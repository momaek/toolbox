package mysql

import (
	"math/rand"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//"github.com/momaek/toolbox/logger"
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

		tag := conf.GetTag()
		clientMap[tag] = append(clientMap[tag], &DB{db})
	}
}

// GetByTag get db instance by tag
func GetByTag(tag string, xReqID ...string) *DB {
	/*
		var reqID = ""
		if len(xReqID) > 0 {
			reqID = xReqID[0]
		} else {
			reqID = logger.GenReqID()
		}
	*/

	if tag == "" {
		tag = defaultTag
	}

	clients := clientMap[tag]
	client := clients[rand.Intn(len(clients))]

	db := client.Session(&gorm.Session{Logger: nil})
	return &DB{db}
}

// GetByTagReadOnly get tag readonly mysql
func GetByTagReadOnly(tag string, xReqID ...string) *DB {
	return GetByTag(tag+readOnlyTagSuffix, xReqID...)
}

// Get get default tag mysql
func Get(xReqID ...string) *DB {
	return GetByTag(defaultTag, xReqID...)
}

// GetReadOnly get default read only tag mysql
func GetReadOnly(xReqID ...string) *DB {
	return GetByTag(defaultReadOnlyTag, xReqID...)
}
