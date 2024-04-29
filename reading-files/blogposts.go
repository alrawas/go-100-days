package readingfiles

import (
	"io/fs"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func NewPostsFromFS(fileSystem fs.FS) []Post {
	return []Post{{}, {}}
}
