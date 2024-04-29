package readingfiles_test

import (
	blogposts "alrawas/100daysofgo/reading-files"
	"testing"
	"testing/fstest"
)

// making the test file in another package _test so
// we not tend to couple the tests with internal implementation
// and test the exported functions instead

// func Test
func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello-world.md": {Data: []byte("hi")},
		"hola-world.md":  {Data: []byte("hola")},
	}

	posts := blogposts.NewPostsFromFS(fs)

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}
}
