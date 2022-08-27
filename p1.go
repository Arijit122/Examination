package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Student struct {
	id        int    `gorm:"primary key:true;autoIncrement=true;column:id"`
	name      string `gorm:"column:name"`
	age       int    `gorm:"column:age"`
	address   string `gorm:"column:address"`
	examType  string `gorm:"column:examType"`
	Marksheet Marksheet
}

type Marksheet struct {
	id         int    `gorm:"primary key:true;autoIncrement=true;column:id"`
	Student_id int    `gorm:"column:student_id"`
	subject    string `gorm:"column:subject"`
	marks      int    `gorm:"column:marks"`
	StudentID  int
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
	//-------------------------------------------------------------------------------------------------
	db.AutoMigrate(&Student{})
	db.AutoMigrate(&Marksheet{})
	// var std Student
	// db.Create(&std)

	// var mk Marksheet
	// mk.StudentID = std.id
	// db.Create(&mk)

	// var std Student
	// std.name = "Arijit"
	// std.age = 22
	// std.address = "Hridaypur"
	// std.examType = "Semester 4"
	// db.Create(&std)

	// var mk Marksheet
	// mk.id = 101
	// mk.subject = "CMSA"
	// mk.marks = 92
	// mk.StudentID = std.id
	// db.Create(&mk)
}
