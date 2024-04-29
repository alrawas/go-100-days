package readingfiles

import (
	"io/fs"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dirEntry, err := fs.ReadDir(fileSystem, ".")

	if err != nil {
		return nil, err
	}

	var posts []Post
	for range dirEntry {
		posts = append(posts, Post{})
	}
	return posts, nil
}
