package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-training/proj/core/model"
	"go-training/proj/core/repository"
	"net/http"
	"strconv"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func WriteErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	WriteJSONResponse(w, ErrorResponse{Code: statusCode, Message: message}, statusCode)
}

func WriteJSONResponse(w http.ResponseWriter, obj interface{}, statusCode int) {
	body, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if _, err := fmt.Fprintln(w, string(body)); err != nil {
		panic(err)
	}

}

func Write404Response(w http.ResponseWriter) {
	WriteErrorResponse(w, "Entity is not found", http.StatusNotFound)
}

func WriteEmptyResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func RetrievePostFromRequest(r *http.Request, postRep repository.PostRepository) *model.Post {
	postId, _ := strconv.Atoi(mux.Vars(r)["postId"])
	if postId == 0 {
		return nil
	}

	post, err := postRep.GetPostById(int32(postId))
	if err != nil {
		panic(err)
	}

	return post
}
