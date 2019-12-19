package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-training/proj/core/model"
	"go-training/proj/core/repository"
	"go-training/proj/server"
	"io/ioutil"
	"net/http"
	"strconv"
)

type PostHandler struct {
	PostRepository repository.PostRepository
}

func (h *PostHandler) GetList(w http.ResponseWriter, r *http.Request) {
	categoryId, _ := strconv.ParseInt(r.URL.Query().Get("category"), 10, 32)
	posts, err := h.PostRepository.GetPostList(int32(categoryId))
	if err != nil {
		panic(err)
	}

	server.WriteJSONResponse(w, posts, http.StatusOK)
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		server.WriteErrorResponse(w, err.Error(), http.StatusBadRequest)
	}
	var post model.Post
	if err = json.Unmarshal(body, &post); err != nil {
		server.WriteErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !post.IsValid() {
		server.WriteErrorResponse(w, "Entity is not valid", http.StatusBadRequest)
		return
	}
	// Validate category

	user := r.Context().Value("User").(*model.User)
	if err = h.PostRepository.AddPost(&post, user.Id); err != nil {
		panic(err)
	}

	server.WriteJSONResponse(w, post, http.StatusOK)
}

func (h *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	postId, _ := strconv.Atoi(mux.Vars(r)["postId"])
	if postId == 0 {
		server.Write404Response(w)
		return
	}

	post, err := h.PostRepository.GetPostById(int32(postId))
	if err != nil {
		panic(err)
	}

	if post.Id == 0 {
		server.Write404Response(w)
		return
	}

	server.WriteJSONResponse(w, post, http.StatusOK)
}

func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	post := server.RetrievePostFromRequest(r, h.PostRepository)
	if post == nil {
		server.Write404Response(w)
		return
	}
	user := r.Context().Value("User").(*model.User)
	if user.Id != post.UserId {
		server.WriteErrorResponse(w, "Access Denied", http.StatusForbidden)
	}
	_ = h.PostRepository.DeletePost(post)
	server.WriteEmptyResponse(w)
}

func (h *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	post := server.RetrievePostFromRequest(r, h.PostRepository)
	if post == nil {
		server.Write404Response(w)
		return
	}
	user := r.Context().Value("User").(*model.User)
	fmt.Print(user.Id, post.UserId)
	if user.Id != post.UserId {
		server.WriteErrorResponse(w, "Access Denied", http.StatusForbidden)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	var JSONPost model.Post
	if err = json.Unmarshal(body, &JSONPost); err != nil {
		server.WriteErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	post.Title = JSONPost.Title
	post.Body = JSONPost.Body
	post.CategoryId = JSONPost.CategoryId

	if !post.IsValid() {
		server.WriteErrorResponse(w, "Entity is not valid", http.StatusBadRequest)
		return
	}

	_ = h.PostRepository.UpdatePost(post)
	server.WriteJSONResponse(w, post, http.StatusOK)
}
