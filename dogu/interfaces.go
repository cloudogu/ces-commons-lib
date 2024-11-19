package dogu

import (
	"context"
	"github.com/cloudogu/cesapp-lib/core"
)

type RemoteDoguDescriptorRepository interface {
	// GetLatest returns the dogu descriptor for a dogu from the remote server.
	// DoguDescriptorNotFoundError if there is no descriptor for that dogu
	// ConnectionError if there are any connection issues
	// Generic Error if there are any other issues
	GetLatest(context.Context, QualifiedName) (*core.Dogu, error)
	// Get returns a version specific dogu descriptor.
	// DoguDescriptorNotFoundError if there is no descriptor for that dogu
	// ConnectionError if there are any connection issues
	// Generic Error if there are any other issues
	Get(context.Context, QualifiedVersion) (*core.Dogu, error)
}

type VersionRegistry interface {
	// GetCurrent returns the current installed dogu version.
	// NotFound Error if the dogu is not installed.
	// ConnectionError if there are any connection issues.
	// Generic Error if there are any other issues.
	GetCurrent(context.Context, SimpleName) (SimpleNameVersion, error)
	// GetCurrentOfAll returns all current installed dogu versions.
	// ConnectionError if there are any connection issues.
	// Generic Error if there are any other issues.
	GetCurrentOfAll(context.Context) ([]SimpleNameVersion, error)
	// IsEnabled returns true if the dogu version is installed.
	// False otherwise.
	// ConnectionError if there are any connection issues
	// Generic Error if there are any other issues.
	IsEnabled(context.Context, SimpleNameVersion) (bool, error)
	// Enable the dogu version as current. It returns
	// ConnectionError if there are any connection issues
	// Generic Error if there are any other issues.
	Enable(context.Context, SimpleNameVersion) error
	// WatchAllCurrent watches the version registry and notifies via the returned channel if changes happen.
	// ConnectionError if there are any connection issues
	// Generic Error if there are any other issues.
	WatchAllCurrent(context.Context) (<-chan CurrentVersionsWatchResult, error)
}

// LocalDoguDescriptorRepository is an append-only Repository, no updates will happen
type LocalDoguDescriptorRepository interface {
	// Get returns the dogu descriptor from the local registry.
	// ConnectionError if there are any connection issues
	// Generic Error if there are any other issues.
	Get(context.Context, SimpleNameVersion) (*core.Dogu, error)
	// GetAll returns all dogu descriptor by parameter from the local registry.
	// ConnectionError if there are any connection issues
	// Generic Error if there are any other issues.
	GetAll(context.Context, []SimpleNameVersion) (map[SimpleNameVersion]*core.Dogu, error)
	// Add adds a dogu descriptor to the local registry.
	// ConnectionError if there are any connection issues
	// Generic Error if there are any other issues.
	Add(context.Context, SimpleName, *core.Dogu) error
	// DeleteAll deletes all dogu descriptors for a given dogu name.
	// ConnectionError if there are any connection issues
	// Generic Error if there are any other issues.
	DeleteAll(context.Context, SimpleName) error
}
