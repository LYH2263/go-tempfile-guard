package writer

import (
	"os"

	"github.com/LYH2263/go-tempfile-guard/internal/fsx"
	"github.com/LYH2263/go-tempfile-guard/internal/tempfile"
)

func AtomicWrite(path string, data []byte) error {
	g := tempfile.New(fsx.TempBeside(path), path)
	return g.Save(func(tmp string) error {
		return os.WriteFile(tmp, data, 0o644)
	})
}

func AtomicWriteFail(path string) error {
	g := tempfile.New(fsx.TempBeside(path), path)
	return g.Save(func(tmp string) error {
		return os.ErrPermission
	})
}
