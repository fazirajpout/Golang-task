package main


import (
   "encoding/json"
   "net/http"
   "gorm.io/driver/mysql"
   "gorm.io/gorm"
   "fmt"
   "github.com/gorilla/mux"
    "time"
    "github.com/dgrijalva/jwt-go"
 )
var DB *gorm.DB
var err error
const DNS = "root:henry37572@tcp(host.docker.internal:3306)/go_db?parseTime=true"

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
var jwtKey = []byte("secret_key")
type Claims struct {
	Name string `json:"name"`
	jwt.StandardClaims
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
      cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
      DB.Find(&students)
      json.NewEncoder(w).Encode(students)

}
func GetStudent(w http.ResponseWriter, r *http.Request){
      w.Header().Set("content-Type", "application/json")
      params := mux.Vars(r)
      var student Student
      cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
      DB.Find(&student, params["id"]) 
      json.NewEncoder(w).Encode(student)
}
func CreateStudent(w http.ResponseWriter, r *http.Request){
     w.Header().Set("Content-Type", "application/json")
     var student Student
     json.NewDecoder(r.Body).Decode(&student)
     cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
     DB.Create(&student)
     json.NewEncoder(w).Encode(student)
}
func UpdateStudent(w http.ResponseWriter, r *http.Request){
     w.Header().Set("content-Type", "application/json")
     params := mux.Vars(r)
     var student Student
     DB.First(&student, params["id"])
     cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
     json.NewDecoder(r.Body).Decode(&student)
     DB.Save(&student)
     json.NewEncoder(w).Encode(student)
}
func DeleteStudent(w http.ResponseWriter, r *http.Request){
     w.Header().Set("content-Type", "application/json")
      params := mux.Vars(r)
      var student Student
      cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
      DB.Delete(&student, params["id"])
      json.NewEncoder(w).Encode("The User is Successfully Deleted")
}
func Login(w http.ResponseWriter, r *http.Request) {
	var teacher Teacher
	err := json.NewDecoder(r.Body).Decode(&teacher)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	DB.Where("name = ? AND password = ?",teacher.Name,teacher.Password).First(&teacher)

	if teacher.Name == "" &&  teacher.Password == ""  {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(time.Minute * 5)

	claims := &Claims{
		Name: teacher.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
        json.NewEncoder(w).Encode("LogIn is success full !! jwt Cookie is created in Postman cookies i.e is valid for 5 minute")
}