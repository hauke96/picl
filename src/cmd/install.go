package cmd

import (
	"net/url"
	"os"

	"github.com/hauke96/picl/src/log"
)

func Install(installPackageName string, outputBaseFolder *os.File, remoteBaseUrl *url.URL) {
	log.Info(installPackageName)
	log.Info(outputBaseFolder.Name())
	log.Info(remoteBaseUrl.EscapedPath())
}
