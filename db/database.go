package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

// Database option for configurations
type Config struct {
	User        string `json:"user"`
	Password    string `json:"password"`
	Host        string `json:"host"`
	Port        string `json:"port"`
	Name        string `json:"name"`
	Charset     string `json:"charset"`
	TablePrefix string `json:"tablePrefix"`
}

// Database is wrapped struct of *sql.DB
type Database struct {
	*gorm.DB
}

// Open generate a database client
func Open(conf *Config) *gorm.DB {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&tls=skip-verify&autocommit=true&loc=Local&parseTime=true", conf.User, conf.Password, conf.Host, conf.Port, conf.Name, conf.Charset)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   conf.TablePrefix, // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true,             // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	return db
}

// NewDatabase waper gorm.DB
func NewDatabase(conf *Config) *Database {
	return &Database{
		Open(conf),
	}
}
