package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/hauke96/picl/src/log"
)

func Install(installPackageName string, outputBaseFolder string, remoteBaseUrl *url.URL) {
	log.Info("Start installing...")

	// Create output dir of not exists
	if _, err := os.Stat(outputBaseFolder); os.IsNotExist(err) {
		log.Info(fmt.Sprintf("Output folder %s does not exist. I'll create it...", outputBaseFolder))
		os.Mkdir(outputBaseFolder, os.ModeDir)
	}
}
