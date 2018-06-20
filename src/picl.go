package main

import (
	"bufio"
	"fmt"
	"math"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/hauke96/kingpin"
	"github.com/hauke96/picl/src/cmd"
)

var (
	app           = kingpin.New("picl", "Maybe the dumbest package manager ever")
	appConfigFile = app.Flag("config", "Specifies the configuration file that should be used. This is \"./picl.conf\" by default.").Short('c').Default("./picl.conf").File()

	installCmd          = app.Command("install", "Installs the given library")
	installOutputFolder = installCmd.Flag("output", "Specifies the output folder where all libraries should be stored.").Short('o').File()
	installUrl          = installCmd.Flag("url", "The base url where picl downloads files from").Short('u').URL()
	installPackageName  = installCmd.Arg("package", "The library to install").String()

	removeCmd         = app.Command("remove", "Uninstalls/removes the given library")
	removePackageName = removeCmd.Arg("package", "The library to remove").String()

	configFileCommentRegex   = regexp.MustCompile("^#\\s*\\S*")
	configFileBlankLineRegex = regexp.MustCompile("^\\s*$")
	configFileValidRegex     = regexp.MustCompile("^\\s*\\S+\\s*:\\s*\\S+\\s*$")

	configOutputFolder *os.File
	configRemoteUrl    *url.URL
)

func readConfig(configFile *os.File) {
	lines, err := readFile(configFile)
	if err != nil {
		panic(err.Error())
	}

	pairs := make(map[string]string)

	for i, line := range lines {
		switch {
		case configFileBlankLineRegex.MatchString(line):
			continue
		case configFileCommentRegex.MatchString(line):
			continue
		case configFileValidRegex.MatchString(line):
			splittedLine := strings.SplitN(line, ":", 2)

			if len(splittedLine) != 2 {
				fmt.Errorf("Parsing line %d failed. This could be an error in the regex, the splitting or the line itself", i)
				continue
			}

			key := splittedLine[0]
			value := splittedLine[1]

			pairs[key] = value
		default:
			// To print the first 20 characters of the line, we have to be careful with the bounds of this line
			upperBound := int(math.Min(float64(len(line)), 20))

			fmt.Fprintf(os.Stderr, "Malformed config-entry in line %d: %s...\n", i, line[:upperBound])
		}
	}

	if value, ok := pairs["url"]; ok {
		urlPtr, err := url.Parse(value)

		if err != nil {
			fmt.Errorf("Error parsing key 'url' from config\n")
			// TODO further error handling?
		} else {
			configRemoteUrl = urlPtr
		}
	}

	if value, ok := pairs["output_folder"]; ok {
		filePtr, err := os.Open(value)

		if err != nil {
			fmt.Errorf("Error parsing key 'output_folder' from config\n")
			// TODO further error handling?
		} else {
			configOutputFolder = filePtr
		}
	}
}

func readFile(file *os.File) ([]string, error) {
	lines := make([]string, 0)

	// defer closing
	defer file.Close()

	// read lines
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// return lines or error
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// When a field (e.g. configOutputFolder) is not set yet, but has been specified
// via a command line argument, it'll be set here. Here we also overwrite
// existing values (e.g. the remote URL).
func setConfigFromArguments() {
	if *installOutputFolder != nil {
		configOutputFolder = *installOutputFolder
	}

	if *installUrl != nil {
		configRemoteUrl = *installUrl
	}
}

func main() {
	app.Author("Hauke Stieler")
	app.Version("0.1")

	app.CustomDescription("Package Name", `This name if the library name including the version you wan't do deal with. The name has the following format:

      my-library@3.5.1

There must be a name and there must be a version. The version is basically the string that is behind the "@" and is not parsed. It just has to exist on the server, but the format "x.y.z" (e.g. 3.5.1) is only recommended.`)

	app.HelpFlag.Short('h')
	app.VersionFlag.Short('v')

	command, err := app.Parse(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while parsing arguments:\n%s\n", err)
		os.Exit(1)
	}

	readConfig(*appConfigFile)

	setConfigFromArguments()

	// TODO handle invalid/not set configurations

	switch command {
	case installCmd.FullCommand():
		cmd.Install(*installPackageName, configOutputFolder, configRemoteUrl)
	case removeCmd.FullCommand():
		fmt.Errorf("Not implemented yet\n")
	}
}
