package document

import (
	"testing"
	"log"
)

func TestLocalFiles_Get(t *testing.T) {
	l := LocalFiles{"."}
	if err := l.Get(); err != nil {
		log.Fatalf("Error running Get: %s", err)
	}
}

func TestLocalFiles_Path(t *testing.T) {
	exp := "."
	l := LocalFiles{WorkDir: exp}
	r := l.Path()
	if r != "." {
		log.Fatalf("Expected %q, got %s", r, exp)
	}
}

func TestLocalFiles_CleanUp(t *testing.T) {
	l := LocalFiles{"."}
	l.CleanUp()
}