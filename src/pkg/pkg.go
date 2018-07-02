package pkg

import (
	"fmt"
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
	return p.VersionedNameString()
}

// Create the raw string for this package, which is "Name@Version" (so e.g.
// foo@1.2.3).
func (p *Package) VersionedNameString() string {
	return p.Name + "@" + p.Version
}
