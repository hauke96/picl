package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/hauke96/picl/src/log"
)

func Install(installPackageName string, outputBaseFolder *os.File, remoteBaseUrl *url.URL) {
	log.Info("Start installing...")

	// Create output dir of not exists
	if _, err := os.Stat(outputBaseFolder.Name()); os.IsNotExist(err) {
		log.Info(fmt.Sprintf("Output folder %s does not exist. I'll create it...", outputBaseFolder.Name()))
		os.Mkdir(outputBaseFolder.Name(), os.ModeDir)
	}
}
