package readingfiles_test

import (
	blogposts "alrawas/100daysofgo/reading-files"
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

// making the test file in another package _test so
// we not tend to couple the tests with internal implementation
// and test the exported functions instead

func TestNewBlogPosts(t *testing.T) {

	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker`
	)

	t.Run("happy path", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello-world.md": {Data: []byte(firstBody)},
			"hola-world.md":  {Data: []byte(secondBody)},
		}

		posts, err := blogposts.NewPostsFromFS(fs)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fs) {
			t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
		}

		assertPost(t, posts[0], blogposts.Post{Title: "Post 1", Description: "Description 1", Tags: []string{"tdd", "go"}})
	})

	t.Run("fs open should fail", func(t *testing.T) {
		_, err := blogposts.NewPostsFromFS(StubFailingFS{})
		if err == nil {
			t.Fatal("fs open should fail, it didn't")
		}
	})

}

func assertPost(t *testing.T, got, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}
