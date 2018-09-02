package pkg

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	validStringRegex = regexp.MustCompile("^\\S+@\\S+$")
)

type Package struct {
	Name    string
	Version string
}

func IsValid(versionedName string) bool {
	return validStringRegex.MatchString(versionedName)
}

func ParsePackage(versionedName string) (*Package, error) {
	if !IsValid(versionedName) {
		return nil, errors.New("Invalid package name string '" + versionedName + "'")
	}

	splittedName := strings.Split(versionedName, "@")

	return &Package{
		Name:    splittedName[0],
		Version: splittedName[1],
	}, nil
}

func (p *Package) String() string {
	return p.VersionedNameString()
}

// Create the raw string for this package, which is "Name@Version" (so e.g.
// foo@1.2.3).
func (p *Package) VersionedNameString() string {
	return p.Name + "@" + p.Version
}

func (p *Package) MetaFileName() string {
	return fmt.Sprintf("meta_%s", p.VersionedNameString())
}
