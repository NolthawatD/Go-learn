package main

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n==========================================================\n", sql)
}

func main() {
	dsn := "root:secret@tcp(127.0.0.1:3306)/gorm?parseTime=true"
	dial := mysql.Open(dsn)                  // dsn is datasource name
	db, err := gorm.Open(dial, &gorm.Config{ // dialector is db driver for your connect
		Logger: &SqlLogger{},
		DryRun: true, // test view sql script only
	})
	if err != nil {
		panic(err)
	}

	db.Migrator().CreateTable(Test{})
	// db.AutoMigrate(Gender{})
}

type Test struct {
	ID   uint
	Code uint   `gorm:"primaryKey; comment:This is Code"`
	Name string `gorm:"column:myname; type:varchar(50); unique; default:Hello: not null"` // size:20
}

// change table name convention
func (t Test) TableName() string {
	return "MyTest"
}
