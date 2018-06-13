package main

import (
	"fmt"
	"os"

	"github.com/hauke96/kingpin"
)

var (
	app = kingpin.New("picl", "Maybe the dumbest package manager ever")

	install             = app.Command("install", "Installs the given library")
	installConfigFile   = install.Flag("config", "Specifies the configuration file that should be used. This is \"./picl.conf\" by default.").Short('c').Default("./picl.conf").File()
	installOutputFolder = install.Flag("output", "Specifies the output folder where all libraries should be stored. This is \"./libs\" by default.").Short('o').Default("./libs/").File()
	installUrl          = install.Flag("url", "The base url where picl downloads files from").Short('u').URL()
	installPackageName  = install.Arg("package", "The library to install").String()

	remove            = app.Command("remove", "Uninstalls/removes the given library")
	removePackageName = remove.Arg("package", "The library to remove").String()
)

func main() {
	app.Author("Hauke Stieler")
	app.Version("0.1")

	app.CustomDescription("Package Name", `This name if the library name including the version you wan't do deal with. The name has the following format:

      my-library@3.5.1

There must be a name and there must be a version. The version is basically the string that is behind the "@" and is not parsed. It just has to exist on the server, but the format "x.y.z" (e.g. 3.5.1) is only recommended.`)

	app.HelpFlag.Short('h')
	app.VersionFlag.Short('v')

	app.Parse(os.Args[1:])

	fmt.Println(*installPackageName)
}
