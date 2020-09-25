package mysql

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB ..
type DB struct {
	*gorm.DB
}

var (
	clientMap = map[string][]*DB{}
)

const (
	defaultTag         = "default"
	readOnlyTagSuffix  = "_readonly"
	defaultReadOnlyTag = defaultTag + readOnlyTagSuffix

	maxOpenConns    = 16384
	connMaxLifeTime = 5 // second
)

// Init init mysql clients
func Init(configs ...*Config) {
	for _, conf := range configs {
		url := conf.GetDSN()
		db, err := gorm.Open(mysql.Open(url), &gorm.Config{NowFunc: func() time.Time { return time.Now().Local() }})
		if err != nil {
			panic(err)
		}

		database, err := db.DB()
		if err != nil {
			panic(err)
		}

		if err = database.Ping(); err != nil {
			panic(err)
		}

		if conf.MaxOpenConns > 0 {
			if conf.MaxOpenConns > maxOpenConns {
				conf.MaxOpenConns = maxOpenConns
			}
			database.SetMaxOpenConns(conf.MaxOpenConns)
		}

		if conf.MaxIdleConns > 0 {
			if conf.MaxIdleConns > conf.MaxOpenConns {
				conf.MaxIdleConns = conf.MaxOpenConns
			}
			database.SetMaxIdleConns(conf.MaxIdleConns)
		}

		if conf.ConnMaxLifeTime == 0 {
			conf.ConnMaxLifeTime = connMaxLifeTime
		}
		database.SetConnMaxLifetime(time.Duration(conf.ConnMaxLifeTime) * time.Second)
	}
}

// GetByTag get db instance by tag
func GetByTag(tag string, xReqID ...string) *DB {
	return nil
}
