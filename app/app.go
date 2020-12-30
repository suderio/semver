package app

import (
	"errors"
	"regexp"
	"strconv"
)

var valid = regexp.MustCompile(`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)

// Semver representa uma versão semântica
type Semver struct {
	Major   int
	Minor   int
	Patch   int
	Release string
	Build   string
}

// ParseSemVer parses a string and returns its semantic versioning parts
func ParseSemVer(semver string) (Semver, error) {
	var result = valid.FindStringSubmatch(semver)
	if result == nil || len(result) != 6 {
		return Semver{}, errors.New("semver not found")
	}
	major, err := strconv.Atoi(result[1])
	if err != nil {
		return Semver{}, errors.New("major version is not a number")
	}
	minor, err := strconv.Atoi(result[2])
	if err != nil {
		return Semver{}, errors.New("minor version is not a number")
	}
	patch, err := strconv.Atoi(result[3])
	if err != nil {
		return Semver{}, errors.New("patch version is not a number")
	}
	release := result[4]

	build := result[5]

	return Semver{major, minor, patch, release, build}, nil
}
