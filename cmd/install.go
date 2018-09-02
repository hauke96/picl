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

	metaFile, err := downloadMetaFile(remoteBaseUrl.String(), outputBaseFolder, pkg)
	if err != nil {
		return err
	}

	err = downloadPackageFile(remoteBaseUrl.String(), outputBaseFolder, pkg, metaFile)
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

func downloadMetaFile(remoteBaseUrl, outputBaseFolder string, pkgObj *pkg.Package) (*pkg.MetaFile, error) {
	versionedPackageName := pkgObj.VersionedNameString()

	url := fmt.Sprintf("%s/%s/%s", remoteBaseUrl, versionedPackageName, "meta")
	metaFile := fmt.Sprintf("%s/%s", outputBaseFolder, pkgObj.MetaFileName())

	err := util.DownloadFile(url, metaFile)
	if err != nil {
		return nil, err
	}

	sigolo.Info(fmt.Sprintf("Downloaded meta-file for %s", versionedPackageName))
	sigolo.Debug(fmt.Sprintf("Downloaded meta-file from %s to %s", url, metaFile))

	file, err := os.Open(metaFile)
	if err != nil {
		return nil, fmt.Errorf("Could not open meta file %s after downloading", metaFile)
	}

	return pkg.ParseMetaFile(file)
}

func downloadPackageFile(remoteBaseUrl, outputBaseFolder string, pkg *pkg.Package, meta *pkg.MetaFile) error {
	versionedPackageName := pkg.VersionedNameString()

	url := fmt.Sprintf("%s/%s/%s%s", remoteBaseUrl, versionedPackageName, versionedPackageName, meta.Ext)
	packageFile := fmt.Sprintf("%s/%s%s", outputBaseFolder, versionedPackageName, meta.Ext)

	err := util.DownloadFile(url, packageFile)
	if err != nil {
		return err
	}

	sigolo.Info(fmt.Sprintf("Downloaded package-file for %s", versionedPackageName))
	sigolo.Debug(fmt.Sprintf("Downloaded package-file from %s to %s", url, packageFile))
	return nil
}
