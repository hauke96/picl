package pkg

import (
	"strings"
)

type Package struct {
	Name    string
	Version string
}

func ParsePackage(versionedName string) *Package {
	splittedName := strings.Split(versionedName, "@")

	// TODO check if array is valid

	return &Package{
		Name:    splittedName[0],
		Version: splittedName[1],
	}
}

func (p *Package) String() string {
	return p.Name + "@" + p.Version
}
