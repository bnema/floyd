package version

import "runtime/debug"

var (
	// Version holds the current version number
	Version string
	// CommitSHA holds the git commit SHA at build time
	CommitSHA string
)

func setVersionInfo() {
	if Version == "" {
		if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "(devel)" {
			Version = info.Main.Version
		} else {
			Version = "dev"
		}
	}
	if CommitSHA == "" {
		CommitSHA = "unknown"
	}
}

func PrintVersion() string {
	setVersionInfo()
	return Version
}
