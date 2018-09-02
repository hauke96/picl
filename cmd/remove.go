package cmd

import (
	"fmt"
	"os"

	"github.com/hauke96/sigolo"

	"github.com/hauke96/picl/pkg"
)

func Remove(pkgObj *pkg.Package, outputBaseFolder string) error {
	sigolo.Info("Start removing...")

	// Parse meta file
	metaFilePath := fmt.Sprintf("%s/%s", outputBaseFolder, pkgObj.MetaFileName())

	// TODO check if meta file exists. If not, print error
	metaFileRef, err := os.Open(metaFilePath)
	if err != nil {
		return fmt.Errorf("Error opening meta file %s: %s", metaFilePath, err.Error())
	}

	metaFile, err := pkg.ParseMetaFile(metaFileRef)
	if err != nil {
		return fmt.Errorf("Error reading meta file %s: %s", metaFilePath, err.Error())
	}

	// Determine/build package file path
	packageFile := fmt.Sprintf("%s/%s%s", outputBaseFolder, pkgObj.VersionedNameString(), metaFile.Ext)

	// TODO check if package file exists. If not, print error

	// Actually remove files
	sigolo.Info(fmt.Sprintf("Remove meta-file for %s", pkgObj.VersionedNameString()))
	os.Remove(metaFilePath)

	sigolo.Info(fmt.Sprintf("Remove package-file for %s", pkgObj.VersionedNameString()))
	os.Remove(packageFile)

	sigolo.Info(fmt.Sprintf("Removal of %s finished", pkgObj.VersionedNameString()))
	return nil
}
