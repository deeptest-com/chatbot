package _comm

import (
	"os"
	"path/filepath"
	"strings"
)

func IsRelease() bool {
	arg1 := strings.ToLower(os.Args[0])
	name := filepath.Base(arg1)

	return strings.Index(name, "___") != 0 && strings.Index(arg1, "go-build") < 0
}
