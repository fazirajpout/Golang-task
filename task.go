package main


import (
  "log"
  "net/http"
  "github.com/gorilla/mux"
 )
func initializeRouter(){
     r :=mux.NewRouter()
     //Getting list of student
     r.HandleFunc("/students",GetStudents).Methods("GET")
     //Getting student on bases of id
     r.HandleFunc("/students/{id}",GetStudent).Methods("GET")
     //Creating Teacher
     r.HandleFunc("/teachers",CreateTeacher).Methods("POST")
     //Creating Student
     r.HandleFunc("/students",CreateStudent).Methods("POST")
     //Teacher log 
     r.HandleFunc("/login",Login).Methods("POST")
     //Update students on basis of id
     r.HandleFunc("/students/{id}",UpdateStudent).Methods("PUT")
     //Deleting student on basis of id
     r.HandleFunc("/students/{id}",DeleteStudent).Methods("DELETE")
     
     log.Fatal(http.ListenAndServe(":9000",r))
     
}
func main(){
     Migration()
     IntialMigration()
     initializeRouter()
     
}