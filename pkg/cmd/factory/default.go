package factory

import "github.com/psyllos/rally-cli/pkg/cmdutil"

func New(buildVersion string) *cmdutil.Factory {
	return &cmdutil.Factory{}
}
