package utils

import (
	"fmt"
)

func CountSize(n int) string {
	// log.Infof("n: %d  nf: %f", n, nf)

	// var num float64
	// var unit string
	// if n < 1024 {
	// 	unit = "KB"
	// 	num = nf
	// } else if 1024 <= n && n < (1024*1024) {
	// 	unit = "MB"
	// 	num = nf / 1024
	// } else if (1024*1024) <= n && n < (1024*1024*1024) {
	// 	unit = "GB"
	// 	num = nf / (1024 * 1024)
	// } else if (1024 * 1024 * 1024) <= n {
	// 	unit = "TB"
	// 	num = nf / (1024 * 1024 * 1024)
	// }
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

	// log.Infof("final: %s", fmt.Sprintf("%.2f%s", math.Trunc(num*1e2+0.5)*1e-2, unit))

	// return fmt.Sprintf("%f%s", math.Trunc(num), unit)

}
