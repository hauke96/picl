package main

import (
	"math"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/hauke96/sigolo"

	"github.com/hauke96/picl/util"
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
		sigolo.Fatal("Error reading config file: %s", err)
	}

	// TODO extract this parsing routine into own functions used also by meta.go
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
				sigolo.Error("Parsing line %d failed. This could be an error in the regex, the splitting or the line itself", i)
				continue
			}

			key := splittedLine[0]
			value := splittedLine[1]

			pairs[key] = value
		default:
			// To print the first 20 characters of the line, we have to be careful with the bounds of this line
			upperBound := int(math.Min(float64(len(line)), 20))

			sigolo.Error("Malformed config-entry in line %d: %s...", i, line[:upperBound])
		}
	}

	if value, ok := pairs["url"]; ok {
		urlPtr, err := url.Parse(value)

		if err != nil {
			sigolo.Error("Error parsing key 'url' from config")
			// TODO further error handling?
		} else {
			configRemoteUrl = urlPtr
		}
	}

	if value, ok := pairs["output_folder"]; ok {
		configOutputFolder = value
	}
}
