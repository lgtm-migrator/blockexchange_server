package web

import (
	"blockexchange/db"
	"blockexchange/types"
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gorilla/mux"
)

type UserApi struct {
	UserRepo db.UserRepository
}

func (api UserApi) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := api.UserRepo.GetUsers()
	if err != nil {
		SendError(w, err.Error())
		return
	}
	sanitizedUsers := make([]types.User, len(users))
	for i, user := range users {
		sanitizedUsers[i] = types.User{
			ID:      user.ID,
			Created: user.Created,
			Name:    user.Name,
			Type:    user.Type,
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(&sanitizedUsers)
}

func (api UserApi) UpdateUser(w http.ResponseWriter, r *http.Request, ctx *SecureContext) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		SendError(w, err.Error())
		return
	}

	if id != ctx.Token.UserID {
		SendError(w, "user mismatch")
		return
	}

	sentuser := types.User{}
	err = json.NewDecoder(r.Body).Decode(&sentuser)
	if err != nil {
		SendError(w, err.Error())
		return
	}

	valid, _, err := api.ValidateUsername(sentuser.Name)
	if err != nil {
		SendError(w, err.Error())
		return
	}

	user, err := api.UserRepo.GetUserById(id)
	if err != nil {
		SendError(w, err.Error())
		return
	}

	if valid {
		// only update username if it is valid
		user.Name = sentuser.Name
	}
	user.Mail = sentuser.Mail

	err = api.UserRepo.UpdateUser(user)
	if err != nil {
		SendError(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(&user)
}

func (api UserApi) PostValidateUsername(w http.ResponseWriter, r *http.Request) {
	sentuser := types.User{}
	err := json.NewDecoder(r.Body).Decode(&sentuser)
	if err != nil {
		SendError(w, err.Error())
		return
	}

	valid, msg, err := api.ValidateUsername(sentuser.Name)
	if err != nil {
		SendError(w, err.Error())
		return
	}

	SendJson(w, types.ValidationResult{
		Valid:   valid,
		Message: msg,
	})
}

var username_regex = regexp.MustCompile("^[a-zA-Z0-9_-]+$")

func (api UserApi) ValidateUsername(username string) (bool, string, error) {
	if !username_regex.MatchString(username) {
		return false, "Username can only contain characters,numbers,dashes and underlines", nil
	}
	existinguser, err := api.UserRepo.GetUserByName(username)
	if err != nil {
		return false, "", err
	}
	if existinguser != nil {
		return false, "User already exists", nil
	}

	return true, "", nil
}
