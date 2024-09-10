package config

import (
	"time"

	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type config struct {
	*viper.Viper
}

var DefaultInstance = &config{viper.New()}

func init() {
	DefaultInstance.SetConfigName("config")
	DefaultInstance.AddConfigPath("./conf")
	if err := DefaultInstance.ReadInConfig(); err != nil {
		logger.Warnf("No config file: %s ", err)
	}
	DefaultInstance.AutomaticEnv()
}

func (c *config) SetConfigName(name string) {
	c.Viper.SetConfigName(name)
}
func (c *config) AddConfigPath(path string) {
	c.Viper.AddConfigPath(path)
}
func (c *config) ReadInConfig() error {
	return c.Viper.ReadInConfig()
}
func (c *config) AutomaticEnv() {
	c.Viper.AutomaticEnv()
}

func (c *config) GetString(key string, defaultValue ...string) string {
	if c.Viper.IsSet(key) {
		return c.Viper.GetString(key)
	}
	if len(defaultValue) == 0 {
		return ""
	}
	return defaultValue[0]
}
func (c *config) GetInt(key string, defaultValue ...int) int {
	if c.Viper.IsSet(key) {
		return c.Viper.GetInt(key)
	}
	if len(defaultValue) == 0 {
		return 0
	}
	return defaultValue[0]
}
func (c *config) GetBool(key string, defaultValue ...bool) bool {
	if c.Viper.IsSet(key) {
		return c.Viper.GetBool(key)
	}
	if len(defaultValue) == 0 {
		return false
	}
	return defaultValue[0]
}
func (c *config) GetUint(key string, defaultValue ...uint) uint {
	if c.Viper.IsSet(key) {
		return c.Viper.GetUint(key)
	}
	if len(defaultValue) == 0 {
		return 0
	}
	return defaultValue[0]
}
func (c *config) GetUint64(key string, defaultValue ...uint64) uint64 {
	if c.Viper.IsSet(key) {
		return c.Viper.GetUint64(key)
	}
	if len(defaultValue) == 0 {
		return 0
	}
	return defaultValue[0]
}
func (c *config) GetFloat64(key string, defaultValue ...float64) float64 {
	if c.Viper.IsSet(key) {
		return c.Viper.GetFloat64(key)
	}
	if len(defaultValue) == 0 {
		return 0
	}
	return defaultValue[0]
}
func (c *config) IsSet(key string) bool {
	return c.Viper.IsSet(key)
}
func (c *config) GetStringSlice(key string) []string {
	return c.Viper.GetStringSlice(key)
}
func (c *config) GetStringMapString(key string) map[string]string {
	return c.Viper.GetStringMapString(key)
}
func (c *config) GetDuration(key string) time.Duration {
	return c.Viper.GetDuration(key)
}
func (c *config) WatchConfig() {
	c.Viper.WatchConfig()
}
func (c *config) Set(key string, value interface{}) {
	c.Viper.Set(key, value)
}
func (c *config) WriteConfig() error {
	return c.Viper.WriteConfig()
}

/////////////////////////////////////////////////////////////////////
//////// 默认配置获取 //////

func SetConfigName(name string) {
	DefaultInstance.SetConfigName(name)
}
func AddConfigPath(path string) {
	DefaultInstance.AddConfigPath(path)
}
func ReadInConfig() error {
	return DefaultInstance.ReadInConfig()
}
func AutomaticEnv() {
	DefaultInstance.AutomaticEnv()
}

func GetString(key string, defaultValue ...string) string {
	return DefaultInstance.GetString(key, defaultValue...)
}
func GetInt(key string, defaultValue ...int) int {
	return DefaultInstance.GetInt(key, defaultValue...)
}
func GetBool(key string, defaultValue ...bool) bool {
	return DefaultInstance.GetBool(key, defaultValue...)
}
func GetUint(key string, defaultValue ...uint) uint {
	return DefaultInstance.GetUint(key, defaultValue...)
}
func GetUint64(key string, defaultValue ...uint64) uint64 {
	return DefaultInstance.GetUint64(key, defaultValue...)
}
func GetFloat64(key string, defaultValue ...float64) float64 {
	return DefaultInstance.GetFloat64(key, defaultValue...)
}
func IsSet(key string) bool {
	return DefaultInstance.IsSet(key)
}
func GetStringSlice(key string) []string {
	return DefaultInstance.GetStringSlice(key)
}
func GetIntSlice(key string) []int {
	return DefaultInstance.GetIntSlice(key)
}
func GetStringMapString(key string) map[string]string {
	return DefaultInstance.GetStringMapString(key)
}
func GetDuration(key string) time.Duration {
	return DefaultInstance.GetDuration(key)
}
func WatchConfig() {
	DefaultInstance.WatchConfig()
}
func Set(key string, value interface{}) {
	DefaultInstance.Set(key, value)
}
func WriteConfig() error {
	return DefaultInstance.WriteConfig()
}
