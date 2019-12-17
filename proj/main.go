package proj

import (
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

	apiRouterMux := http.NewServeMux()
	apiRouterMux.HandleFunc("/api/auth", securityHandler.Authenticate)
	apiRouterMux.HandleFunc("/api/posts", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			securityMiddlewares.Authorization(postHandler.CreatePost).ServeHTTP(w, r)
		} else {
			postHandler.GetList(w, r)
		}
	})

	globalMux := http.NewServeMux()
	globalMux.Handle("/api/", apiRouterMux)
	globalMux.Handle("/", http.FileServer(http.Dir("proj/static")))

	globalHandler := middleware.ServerErrorHandler(securityMiddlewares.Authentication(globalMux))

	if err := http.ListenAndServe(":3000", globalHandler); err != nil {
		panic(err)
	}

}
