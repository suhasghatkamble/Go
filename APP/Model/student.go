package model

type Student struct {
    Name         string `gorm:"column:name;type:varchar(255);not null"`
    Roll_No      int    `gorm:"column:roll_no;type:int;primary_key"`
    Std          string `gorm:"column:std;type:varchar(10)"`
    Parents_Name string `gorm:"column:parents_name;type:varchar(255)"`
    Phone_No     int    `gorm:"column:phone_no;type:int"`
    Address      string `gorm:"column:address;type:varchar(255)"`
    DOB          string `gorm:"column:DOB;type:date"`
}