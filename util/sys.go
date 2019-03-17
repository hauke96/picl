package util

import (
	"github.com/hauke96/sigolo"
)

func ExitOnError(err error) {
	if err != nil {
		sigolo.Fatal(err.Error())
	}
}
