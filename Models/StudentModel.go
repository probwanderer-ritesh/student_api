package Models

type Student struct {
	Id        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	DOB       string `json:"DOB"`
	Address   string `json:"address"`
	Marks     Marks  `json:"marks"`
}
type Marks struct {
	Maths   int `json:"maths"`
	English int `json:"english"`
}

func (b *Student) TableName() string {
	return "student"
}
