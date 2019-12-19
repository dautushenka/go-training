package handlers

import (
	"go-training/proj/core/model"
	"go-training/proj/core/repository"
	"go-training/proj/core/security"
	"go-training/proj/server"
	"net/http"
)

type SecurityHandler struct {
	UsersRepository repository.UserRepository
}

type authenticateResponse struct {
	Token string      `json:"token,omitempty"`
	User  *model.User `json:"user"`
}

func (h *SecurityHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")

	if login == "" || password == "" {
		server.WriteErrorResponse(w, "Login or password is empty", 400)
		return
	}

	user, err := h.UsersRepository.GetUserByLogin(login)
	if err != nil || !security.IsPasswordValid(password, user.PasswordHash) {
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
