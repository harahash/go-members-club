package entity

import "time"

type Member struct {
	Id               int
	Name             string
	Email            string
	RegistrationDate string
}

func NewMember(name string, email string) Member {
	id := len(DB.Members) + 1
	registrationDate := time.Now().Format("02.01.2006") //сегодняшняя дата в формате ДД.ММ.ГГГГ

	newMember := Member{id, name, email, registrationDate}
	return newMember
}
