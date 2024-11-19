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
	GetCurrent(context.Context, SimpleName) (SimpleNameVersion, error)
	GetCurrentOfAll(context.Context) ([]SimpleNameVersion, error)
	IsEnabled(context.Context, SimpleNameVersion) (bool, error)
	Enable(context.Context, SimpleNameVersion) error
	WatchAllCurrent(context.Context) (<-chan CurrentVersionsWatchResult, error)
}

// LocalDoguDescriptorRepository is an append-only Repository, no updates will happen
type LocalDoguDescriptorRepository interface {
	Get(context.Context, SimpleNameVersion) (*core.Dogu, error)
	GetAll(context.Context, []SimpleNameVersion) (map[SimpleNameVersion]*core.Dogu, error)
	Add(context.Context, SimpleName, *core.Dogu) error
	DeleteAll(context.Context, SimpleName) error
}
