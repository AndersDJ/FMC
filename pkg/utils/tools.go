package utils

import (
	"fmt"

	"github.com/spf13/afero"
)

func CountSize(n int) string {
	if n < 1024 {
		//return strconv.FormatInt(fileSize, 10) + "B"
		return fmt.Sprintf("%.2f B", float64(n)/float64(1))
	} else if n < (1024 * 1024) {
		return fmt.Sprintf("%.2f KB", float64(n)/float64(1024))
	} else if n < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2f MB", float64(n)/float64(1024*1024))
	} else if n < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2f GB", float64(n)/float64(1024*1024*1024))
	} else if n < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2f TB", float64(n)/float64(1024*1024*1024*1024))
	} else { //if fileSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fmt.Sprintf("%.2f EB", float64(n)/float64(1024*1024*1024*1024*1024))
	}
}

func Exists(path string) (bool, error) {
	a := afero.NewOsFs()
	b, err := afero.Exists(a, path)
	if err != nil {
		return b, err
	}
	return b, nil
}
