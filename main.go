package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"

	entity "github.com/harahash/go-members-club/entity"
)

//храним использованные эмейлы. почему map, а не slice? потому что в map поиск по ключу будет выполняться быстрее
//почему значение bool? потому что bool занимает всего 1 байт памяти. можно поиграться с вариантом map[Email]void, но мне лень проверять, что лучше по памяти
var usedEmails map[string]bool = make(map[string]bool)

//функция отоборажения данных
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Index page has been required")
	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, entity.DB)
	entity.DB.ErrorMessage = ""
}

func CreateMemberHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Method is %v\n", r.Method) //можешь удалить, а можешь оставить. тупо интересно, что будет при редиректе

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Printf("Error while parsing form:%v\n", err)
		}

		name := r.FormValue("name")
		email := r.FormValue("email")
		_, present := usedEmails[email]
		if present {
			entity.DB.ErrorMessage = fmt.Sprintf("Email '%s' already taken", email)
		} else if !NameIsValid(name) {
			entity.DB.ErrorMessage = fmt.Sprintf("Name '%s' is not valid", name)
		} else if !EmailIsValid(email) {
			entity.DB.ErrorMessage = fmt.Sprintf("Email '%s' is not valid", email)
		} else {
			newMember := entity.NewMember(name, email)
			log.Printf("New member=%v", newMember)

			entity.DB.Members = append(entity.DB.Members, newMember)
			usedEmails[email] = true
		}
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func NameIsValid(name string) bool {
	matched, _ := regexp.MatchString(`^[A-Za-z. ]+$`, name)
	log.Println("name matched=", matched)
	return matched
}

func EmailIsValid(email string) bool {
	matched, _ := regexp.MatchString(`^\w+@[a-z]+\.[a-z]+$`, email)
	log.Println("email matched=", matched)
	return matched
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/create_member", CreateMemberHandler)

	log.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
