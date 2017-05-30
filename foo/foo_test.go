package foo

import (
	"os"
	"testing"
)

func TestEmpty(t *testing.T) {
	dir := os.TempDir()
	defer os.RemoveAll(dir)
}
