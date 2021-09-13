package main


import (
  "log"
  "net/http"
  "github.com/gorilla/mux"
 )
func initializeRouter(){
     r :=mux.NewRouter()
     r.HandleFunc("/students",GetStudents).Methods("GET")
     r.HandleFunc("/students/{id}",GetStudent).Methods("GET")
     r.HandleFunc("/teachers",CreateTeacher).Methods("POST")
     r.HandleFunc("/students",CreateStudent).Methods("POST")
     r.HandleFunc("/login",Login).Methods("POST")
     r.HandleFunc("/students/{id}",UpdateStudent).Methods("PUT")
     r.HandleFunc("/students/{id}",DeleteStudent).Methods("DELETE")
     
     log.Fatal(http.ListenAndServe(":9000",r))
     
}
func main(){
     Migration()
     IntialMigration()
     initializeRouter()
     
}