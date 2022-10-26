package repository

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/renoinn/sample-api/model/entity"
)

type BookmarkRepository interface {
    GetBookmarks() (bookmarks []entity.BookmarkEntity, err error)
}

type bookmarkRepository struct {
}

func NewBookmarkRepository() BookmarkRepository {
    return &bookmarkRepository{}
}

func (br *bookmarkRepository) GetBookmarks() (bookmarks []entity.BookmarkEntity, err error) {
    bookmarks = []entity.BookmarkEntity{}

    rows, err := Db.
		Query("SELECT id, title, url, note FROM bookmark ORDER BY id DESC")
	if err != nil {
		log.Print(err)
		return
	}

    for rows.Next() {
        bookmark := entity.BookmarkEntity{}
        err = rows.Scan(&bookmark.Id, &bookmark.Title, &bookmark.Url, &bookmark.Note)
        if err != nil {
            log.Print(err)
            return
        }
        bookmarks = append(bookmarks, bookmark)
    }

    return
}
