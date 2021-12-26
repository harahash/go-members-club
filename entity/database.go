package entity

type Database struct {
	Members      []Member
	ErrorMessage string
}

var DB Database = Database{}
