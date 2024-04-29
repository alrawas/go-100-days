package readingfiles_test

import (
	blogposts "alrawas/100daysofgo/reading-files"
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"
)

// making the test file in another package _test so
// we not tend to couple the tests with internal implementation
// and test the exported functions instead

// func Test
func TestNewBlogPosts(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello-world.md": {Data: []byte("hi")},
			"hola-world.md":  {Data: []byte("hola")},
		}

		posts, err := blogposts.NewPostsFromFS(fs)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fs) {
			t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
		}
	})

	t.Run("fs open should fail", func(t *testing.T) {
		_, err := blogposts.NewPostsFromFS(StubFailingFS{})
		if err == nil {
			t.Fatal("fs open should fail, it didn't")
		}
	})

}

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}
