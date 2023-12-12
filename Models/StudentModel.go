package Models

import "github.com/jinzhu/gorm"

type Student struct {
	gorm.Model
	Id        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	DOB       string `json:"DOB"`
	Address   string `json:"address"`
	MarksId   uint
	Marks     Marks `json:"marks"`
}

type Marks struct {
	gorm.Model
	Maths   int `json:"maths"`
	English int `json:"english"`
}

func (b *Student) TableName() string {
	return "student"
}

func (m *Marks) TableName() string {
	return "marks"
}
