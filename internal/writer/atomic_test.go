package writer_test

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/LYH2263/go-tempfile-guard/internal/writer"
)

func TestFailPreservesError(t *testing.T) {
	p := filepath.Join(t.TempDir(), "x.txt")
	err := writer.AtomicWriteFail(p)
	if !errors.Is(err, os.ErrPermission) {
		t.Fatalf("want permission, got %v", err)
	}
}

func TestOK(t *testing.T) {
	p := filepath.Join(t.TempDir(), "y.txt")
	if err := writer.AtomicWrite(p, []byte("ok")); err != nil {
		t.Fatal(err)
	}
}
