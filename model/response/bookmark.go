package response

type BookmarkResponse struct {
    Id int `json:"id"`
    Title string `json:"title"`
    Url string `json:"url"`
    Note string `json:"note"`
}

type BookmarksResponse struct {
    Bookmarks []BookmarkResponse `json:"bookmarks"`
}
