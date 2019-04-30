/*
	建立Psql数据库连接
*/
package connect

import (
	"fmt"
	conf "mcc-core/config"
	"mcc-core/log"
	"os"
	"sync"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

var (
	ViewPath   string
	StaticPath string
)

var (
	PsqlEngine *xorm.Engine
	onePsql    sync.Once
)

func PsqlInit() {
	NewPsql(&conf.Conf.Psql)
}

// 数据库连接,单例
func NewPsql(config *conf.PsqlConfig) {
	onePsql.Do(func() {
		log.InfoF("connect to psql %s:%s",
			config.Host, config.Port)

		driveSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			config.Host, config.Port, config.User, config.Pwd, config.DB, config.Sslmode)

		engine, err := xorm.NewEngine("postgres", driveSource)
		if err != nil {
			log.Info("psql connect error", err)
			os.Exit(0)
		}

		// Debug模式，打印全部的SQL语句，帮助对比，看ORM与SQL执行的对照关系
		engine.ShowSQL(conf.Conf.Psql.IsShowSQL)
		engine.SetTZLocation(conf.SysTimeChina)

		//设置连接数
		engine.SetMaxIdleConns(20)
		engine.SetMaxOpenConns(100)

		// 性能优化的时候才考虑，加上本机的SQL缓存
		// cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
		// engine.SetDefaultCacher(cacher)

		err = engine.Ping()
		if err != nil {
			log.Crit(err.Error())
			os.Exit(0)
		} else {
			log.Info("psql had connected")
		}
		PsqlEngine = engine
	},
	)
}
