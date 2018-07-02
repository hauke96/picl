package cmd

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/hauke96/picl/src/log"
	"github.com/hauke96/picl/src/pkg"
)

func Install(pkg *pkg.Package, outputBaseFolder string, remoteBaseUrl *url.URL) {
	log.Info("Start installing...")

	ensureOutputFolder(outputBaseFolder)

	downloadMetaFile(remoteBaseUrl.String(), outputBaseFolder, pkg.VersionedNameString())
}

// Create the output folder if it doesn't exist
func ensureOutputFolder(outputBaseFolder string) {
	if _, err := os.Stat(outputBaseFolder); os.IsNotExist(err) {
		log.Info(fmt.Sprintf("Output folder %s does not exist. I'll create it...", outputBaseFolder))

		// TODO FIX Permissions on new folder are 000 (instead of e.g. 644)
		os.Mkdir(outputBaseFolder, os.ModeDir)
	}
}

func downloadMetaFile(remoteBaseUrl, outputBaseFolder, versionedPackageName string) {
	url := fmt.Sprintf("%s/%s/%s", remoteBaseUrl, versionedPackageName, "meta")
	metaFile := fmt.Sprintf("%s/meta_%s", outputBaseFolder, versionedPackageName)

	log.Info(fmt.Sprintf("Download meta-file for %s", versionedPackageName))
	log.Debug(fmt.Sprintf("Download meta-file from %s to %s", url, metaFile))
	downloadFile(url, metaFile)
}

func downloadFile(url string, fileName string) error {
	// Create output file
	// TODO: check file existence first with io.IsExist
	output, err := os.Create(fileName)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while creating %s: %s", fileName, err.Error()))
	}
	defer output.Close()

	// Donwload data
	response, err := http.Get(url)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while donwloading %s: %s", url, err.Error()))
	}
	defer response.Body.Close()

	// Write data into output File
	_, err = io.Copy(output, response.Body)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while copying response stream from %s to %s: %s", url, fileName, err.Error()))
	}

	return nil
}
