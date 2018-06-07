package main

import (
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
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

	app.HelpFlag.Short('h')
	app.VersionFlag.Short('v')

	app.Parse(os.Args[1:])

	fmt.Println(*installPackageName)
}
