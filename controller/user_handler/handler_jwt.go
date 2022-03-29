package userhandler

import (
	driver "SQLite_Repo_Pattern/driver"
	models "SQLite_Repo_Pattern/model/user"
	repoImpl "SQLite_Repo_Pattern/repository/repoimpl"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("abcdefghijklmnopq")

type Claims struct {
	Email       string `json:"email"`
	DisplayName string `json:"displayName"`
	jwt.StandardClaims
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applycation/json")
	var regData models.RegistrationData
	err := json.NewDecoder(r.Body).Decode(&regData)
	if err != nil {
		ResponseErr(w, http.StatusBadRequest)
	}
	UserRepo := repoImpl.NewUserRepo(driver.SQLite.SQL)
	fmt.Println(regData.Email)
	val, _ := UserRepo.FindUserByEmail(regData.Email)
	fmt.Println(val)
	if val == true {
		ResponseErr(w, http.StatusConflict)
	}

	user := models.User{
		Email:       regData.Email,
		Password:    regData.Password,
		DisplayName: regData.DisplayName,
	}
	UserRepoo := repoImpl.NewUserRepo(driver.SQLite.SQL)
	err = UserRepoo.Isert(user)
	if err != nil {
		ResponseErr(w, http.StatusInternalServerError)
		return
	}

	var tokenString string
	tokenString, err = GenToken(user)
	if err != nil {
		ResponseErr(w, http.StatusInternalServerError)
		return
	}

	ResponseOk(w, models.RegisterResponse{
		Token:  tokenString,
		Status: http.StatusOK,
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applycation/json")
	var loginData models.LoginData
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		ResponseErr(w, http.StatusBadRequest)
		return
	}
	UserRepo := repoImpl.NewUserRepo(driver.SQLite.SQL)
	user, _ := UserRepo.CheckLoginInfo(loginData.Email, loginData.Password)
	// if err != nil {
	// 	ResponseErr(w, http.StatusUnauthorized)
	// 	return
	// }

	var tokenString string
	tokenString, err = GenToken(user)
	if err != nil {
		ResponseErr(w, http.StatusInternalServerError)
		return
	}

	ResponseOk(w, models.RegisterResponse{
		Token:  tokenString,
		Status: http.StatusOK,
	})
}

func GenToken(user models.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Email:       user.Email,
		DisplayName: user.DisplayName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ResponseErr(w http.ResponseWriter, statusCode int) {
	jData, err := json.Marshal(models.Error{
		Status:  statusCode,
		Message: http.StatusText(statusCode),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

func ResponseOk(w http.ResponseWriter, data interface{}) {
	if data == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}
