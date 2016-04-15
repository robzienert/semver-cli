package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kingpin"
	"github.com/coreos/go-semver/semver"
)

var (
	app              = kingpin.New("semver", "A command-line utility for managing semver")
	metadata         = app.Flag("metadata", "Metadata version information").Short('m').String()
	pre              = app.Flag("prerelease", "Prerelease version information").Short('p').String()
	bumpMajor        = app.Command("bumpMajor", "Bump the major version of a semver string")
	bumpMajorVersion = bumpMajor.Arg("version", "The version to bump").String()
	bumpMinor        = app.Command("bumpMinor", "Bump the minor version of a semver string")
	bumpMinorVersion = bumpMinor.Arg("version", "The version to bump").String()
	bumpPatch        = app.Command("bumpPatch", "Bump the patch version of a semver string")
	bumpPatchVersion = bumpPatch.Arg("version", "The version to bump").String()
)

func main() {
	var version *semver.Version
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case bumpMajor.FullCommand():
		v, err := semver.NewVersion(*bumpMajorVersion)
		app.FatalIfError(err, "could not parse version")
		version = v
		version.BumpMajor()

	case bumpMinor.FullCommand():
		v, err := semver.NewVersion(*bumpMinorVersion)
		app.FatalIfError(err, "could not parse version")
		version = v
		version.BumpMinor()

	case bumpPatch.FullCommand():
		v, err := semver.NewVersion(*bumpPatchVersion)
		app.FatalIfError(err, "could not parse version")
		version = v
		version.BumpPatch()
	}

	applyMetadata(version, *metadata)
	applyPreRelease(version, *pre)

	fmt.Println(version.String())
}

func applyMetadata(v *semver.Version, m string) {
	if m != "" {
		v.Metadata = m
	}
}

func applyPreRelease(v *semver.Version, p string) {
	if p != "" {
		v.PreRelease = semver.PreRelease(p)
	}
}
