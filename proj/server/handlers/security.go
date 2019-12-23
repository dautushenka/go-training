package handlers

import (
	"encoding/json"
	"go-training/proj/core/model"
	"go-training/proj/core/repository"
	"go-training/proj/core/security"
	"go-training/proj/server"
	"io/ioutil"
	"net/http"
)

type SecurityHandler struct {
	UsersRepository repository.UserRepository
}

type authenticateResponse struct {
	Token string      `json:"token,omitempty"`
	User  *model.User `json:"user"`
}

type authenticateRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func CreateSecurityHandler(repository repository.UserRepository) *SecurityHandler {
	return &SecurityHandler{repository}
}

func (h *SecurityHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}

	var request authenticateRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		server.WriteErrorResponse(w, err.Error(), 400)
		return
	}

	if request.Login == "" || request.Password == "" {
		server.WriteErrorResponse(w, "Login or password is empty", 400)
		return
	}

	user, err := h.UsersRepository.GetUserByLogin(request.Login)
	if err != nil || !security.IsPasswordValid(request.Password, user.PasswordHash) {
		server.WriteErrorResponse(w, "Login/Password is invalid", 400)
		return
	}

	response := authenticateResponse{
		Token: security.GetToken(user),
		User:  user,
	}

	server.WriteJSONResponse(w, response, 200)
}

func (h *SecurityHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	if user := r.Context().Value("User"); user != nil {
		response := authenticateResponse{
			User: user.(*model.User),
		}
		server.WriteJSONResponse(w, response, 200)
		return
	}
	server.WriteErrorResponse(w, "Unauthorized", 401)
}
