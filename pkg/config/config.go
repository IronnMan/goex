package config

import (
	"github.com/spf13/cast"
	viperlib "github.com/spf13/viper"
	"goex/pkg/helpers"
	"os"
)

var viper *viperlib.Viper

// ConfigFunc Dynamically load configuration information
type ConfigFunc func() map[string]interface{}

var ConfigFuncs map[string]ConfigFunc

func init() {
	// Initialize library viper
	viper = viperlib.New()
	// Configuration type, support "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
	viper.SetConfigType("env")
	// The path to look for the environment variable configuration file, relative to `main.go`
	viper.AddConfigPath(".")
	// Set the environment variable prefix to distinguish Golang system environment variables
	viper.SetEnvPrefix("appenv")
	// Read environment variables
	viper.AutomaticEnv()

	ConfigFuncs = make(map[string]ConfigFunc)
}

func InitConfig(env string) {
	// 1. load environment variables
	loadEnv(env)
	// 2. register configuration information
	loadConfig()
}

func loadConfig() {
	for name, fn := range ConfigFuncs {
		viper.Set(name, fn())
	}
}

func loadEnv(envSuffix string) {
	envPath := ".env"
	if len(envSuffix) > 0 {
		filepath := ".env." + envSuffix
		if _, err := os.Stat(filepath); err == nil {
			envPath = filepath
		}
	}

	// load env
	viper.SetConfigName(envPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// Monitor .env files, reload when changed
	viper.WatchConfig()
}

// Env read environment variables, support default values
func Env(envName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return internalGet(envName, defaultValue[0])
	}
	return internalGet(envName)
}

// Add new configuration items
func Add(name string, configFn ConfigFunc) {
	ConfigFuncs[name] = configFn
}

// Get get configuration items
func Get(path string, defaultValue ...interface{}) string {
	return GetString(path, defaultValue)
}

func internalGet(path string, defaultValue ...interface{}) interface{} {
	if !viper.IsSet(path) || helpers.Empty(viper.Get(path)) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return viper.Get(path)
}

// GetString Get string type config information
func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(internalGet(path, defaultValue...))
}

// GetInt Get int type config information
func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(internalGet(path, defaultValue...))
}

// GetInt64 Get int type config information
func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(internalGet(path, defaultValue...))
}

// GetFloat64 Get float64 type config information
func GetFloat64(path string, defaultValue ...interface{}) float64 {
	return cast.ToFloat64(internalGet(path, defaultValue...))
}

// GetUint Get uint type config information
func GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(internalGet(path, defaultValue...))
}

// GetBool Get bool type config information
func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(internalGet(path, defaultValue...))
}

// GetStringMapString 获取结构数据
func GetStringMapString(path string) map[string]string {
	return viper.GetStringMapString(path)
}
