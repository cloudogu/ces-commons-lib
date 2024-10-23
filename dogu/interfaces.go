package dogu

import (
	"github.com/cloudogu/cesapp-lib/core"
)

type SimpleDoguName string
type DoguNamespace string

type QualifiedDoguVersion struct {
	Name    QualifiedDoguName
	Version core.Version
}
type QualifiedDoguName struct {
	SimpleName SimpleDoguName
	Namespace  DoguNamespace
}
