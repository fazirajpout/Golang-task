package main

import (
   "encoding/json"
   "net/http"
   "gorm.io/driver/mysql"
   "gorm.io/gorm"
   "fmt"
   "github.com/gorilla/mux"
 )
var DB *gorm.DB
var err error
const DNS = "root:henry37572@tcp(127.0.0.1:3306)/go_db?parseTime=true"

type Student struct{
     gorm.Model
     Name string `json:"name" gorm:"unique"`
     Age  string `json:"Age"`
     City string `json:"city"`
     Subject string `json:"subject"`
     Email string `json:"email" gorm:"unique"`
     Class string `json:"class"`

     
}
type Teacher struct{
        gorm.Model
	Name string `json:"name"`
        Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}
func IntialMigration(){
     DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
     if err!= nil{
     fmt.Println(err.Error())
     panic("cannot connect to DataBase")
     }
     DB.AutoMigrate(&Student{}) 
}
func Migration(){
     DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
     if err!= nil{
     fmt.Println(err.Error())
     panic("cannot connect to DataBase")
     }
     DB.AutoMigrate(&Teacher{})
}
func CreateTeacher(w http.ResponseWriter, r *http.Request){
     w.Header().Set("Content-Type", "application/json")
     var teacher Teacher
     json.NewDecoder(r.Body).Decode(&teacher)
     DB.Create(&teacher)
     json.NewEncoder(w).Encode(teacher)
}

func GetStudents(w http.ResponseWriter, r *http.Request){
      w.Header().Set("content-Type", "application/json")
      var students []Student
      DB.Find(&students)
      json.NewEncoder(w).Encode(students)

}
func GetStudent(w http.ResponseWriter, r *http.Request){
      w.Header().Set("content-Type", "application/json")
      params := mux.Vars(r)
      var student Student
      DB.First(&student, params["id"])
      json.NewEncoder(w).Encode(student)
}
func CreateStudent(w http.ResponseWriter, r *http.Request){
     w.Header().Set("Content-Type", "application/json")
     var student Student
     json.NewDecoder(r.Body).Decode(&student)
     DB.Create(&student)
     json.NewEncoder(w).Encode(student)
}
func UpdateStudent(w http.ResponseWriter, r *http.Request){
     w.Header().Set("content-Type", "application/json")
     params := mux.Vars(r)
     var student Student
     DB.First(&student, params["id"])
     json.NewDecoder(r.Body).Decode(&student)
     DB.Save(&student)
     json.NewEncoder(w).Encode(student)
}
func DeleteStudent(w http.ResponseWriter, r *http.Request){
     w.Header().Set("content-Type", "application/json")
      params := mux.Vars(r)
      var student Student
      DB.Delete(&student, params["id"])
      json.NewEncoder(w).Encode("The User is Successfully Deleted")
}