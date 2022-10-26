package controller

import (
	"net/http"
)

type Router interface {
    HandleBookmarkRequest(w http.ResponseWriter, r *http.Request)
}

type router struct {
    bc BookmarkController
}

func NewRouter(bc BookmarkController) Router {
    return &router{bc}
}

func (ro *router) HandleBookmarkRequest(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
        case "GET":
            ro.bc.HandleBookmarkRequest(w, r,)
        default:
            w.WriteHeader(405)
    }
}
