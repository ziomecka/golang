package copyfile

import (
	"io"
	"os"
)

func Copy(srcFile, destFile string) (written int64, err error) {
	src, err := os.Open(srcFile)
	if err != nil {
		return
	}

	dest, err := os.Create(destFile)
	defer src.Close()
	if err != nil {
		return
	}

	written, err = io.Copy(dest, src)

	dest.Close()

	return
}
