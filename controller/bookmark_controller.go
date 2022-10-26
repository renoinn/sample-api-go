package controller

import (
	"encoding/json"
	"net/http"

	"github.com/renoinn/sample-api/model/repository"
	"github.com/renoinn/sample-api/model/response"
)

type BookmarkController interface {
    HandleBookmarkRequest(w http.ResponseWriter, r *http.Request)
}

type bookmarkController struct {
    br repository.BookmarkRepository
}

func NewController(br repository.BookmarkRepository) Router {
    return &bookmarkController{br}
}

func (bc *bookmarkController) HandleBookmarkRequest(w http.ResponseWriter, r *http.Request) {
    bookmarks, err := bc.br.GetBookmarks()
    if err != nil {
        w.WriteHeader(500)
        return
    }
    
    var bookmarkResponse []response.BookmarkResponse
    for _, v := range bookmarks {
        bookmarkResponse = append(bookmarkResponse, response.BookmarkResponse{Id: v.Id, Title: v.Title, Url: v.Url, Note: v.Note})
    }
    var bookmarksResponse response.BookmarksResponse
    bookmarksResponse.Bookmarks = bookmarkResponse

    output, _ := json.MarshalIndent(bookmarksResponse.Bookmarks, "", "\t\t")
    w.Header().Set("Contet-Type", "application/json")
    w.Write(output)
}
