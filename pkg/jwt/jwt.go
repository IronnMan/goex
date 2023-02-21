package jwt

import (
	"errors"
	"github.com/gin-gonic/gin"
	jwtpkg "github.com/golang-jwt/jwt/v4"
	"goex/pkg/app"
	"goex/pkg/config"
	"goex/pkg/logger"
	"strings"
	"time"
)

var (
	ErrTokenExpired           error = errors.New("token has expired")
	ErrTokenExpiredMaxRefresh error = errors.New("the token has passed the maximum refresh time")
	ErrTokenMalformed         error = errors.New("malformed request token")
	ErrTokenInvalid           error = errors.New("invalid request token")
	ErrHeaderEmpty            error = errors.New("authentication is required to access")
	ErrHeaderMalformed        error = errors.New("authorization in the request header is malformed formatted")
)

type JWT struct {
	SignKey    []byte
	MaxRefresh time.Duration
}

type UserInfo struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
}

type CustomJWTClaims struct {
	UserInfo
	ExpireAtTime int64 `json:"expire_at_time"`

	jwtpkg.RegisteredClaims
}

func NewJWT() *JWT {
	return &JWT{
		SignKey:    []byte(config.GetString("app.key")),
		MaxRefresh: time.Duration(config.GetInt64("jwt.max_refresh_time")) * time.Minute,
	}
}

func (j *JWT) ParserToken(ctx *gin.Context) (*CustomJWTClaims, error) {

	tokenStr, parseErr := j.getTokenFromHeader(ctx)
	if parseErr != nil {
		return nil, parseErr
	}

	token, err := j.parseTokenString(tokenStr)

	if err != nil {
		validationErr, ok := err.(*jwtpkg.ValidationError)
		if ok {
			if validationErr.Errors == jwtpkg.ValidationErrorMalformed {
				return nil, ErrTokenMalformed
			} else if validationErr.Errors == jwtpkg.ValidationErrorExpired {
				return nil, ErrTokenExpired
			}
		}
		return nil, ErrTokenInvalid
	}

	if claims, ok := token.Claims.(*CustomJWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid
}

func (j *JWT) RefreshToken(ctx *gin.Context) (string, error) {
	tokenStr, parseErr := j.getTokenFromHeader(ctx)
	if parseErr != nil {
		return "", parseErr
	}

	token, err := j.parseTokenString(tokenStr)

	if err != nil {
		validationErr, ok := err.(*jwtpkg.ValidationError)

		if !ok || validationErr.Errors != jwtpkg.ValidationErrorExpired {
			return "", err
		}
	}

	claims := token.Claims.(*CustomJWTClaims)

	t := app.TimeNowInTimezone().Add(-j.MaxRefresh).Unix()

	if claims.IssuedAt.Unix() > t {
		claims.RegisteredClaims.ExpiresAt = jwtpkg.NewNumericDate(j.expireAtTime())
		return j.createToken(*claims)
	}

	return "", ErrTokenExpiredMaxRefresh
}

func (j *JWT) IssueToken(info UserInfo) string {
	expireTime := j.expireAtTime()
	claims := CustomJWTClaims{
		UserInfo: UserInfo{
			UserID:   info.UserID,
			UserName: info.UserName,
		},

		ExpireAtTime: expireTime.Unix(),
		RegisteredClaims: jwtpkg.RegisteredClaims{
			NotBefore: jwtpkg.NewNumericDate(app.TimeNowInTimezone()), // effective time of signature
			IssuedAt:  jwtpkg.NewNumericDate(app.TimeNowInTimezone()), // first signature time (subsequent refresh token will not be updated)
			ExpiresAt: jwtpkg.NewNumericDate(expireTime),              // signature expiration time
			Issuer:    config.GetString("app.name"),                   // signature issuer
		},
	}

	token, err := j.createToken(claims)
	if err != nil {
		logger.LogIf(err)
		return ""
	}
	return token
}

// getTokenFromHeader use jwtpkg.ParseWithClaims to parse token
// Authorization:Bearer xxxxx
func (j *JWT) getTokenFromHeader(ctx *gin.Context) (string, error) {
	authHeader := ctx.Request.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrHeaderEmpty
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", ErrHeaderMalformed
	}
	return parts[1], nil
}

func (j *JWT) parseTokenString(token string) (*jwtpkg.Token, error) {
	return jwtpkg.ParseWithClaims(token, &CustomJWTClaims{}, func(token *jwtpkg.Token) (interface{}, error) {
		return j.SignKey, nil
	})
}

// expireAtTime token expiration time
func (j *JWT) expireAtTime() time.Time {
	timezone := app.TimeNowInTimezone()

	var expireTime int64
	if config.GetBool("app.debug") {
		expireTime = config.GetInt64("jwt.debug_expire_time")
	} else {
		expireTime = config.GetInt64("jwt.expire_time")
	}

	expire := time.Duration(expireTime) * time.Minute
	return timezone.Add(expire)
}

func (j *JWT) createToken(claims CustomJWTClaims) (string, error) {
	t := jwtpkg.NewWithClaims(jwtpkg.SigningMethodHS256, claims)
	return t.SignedString(j.SignKey)
}
