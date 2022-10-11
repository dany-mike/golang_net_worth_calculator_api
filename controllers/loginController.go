package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"golang_net_worth_calculator_api/auth"
	"golang_net_worth_calculator_api/models"
	"golang_net_worth_calculator_api/responses"
	"golang_net_worth_calculator_api/utils/formatError"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string
	Email    string
}

type SigninResponse struct {
	User  User
	Token string
}

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	response, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := formatError.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) SignIn(email, password string) (SigninResponse, error) {

	var err error

	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return SigninResponse{}, err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return SigninResponse{}, err
	}
	token, err := auth.CreateToken(user.ID)

	userInstance := User{user.Username, user.Email}
	return SigninResponse{userInstance, token}, err
}
