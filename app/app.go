package app

import (
	"bytes"
	"errors"
	"io"
	"regexp"
	"strconv"
	"text/template"
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

// Show prints the semantic version
func Show(wr io.Writer, semver Semver, verbose bool, linebreak bool) {
	var sepVersion, sepRelease, sepBuild = " ", " ", " "
	var majorStr, minorStr, patchStr, releaseStr, buildStr = "", "", "", "", ""
	if verbose {
		majorStr = "Major: "
		minorStr = "Minor: "
		patchStr = "Patch: "
		releaseStr = "Release: "
		buildStr = "Build: "
	} else {
		sepVersion, sepRelease, sepBuild = ".", "-", "+"
	}
	if linebreak {
		sepVersion, sepRelease, sepBuild = "\n", "\n", "\n"
	}
	var templ = majorStr + "{{ .Major}}" + sepVersion + minorStr + "{{ .Minor}}" + sepVersion + patchStr + "{{ .Patch}}{{if .Release}}" + sepRelease + releaseStr + "{{ .Release}}{{end}}{{if .Build}}" + sepBuild + buildStr + "{{ .Build}}{{end}}\n"

	t, err := template.New("semver").Parse(templ)
	if err != nil {
		panic(err)
	}
	err = t.Execute(wr, semver)
	if err != nil {
		panic(err)
	}
}

// ParseSemVer parses a string and returns its semantic versioning parts
func ParseSemVer(semver string) (Semver, error) {
	var result = valid.FindStringSubmatch(semver)
	if result == nil || len(result) != 6 {
		return Semver{}, errors.New("semver not found")
	}
	major, _ := strconv.Atoi(result[1])
	minor, _ := strconv.Atoi(result[2])
	patch, _ := strconv.Atoi(result[3])
	release := result[4]
	build := result[5]
	return Semver{major, minor, patch, release, build}, nil
}

// Up increases a semantic version
func Up(semver string, major bool, minor bool) (Semver, error) {
	oldVer, err := ParseSemVer(semver)
	if err != nil {
		return Semver{}, err
	}
	if major {
		oldVer.Major++
		oldVer.Minor = 0
		oldVer.Patch = 0
	} else if minor {
		oldVer.Minor++
		oldVer.Patch = 0
	} else {
		oldVer.Patch++
	}
	oldVer.Release = ""
	oldVer.Build = ""
	return oldVer, nil
}

// Down increases a semantic version
func Down(semver string, major bool, minor bool) (Semver, error) {
	oldVer, err := ParseSemVer(semver)
	if err != nil {
		return Semver{}, err
	}
	if major {
		oldVer.Major--
		oldVer.Minor = 0
		oldVer.Patch = 0
	} else if minor {
		oldVer.Minor--
		oldVer.Patch = 0
	} else {
		oldVer.Patch--
	}
	oldVer.Release = ""
	oldVer.Build = ""
	return oldVer, nil
}

// Set sets a semantic version
func Set(semver string, major string, minor string, patch string, release string, build string) (Semver, error) {
	oldVer, err := ParseSemVer(semver)
	if err != nil {
		return Semver{}, err
	}

	if len(major) > 0 {
		oldVer.Major, err = strconv.Atoi(major)
		if err != nil {
			return Semver{}, err
		}
	}
	if len(minor) > 0 {
		oldVer.Minor, err = strconv.Atoi(minor)
		if err != nil {
			return Semver{}, err
		}
	}
	if len(patch) > 0 {
		oldVer.Patch, err = strconv.Atoi(patch)
		if err != nil {
			return Semver{}, err
		}
	}
	if len(major) > 0 {
		oldVer.Release = release
	}
	if len(major) > 0 {

		oldVer.Build = build
	}
	var buf bytes.Buffer
	Show(&buf, oldVer, false, false)
	v, err := ParseSemVer(buf.String())
	if err != nil {
		return Semver{}, err
	}
	return v, nil
}
