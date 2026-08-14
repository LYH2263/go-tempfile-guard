package fsx

import "path/filepath"

func TempBeside(path string) string {
	return filepath.Join(filepath.Dir(path), "."+filepath.Base(path)+".tmp")
}
