package dogu

import (
	"context"
	"errors"
	"github.com/cloudogu/cesapp-lib/core"
)

var DoguDescriptorNotFoundError = errors.New("No DoguDescriptor found for that dogu")
var ConnectionError = errors.New("There are some connection issues")

type RemoteDoguDescriptorRepository interface {
	// GetLatest returns the dogu descriptor for a dogu from the remote server.
	// DoguDescriptorNotFoundError if there is no descriptor for that dogu
	// ConnectionError if there are any connection issues
	// Generic Error if there are any other issues
	GetLatest(context.Context, QualifiedDoguName) (*core.Dogu, error)
	// Get returns a version specific dogu descriptor.
	// DoguDescriptorNotFoundError if there is no descriptor for that dogu
	// ConnectionError if there are any connection issues
	// Generic Error if there are any other issues
	Get(context.Context, QualifiedDoguVersion) (*core.Dogu, error)
}
