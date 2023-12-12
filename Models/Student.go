package Models

import (
	"fmt"
	"students-api/Config"

	_ "github.com/go-sql-driver/mysql"
)

//GetAllStudent Fetch all student data

func GetAllUsers(student *[]Student) (err error) {
	if err = Config.DB.Find(student).Error; err != nil {
		return err
	}
	return nil
}

// CreateStudent ... Insert New Data
func CreateStudent(student *Student) (err error) {
	newStudent := &Student{
		FirstName: student.FirstName,
		LastName:  student.LastName,
		DOB:       student.DOB,
		Address:   student.Address,
		Marks: Marks{
			Maths:   student.Marks.Maths,
			English: student.Marks.English,
		},
	}
	if err = Config.DB.Create(newStudent).Error; err != nil {
		return err
	}
	return nil
}

// GetStudentByID ... Fetch only one student by Id
func GetStudentByID(student *Student, id string) (err error) {
	if err = Config.DB.Where("id=?", id).First(student).Error; err != nil {
		return err
	}
	return nil
}

// UpdateStudent... Update Student
func UpdateStudent(student *Student, id string) (err error) {
	fmt.Println(student)
	Config.DB.Save(student)
	return nil
}

// DeleteStudent ... Delete Student
func DeleteStudent(student *Student, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(student)
	return nil
}
