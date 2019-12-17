package handlers

import (
	"go-training/proj/core/repository"
	"go-training/proj/server"
	"net/http"
)

type PostHandler struct {
	PostRepository repository.PostRepository
}

func (h *PostHandler) GetList(w http.ResponseWriter, r *http.Request) {
	posts, err := h.PostRepository.GetPostList()
	if err != nil {
		panic(err)
	}

	server.WriteJSONResponse(w, posts, 200)
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test"))
}

func (h *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {

}

func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {

}

func (h *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {

}
