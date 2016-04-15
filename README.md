# semver-cli

A command line utility for managing semantic versions.

# Install

```
$ go get github.com/robzienert/semver-cli
```

# Usage

```
usage: semver [<flags>] <command> [<args> ...]

A command-line utility for managing semver

Flags:
      --help                   Show context-sensitive help (also try --help-long and --help-man).
  -m, --metadata=METADATA      Metadata version information
  -p, --prerelease=PRERELEASE  Prerelease version information

Commands:
  help [<command>...]
    Show help.

  bumpMajor [<version>]
    Bump the major version of a semver string

  bumpMinor [<version>]
    Bump the minor version of a semver string

  bumpPatch [<version>]
    Bump the patch version of a semver string
```
