package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n==========================================================\n", sql)
}

var db *gorm.DB

func main() {
	dsn := "root:secret@tcp(127.0.0.1:3306)/gorm?parseTime=true"
	dial := mysql.Open(dsn) // dsn is datasource name

	var err error
	db, err = gorm.Open(dial, &gorm.Config{ // dialector is db driver for your connect
		Logger: &SqlLogger{},
		DryRun: false, // test view sql script only
	})
	if err != nil {
		panic(err)
	}

	// db.Migrator().CreateTable(Test{})
	// db.AutoMigrate(Gender{}, Test{})

	// CreateGender("XYZ")
	// GetGenders()
	// GetGender(10)
	// GetGenderByName("Male")

	// UpdateGender2(4, "ZYX")
	// DeleteGender(4)

	db.Migrator().CreateTable(Customer{})

	CreateCustomer("Apple", 2)

	UpdateGender2(1, "ชาย")
	GetCustomers()

}

func GetCustomers() {
	customers := []Customer{}
	tx := db.Preload(clause.Associations).Find(&customers)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	for _, customer := range customers {
		fmt.Printf("%v|%v|%v\n", customer.ID, customer.Name, customer.Gender.Name)
	}
}

func CreateCustomer(name string, genderID uint) {
	customer := Customer{Name: name, GenderID: genderID}
	tx := db.Create(&customer)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(customer)
}

type Customer struct {
	ID       uint
	Name     string
	Gender   Gender
	GenderID uint
}

// DELETE
func DeleteGender(id uint) {
	tx := db.Delete(&Gender{}, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println("Deleted")
	GetGender(id)
}

// DELETE PERMANENTLY
func DeleteTest(id uint) {
	db.Unscoped().Delete(&Test{}, id)
}

// UPDATE
func UpdateGender(id uint, name string) {
	gender := Gender{}
	tx := db.First(&gender, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	gender.Name = name
	tx = db.Save(&gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	GetGender(id)
}

func UpdateGender2(id uint, name string) {
	gender := Gender{Name: name}
	tx := db.Model(&Gender{}).Where("id=@myid", sql.Named("myid", id)).Updates(gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	GetGender(id)
}

// READ
func GetGenders() {
	genders := []Gender{}
	tx := db.Order("id").Find(&genders)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(genders)
}

func GetGenderByName(name string) {
	genders := []Gender{}
	tx := db.Where("name=?", name).Find(&genders)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(genders)
}

func GetGender(id uint) {
	gender := Gender{}
	tx := db.First(&gender, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(gender)
}

// INSERT
func CreateGender(name string) {
	gender := Gender{Name: name}
	tx := db.Create(&gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}

	fmt.Println(gender)
}

type Gender struct {
	ID   uint
	Name string `gorm:"unique;size(10)"`
}

type Test struct {
	gorm.Model
	Code uint   `gorm:"primaryKey; comment:This is Code"`
	Name string `gorm:"column:myname; type:varchar(50); unique; default:Hello: not null"` // size:20
}

// change table name convention
func (t Test) TableName() string {
	return "MyTest"
}
