package pkg

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"

	"github.com/hauke96/picl/util"
)

var (
	configFileCommentRegex   = regexp.MustCompile("^#\\s*\\S*")
	configFileBlankLineRegex = regexp.MustCompile("^\\s*$")
	configFileValidRegex     = regexp.MustCompile("^\\s*\\S+\\s*:\\s*\\S+\\s*$")
)

type MetaFile struct {
	Ext string
}

func ParseMetaFile(file *os.File) (*MetaFile, error) {
	lines, err := util.ReadAllLines(file)
	if err != nil {
		return nil, err
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

	if value, ok := pairs["ext"]; ok {
		meta := MetaFile{
			Ext: value,
		}

		return &meta, nil
	} else {
		return nil, fmt.Errorf("Could not find 'ext' in meta file %s", file)
	}
}
