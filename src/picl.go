package main

import (
	"fmt"
	"os"

	"github.com/hauke96/kingpin"
)

var (
	app = kingpin.New("picl", "Maybe the dumbest package manager ever")

	installCmd          = app.Command("install", "Installs the given library")
	installConfigFile   = installCmd.Flag("config", "Specifies the configuration file that should be used. This is \"./picl.conf\" by default.").Short('c').Default("./picl.conf").File()
	installOutputFolder = installCmd.Flag("output", "Specifies the output folder where all libraries should be stored. This is \"./libs\" by default.").Short('o').Default("./libs/").File()
	installUrl          = installCmd.Flag("url", "The base url where picl downloads files from").Short('u').URL()
	installPackageName  = installCmd.Arg("package", "The library to install").String()

	removeCmd         = app.Command("remove", "Uninstalls/removes the given library")
	removePackageName = removeCmd.Arg("package", "The library to remove").String()
)

func main() {
	app.Author("Hauke Stieler")
	app.Version("0.1")

	app.CustomDescription("Package Name", `This name if the library name including the version you wan't do deal with. The name has the following format:

      my-library@3.5.1

There must be a name and there must be a version. The version is basically the string that is behind the "@" and is not parsed. It just has to exist on the server, but the format "x.y.z" (e.g. 3.5.1) is only recommended.`)

	app.HelpFlag.Short('h')
	app.VersionFlag.Short('v')

	cmd, err := app.Parse(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while parsing arguments:\n%s\n", err)
		os.Exit(1)
	}

	switch cmd {
	case installCmd.FullCommand():
		fmt.Errorf("Not implemented yet\n")
	case removeCmd.FullCommand():
		fmt.Errorf("Not implemented yet\n")
	}
}
