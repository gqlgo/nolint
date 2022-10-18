package nolint

import _ "embed"

//go:embed versionfile
var version string

// Version returns version of nolint.
func Version() string {
	return version
}
