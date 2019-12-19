package proj

import (
	"github.com/gorilla/mux"
	"go-training/proj/core"
	"go-training/proj/core/repository"
	"go-training/proj/server/handlers"
	"go-training/proj/server/middleware"
	"net/http"
)

func Start() {
	db := core.InitializeDB()
	defer db.Close()

	usersRepository := repository.CreateUserRepository(db)
	postRepository := repository.CreatePostRepository(db)

	securityMiddlewares := middleware.Security{usersRepository}

	securityHandler := handlers.SecurityHandler{usersRepository}
	postHandler := handlers.PostHandler{postRepository}

	apiRouterMux := mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	apiRouterMux.HandleFunc("/auth", securityHandler.Authenticate).Methods(http.MethodPost)
	apiRouterMux.HandleFunc("/auth", securityHandler.GetUser).Methods(http.MethodGet)
	apiRouterMux.HandleFunc("/post", postHandler.GetList).Methods(http.MethodGet)
	apiRouterMux.HandleFunc("/post/{postId:[0-9]+}", postHandler.GetPost).Methods(http.MethodGet)

	apiRouterMuxAuth := mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	apiRouterMuxAuth.Use(securityMiddlewares.Authorization)
	apiRouterMuxAuth.HandleFunc("/post/{postId:[0-9]+}", postHandler.UpdatePost).Methods(http.MethodPut)
	apiRouterMuxAuth.HandleFunc("/post/{postId:[0-9]+}", postHandler.DeletePost).Methods(http.MethodDelete)
	apiRouterMuxAuth.HandleFunc("/post", postHandler.CreatePost).Methods(http.MethodPost)

	apiRouterMux.Handle("/post{_dummy:.*}", apiRouterMuxAuth)

	globalRouter := mux.NewRouter()
	globalRouter.Use(securityMiddlewares.Authentication)
	globalRouter.Use(middleware.ServerErrorHandler)
	globalRouter.Handle("/api/v1/{_dummy:.*}", apiRouterMux)
	globalRouter.Handle("/", http.FileServer(http.Dir("proj/static")))

	if err := http.ListenAndServe(":3000", globalRouter); err != nil {
		panic(err)
	}

}