package context

import (
	"github.com/psyllos/rally-cli/pkg/client"
)

type CmdContext struct {
	Version   string
	BuildDate string
	Client    *client.Client
}

func NewCmdContext(rallyAPI *client.Client, version, buildDate string) *CmdContext {
	return &CmdContext{
		Version:   version,
		BuildDate: buildDate,
		Client:    rallyAPI,
	}
}
