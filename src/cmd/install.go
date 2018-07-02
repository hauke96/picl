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

func Install(pkg *pkg.Package, outputBaseFolder string, remoteBaseUrl *url.URL) error {
	log.Info("Start installing...")

	ensureOutputFolder(outputBaseFolder)

	err := downloadMetaFile(remoteBaseUrl.String(), outputBaseFolder, pkg.VersionedNameString())
	if err != nil {
		return err
	}

	log.Info(fmt.Sprintf("Installation of %s finished", pkg.VersionedNameString()))
	return nil
}

// Create the output folder if it doesn't exist
func ensureOutputFolder(outputBaseFolder string) {
	if _, err := os.Stat(outputBaseFolder); os.IsNotExist(err) {
		log.Info(fmt.Sprintf("Output folder %s does not exist. I'll create it...", outputBaseFolder))

		// TODO FIX Permissions on new folder are 000 (instead of e.g. 644)
		os.Mkdir(outputBaseFolder, os.ModeDir)
	}
}

func downloadMetaFile(remoteBaseUrl, outputBaseFolder, versionedPackageName string) error {
	url := fmt.Sprintf("%s/%s/%s", remoteBaseUrl, versionedPackageName, "meta")
	metaFile := fmt.Sprintf("%s/meta_%s", outputBaseFolder, versionedPackageName)

	err := downloadFile(url, metaFile)
	if err != nil {
		return err
	}

	log.Info(fmt.Sprintf("Downloaded meta-file for %s", versionedPackageName))
	log.Debug(fmt.Sprintf("Downloaded meta-file from %s to %s", url, metaFile))
	return nil
}

func downloadFile(url string, fileName string) error {
	var err error = nil

	// Create output file
	var output *os.File = nil

	// Remove existing file
	if _, err := os.Stat(fileName); err == nil {
		log.Debug("Remove existing file")
		err = os.Remove(fileName)
		if err != nil {
			return errors.New(fmt.Sprintf("Error removing file %s: %s", fileName, err.Error()))
		}
	}

	// Create output file
	log.Debug("Create new file")
	output, err = os.Create(fileName)
	if err != nil {
		return errors.New(fmt.Sprintf("Error creating file %s: %s", fileName, err.Error()))
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
