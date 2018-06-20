package cmd

import (
	"fmt"
	"net/url"
	"os"
)

func Install(installPackageName string, outputBaseFolder *os.File, remoteBaseUrl *url.URL) {
	fmt.Println(installPackageName)
	fmt.Println(outputBaseFolder.Name())
	fmt.Println(remoteBaseUrl)
}
