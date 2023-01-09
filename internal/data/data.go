/*
 * @Author: gengqingyang gengqingyang@onething.net
 * @Date: 2022-11-26 23:15:23
 * @LastEditTime: 2022-12-18 16:12:38
 * @Description:

 */
package data

import (
	"fmt"
	"krato_study/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewUserRepo, NewArticleRepo, NewCommentRepo, NewTagRepo, NewProfileRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db}, cleanup, nil
}

func NewDB(c *conf.Data) *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       c.Database.Dsn, // DSN data source name
		DefaultStringSize:         256,            // string 类型字段的默认长度
		DisableDatetimePrecision:  true,           // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,           // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,           // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,          // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	if err != nil {
		fmt.Println("init db err is ", err.Error())
		return nil
	}
	if err = db.AutoMigrate(&User{}); err != nil {
		fmt.Println("init db create user err is ", err.Error())
		return nil
	}
	return db
}
