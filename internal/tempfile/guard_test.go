package tempfile_test

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/LYH2263/go-tempfile-guard/internal/tempfile"
)

func TestSaveErrorKept(t *testing.T) {
	dir := t.TempDir()
	g := tempfile.New(filepath.Join(dir, "t"), filepath.Join(dir, "f"))
	err := g.Save(func(tmp string) error { return os.ErrPermission })
	if !errors.Is(err, os.ErrPermission) {
		t.Fatalf("%v", err)
	}
}
