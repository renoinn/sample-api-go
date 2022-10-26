package main

import (
	"net/http"

	"github.com/renoinn/sample-api/controller"
	"github.com/renoinn/sample-api/model/repository"
)

var repo = repository.NewBookmarkRepository()
var cont = controller.NewController(repo)
var router = controller.NewRouter(cont)

func main() {
    server := http.Server{
        Addr: ":8080",
    }
    http.HandleFunc("/bookmarks/", router.HandleBookmarkRequest)
    server.ListenAndServe()
}
