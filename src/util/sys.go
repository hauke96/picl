package util

import (
	"os"

	"github.com/hauke96/picl/src/log"
)

func ExitOnError(err error) {
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}
