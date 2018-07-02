package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/hauke96/picl/src/log"
	"github.com/hauke96/picl/src/pkg"
	"github.com/hauke96/picl/src/util"
)

func Install(pkg *pkg.Package, outputBaseFolder string, remoteBaseUrl *url.URL) error {
	log.Info("Start installing...")

	ensureOutputFolder(outputBaseFolder)

	err := downloadMetaFile(remoteBaseUrl.String(), outputBaseFolder, pkg)
	if err != nil {
		return err
	}

	err = downloadPackageFile(remoteBaseUrl.String(), outputBaseFolder, pkg)
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
		os.Mkdir(outputBaseFolder, 0755)
	}
}

func downloadMetaFile(remoteBaseUrl, outputBaseFolder string, pkg *pkg.Package) error {
	versionedPackageName := pkg.VersionedNameString()

	url := fmt.Sprintf("%s/%s/%s", remoteBaseUrl, versionedPackageName, "meta")
	metaFile := fmt.Sprintf("%s/meta_%s", outputBaseFolder, versionedPackageName)

	err := util.DownloadFile(url, metaFile)
	if err != nil {
		return err
	}

	log.Info(fmt.Sprintf("Downloaded meta-file for %s", versionedPackageName))
	log.Debug(fmt.Sprintf("Downloaded meta-file from %s to %s", url, metaFile))
	return nil
}

func downloadPackageFile(remoteBaseUrl, outputBaseFolder string, pkg *pkg.Package) error {
	versionedPackageName := pkg.VersionedNameString()

	// TODO Use file extension form meta file
	url := fmt.Sprintf("%s/%s/%s", remoteBaseUrl, versionedPackageName, pkg.Name)
	metaFile := fmt.Sprintf("%s/%s", outputBaseFolder, versionedPackageName)

	err := util.DownloadFile(url, metaFile)
	if err != nil {
		return err
	}

	log.Info(fmt.Sprintf("Downloaded package-file for %s", versionedPackageName))
	log.Debug(fmt.Sprintf("Downloaded package-file from %s to %s", url, metaFile))
	return nil
}
