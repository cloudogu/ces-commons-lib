package dogu

import (
	"fmt"
	"github.com/cloudogu/cesapp-lib/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQualifiedNameFromString(t *testing.T) {
	tests := []struct {
		name          string
		qualifiedName string
		want          QualifiedName
		wantErr       assert.ErrorAssertionFunc
	}{
		{
			name:          "ok",
			qualifiedName: "official/postgres",
			want:          QualifiedName{SimpleName("postgres"), Namespace("official")},
			wantErr:       assert.NoError,
		},
		{
			name:          "no ns",
			qualifiedName: "postgres",
			want:          QualifiedName{},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorContains(t, err, "dogu name needs to be in the form 'namespace/dogu' but is 'postgres'")
			},
		},
		{
			name:          "no name",
			qualifiedName: "official/",
			want:          QualifiedName{},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorContains(t, err, "dogu name must not be empty")
			},
		},
		{
			name:          "double ns",
			qualifiedName: "official/test/postgres",
			want:          QualifiedName{},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorContains(t, err, "dogu name needs to be in the form 'namespace/dogu' but is 'official/test/postgres'")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QualifiedNameFromString(tt.qualifiedName)
			if !tt.wantErr(t, err, fmt.Sprintf("QualifiedNameFromString(%v)", tt.qualifiedName)) {
				return
			}
			assert.Equalf(t, tt.want, got, "QualifiedNameFromString(%v)", tt.qualifiedName)
		})
	}
}

func TestNewQualifiedDoguVersion(t *testing.T) {
	tests := []struct {
		name          string
		qualifiedDogu QualifiedName
		version       func(t *testing.T) core.Version
		want          func(t *testing.T) QualifiedVersion
	}{
		{
			name:          "create QualifiedVersion",
			qualifiedDogu: QualifiedName{SimpleName: "postgres", Namespace: "official"},
			version: func(t *testing.T) core.Version {
				return core.Version{Raw: "1.0.0"}
			},
			want: func(t *testing.T) QualifiedVersion {
				return QualifiedVersion{Name: QualifiedName{SimpleName: "postgres", Namespace: "official"}, Version: core.Version{Raw: "1.0.0"}}
			},
		},
		{
			name:          "create QualifiedVersion with Parse",
			qualifiedDogu: QualifiedName{SimpleName: "postgres", Namespace: "official"},
			version: func(t *testing.T) core.Version {
				version, err := core.ParseVersion("1.2.3")
				require.NoError(t, err)
				return version
			},
			want: func(t *testing.T) QualifiedVersion {
				version, err := core.ParseVersion("1.2.3")
				require.NoError(t, err)
				return QualifiedVersion{Name: QualifiedName{SimpleName: "postgres", Namespace: "official"}, Version: version}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want(t), NewQualifiedVersion(tt.qualifiedDogu, tt.version(t)), "NewQualifiedVersion(%v, %v)", tt.qualifiedDogu, tt.version(t))
		})
	}
}
