package dogu

import (
	"fmt"
	"github.com/cloudogu/cesapp-lib/core"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQualifiedDoguNameFromString(t *testing.T) {
	tests := []struct {
		test     string
		given    string
		expected QualifiedDoguName
		wantErr  assert.ErrorAssertionFunc
	}{
		{test: "ok", given: "official/postgres", expected: QualifiedDoguName{SimpleDoguName("postgres"), DoguNamespace("official")}, wantErr: assert.NoError},
		{test: "no ns", given: "postgres", expected: QualifiedDoguName{}, wantErr: assert.Error},
		{test: "no name", given: "official/", expected: QualifiedDoguName{}, wantErr: assert.Error},
		{test: "double namespace", given: "official/test/postgres", expected: QualifiedDoguName{}, wantErr: assert.Error},
	}
	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			got, err := QualifiedDoguNameFromString(tt.given)
			if !tt.wantErr(t, err, fmt.Sprintf("TestQualifiedDoguNameFromString(%v)", tt.given)) {
				return
			}
			assert.Equalf(t, tt.expected, got, "TestQualifiedDoguNameFromString(%v)", tt.given)
		})
	}
}

func TestNewQualifiedDoguVersion(t *testing.T) {
	tests := []struct {
		name          string
		qualifiedDogu QualifiedDoguName
		version       core.Version
		want          QualifiedDoguVersion
	}{
		{
			name:          "create QualifiedDoguVersion",
			qualifiedDogu: QualifiedDoguName{SimpleName: "postgres", Namespace: "official"},
			version:       core.Version{Raw: "1.0.0"},
			want:          QualifiedDoguVersion{Name: QualifiedDoguName{SimpleName: "postgres", Namespace: "official"}, Version: core.Version{Raw: "1.0.0"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewQualifiedDoguVersion(tt.qualifiedDogu, tt.version), "NewQualifiedDoguVersion(%v, %v)", tt.qualifiedDogu, tt.version)
		})
	}
}
