package cmdutil

import (
	"net/http"

	"github.com/psyllos/rally-cli/pkg/iostreams"
)

// Factory represents a command context
type Factory struct {
	IOStreams  *iostreams.IOStreams
	HttpClient func() (*http.Client, error)
}
