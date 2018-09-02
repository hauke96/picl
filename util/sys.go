package util

import (
	"os"

	"github.com/hauke96/sigolo"
)

func ExitOnError(err error) {
	if err != nil {
		sigolo.Error(err.Error())
		os.Exit(1)
	}
}
