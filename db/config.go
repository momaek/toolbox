package db

import (
	"fmt"
	"net/url"
)

// Config mysql config
// support json and viper yaml config tag
type Config struct {
	Host            string  `json:"host"               mapstructure:"host"`
	Port            int     `json:"port"               mapstructure:"port"`
	Username        string  `json:"username"           mapstructure:"username"`
	Password        string  `json:"password"           mapstructure:"password"`
	Database        string  `json:"database"           mapstructure:"database"`
	MaxIdleConns    int     `json:"max_idle_conns"     mapstructure:"max_idle_conns"`
	MaxOpenConns    int     `json:"max_open_conns"     mapstructure:"max_open_conns"`
	ConnMaxLifeTime int     `json:"conn_max_life_time" mapstructure:"conn_max_life_time"`
	ReadOnly        bool    `json:"read_only"          mapstructure:"read_only"`
	Tag             *string `json:"tag"                mapstructure:"tag"`
	SlowThreshold   int64   `json:"slowthreshold"      mapstructure:"slowthreshold"` // DB 时间大于这个值，就Warning 慢查询 ms 毫秒 默认 50ms
}

// GetDSN get mysql connection url
func (conf *Config) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=%s&parseTime=true",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
		url.QueryEscape("Local"))
}

// GetTag ..
func (conf *Config) GetTag() string {
	tag := defaultTag
	if conf.Tag != nil {
		tag = *conf.Tag
	}

	if conf.ReadOnly {
		tag += readOnlyTagSuffix
	}

	return tag
}
