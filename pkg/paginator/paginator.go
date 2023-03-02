package paginator

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"goex/pkg/config"
	"goex/pkg/logger"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math"
	"strings"
)

type Paging struct {
	CurrentPage int
	PerPage     int
	TotalPage   int
	TotalCount  int64
	NextPageURL string
	PrevPageURL string
}

type Paginator struct {
	BaseURL    string
	PerPage    int
	Page       int
	Offset     int
	TotalCount int64
	TotalPage  int
	Sort       string
	Order      string

	query *gorm.DB
	ctx   *gin.Context
}

func Paginate(c *gin.Context, db *gorm.DB, data interface{}, baseURL string, perPage int) Paging {
	p := &Paginator{
		query: db,
		ctx:   c,
	}
	p.initProperties(perPage, baseURL)

	err := p.query.Preload(clause.Associations).
		Order(p.Sort + " " + p.Order).
		Limit(p.PerPage).
		Offset(p.Offset).
		Find(data).
		Error

	if err != nil {
		logger.LogIf(err)
		return Paging{}
	}

	return Paging{
		CurrentPage: p.Page,
		PerPage:     p.PerPage,
		TotalPage:   p.TotalPage,
		TotalCount:  p.TotalCount,
		NextPageURL: p.getNextPageURL(),
		PrevPageURL: p.getPrevPageURL(),
	}
}

func (p Paginator) initProperties(perPage int, baseURL string) {
	p.BaseURL = p.formatBaseURL(baseURL)
	p.PerPage = p.getPerPage(perPage)

	p.Order = p.ctx.DefaultQuery(config.Get("paging.url_query_order"), "asc")
	p.Sort = p.ctx.DefaultQuery(config.Get("paging.url_query_sort"), "id")

	p.TotalCount = p.getTotalCount()
	p.TotalPage = p.getTotalPage()
	p.Page = p.getCurrentPage()
	p.Offset = (p.Page - 1) * p.PerPage
}

func (p Paginator) formatBaseURL(baseURL string) string {
	if strings.Contains(baseURL, "?") {
		baseURL = baseURL + "&" + config.Get("paging.url_query_page") + "="
	} else {
		baseURL = baseURL + "?" + config.Get("paging.url_query_page") + "="
	}
	return baseURL
}

func (p Paginator) getPerPage(perPage int) int {
	queryPerpage := p.ctx.Query(config.Get("paging.url_query_per_page"))
	if len(queryPerpage) > 0 {
		perPage = cast.ToInt(queryPerpage)
	}

	if perPage <= 0 {
		perPage = config.GetInt("paging.perpage")
	}

	return perPage
}

func (p Paginator) getTotalCount() int64 {
	var count int64
	if err := p.query.Count(&count).Error; err != nil {
		return 0
	}
	return count
}

func (p Paginator) getTotalPage() int {
	if p.TotalCount == 0 {
		return 0
	}
	nums := int64(math.Ceil(float64(p.TotalCount) / float64(p.PerPage)))
	if nums == 0 {
		nums = 1
	}
	return int(nums)
}

func (p Paginator) getCurrentPage() int {
	page := cast.ToInt(p.ctx.Query(config.Get("paging.url_query_page")))
	if page <= 0 {
		page = 1
	}
	if p.TotalPage == 0 {
		return 0
	}
	if page > p.TotalPage {
		return p.TotalPage
	}
	return page
}

func (p Paginator) getNextPageURL() string {
	if p.TotalPage > p.Page {
		return p.getPageLink(p.Page + 1)
	}
	return ""
}

func (p Paginator) getPageLink(page int) string {
	return fmt.Sprintf("%v%v&%s=%s&%s=%s&%s=%v",
		p.BaseURL,
		page,
		config.Get("paging.url_query_sort"),
		p.Sort,
		config.Get("paging.url_query_order"),
		p.Order,
		config.Get("paging.url_query_per_page"),
		p.PerPage,
	)
}

func (p Paginator) getPrevPageURL() string {
	if p.Page <= 1 || p.Page > p.TotalPage {
		return ""
	}
	return p.getPageLink(p.Page - 1)
}
