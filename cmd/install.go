package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/hauke96/picl/pkg"
	"github.com/hauke96/picl/util"
	"github.com/hauke96/sigolo"
)

func Install(pkg *pkg.Package, outputBaseFolder string, remoteBaseUrl *url.URL) error {
	sigolo.Info("Start installing...")

	ensureOutputFolder(outputBaseFolder)

	err := downloadMetaFile(remoteBaseUrl.String(), outputBaseFolder, pkg)
	if err != nil {
		return err
	}

	// TODO Parse meta file here. First extract the read-functionality from
	// configs into the util (or other) package, because configs and meta
	// files have the same syntax. Maybe adjust and use the config.go file
	// to read the meta file?

	err = downloadPackageFile(remoteBaseUrl.String(), outputBaseFolder, pkg)
	if err != nil {
		return err
	}

	sigolo.Info(fmt.Sprintf("Installation of %s finished", pkg.VersionedNameString()))
	return nil
}

// Create the output folder if it doesn't exist
func ensureOutputFolder(outputBaseFolder string) {
	if _, err := os.Stat(outputBaseFolder); os.IsNotExist(err) {
		sigolo.Info(fmt.Sprintf("Output folder %s does not exist. I'll create it...", outputBaseFolder))

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

	sigolo.Info(fmt.Sprintf("Downloaded meta-file for %s", versionedPackageName))
	sigolo.Debug(fmt.Sprintf("Downloaded meta-file from %s to %s", url, metaFile))
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

	sigolo.Info(fmt.Sprintf("Downloaded package-file for %s", versionedPackageName))
	sigolo.Debug(fmt.Sprintf("Downloaded package-file from %s to %s", url, metaFile))
	return nil
}
