package main

import (
	"fmt"
	"math"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/hauke96/picl/src/util"
)

var (
	configFileCommentRegex   = regexp.MustCompile("^#\\s*\\S*")
	configFileBlankLineRegex = regexp.MustCompile("^\\s*$")
	configFileValidRegex     = regexp.MustCompile("^\\s*\\S+\\s*:\\s*\\S+\\s*$")

	configOutputFolder string
	configRemoteUrl    *url.URL
)

func readConfig(configFile *os.File) {
	lines, err := util.ReadAllLines(configFile)
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
				fmt.Fprintf(os.Stderr, "Parsing line %d failed. This could be an error in the regex, the splitting or the line itself", i)
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
			fmt.Fprintf(os.Stderr, "Error parsing key 'url' from config\n")
			// TODO further error handling?
		} else {
			configRemoteUrl = urlPtr
		}
	}

	if value, ok := pairs["output_folder"]; ok {
		configOutputFolder = value
	}
}
