package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Student2 struct {
	id         int    `gorm:"primary key:true;autoIncrement=true;column:id"`
	name       string `gorm:"column:name"`
	age        int    `gorm:"column:age"`
	address    string `gorm:"column:address"`
	examType   string `gorm:"column:examType"`
	Marksheet2 Marksheet2
}

type Marksheet2 struct {
	id         int    `gorm:"primary key:true;autoIncrement=true;column:id"`
	Student_id int    `gorm:"column:student_id"`
	subject    string `gorm:"column:subject"`
	marks      int    `gorm:"column:marks"`
	Student2ID int
}

// ---------------------------------------------------------------------------------------------------
func main() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "Arijit123@"
		dbname   = "Examination"
	)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Database Connected...")
	}
	var std Student2
	std.name = "Arijit"
	std.age = 23
	std.address = "Barasat"
	std.examType = "Sem 4"

	var mk Marksheet2
	mk.id = 101
	mk.subject = "BENG"
	mk.marks = 90
	std.Marksheet2 = mk
	db.Create(&std)

}
