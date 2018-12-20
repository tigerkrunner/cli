package shared

import (
	"fmt"
	"runtime"

	"code.cloudfoundry.org/cli/command"
)

func BuildUserAgent(config command.Config) string {
	return fmt.Sprintf("%s/%s (%s; %s %s)", config.BinaryName(), config.BinaryVersion(), runtime.Version(), runtime.GOARCH, runtime.GOOS)
}
