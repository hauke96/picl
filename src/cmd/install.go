package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/hauke96/picl/src/log"
	"github.com/hauke96/picl/src/pkg"
)

func Install(pkg *pkg.Package, outputBaseFolder string, remoteBaseUrl *url.URL) {
	log.Info("Start installing...")

	ensureOutputFolder(outputBaseFolder)
}

// Create the output folder if it doesn't exist
func ensureOutputFolder(outputBaseFolder string) {
	if _, err := os.Stat(outputBaseFolder); os.IsNotExist(err) {
		log.Info(fmt.Sprintf("Output folder %s does not exist. I'll create it...", outputBaseFolder))

		// TODO FIX Permissions on new folder are 000 (instead of e.g. 644)
		os.Mkdir(outputBaseFolder, os.ModeDir)
	}
}
