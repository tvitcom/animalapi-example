package util

import (
	"os"
)

func PanicError(err error) {
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
