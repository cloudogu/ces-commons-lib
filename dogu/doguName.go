package dogu

import (
	"errors"
	"fmt"
	"github.com/cloudogu/cesapp-lib/core"
	"strings"
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

func NewQualifiedDoguName(namespace DoguNamespace, simpleName SimpleDoguName) (QualifiedDoguName, error) {
	doguName := QualifiedDoguName{Namespace: namespace, SimpleName: simpleName}
	err := doguName.Validate()
	if err != nil {
		return QualifiedDoguName{}, err
	}
	return QualifiedDoguName{Namespace: namespace, SimpleName: simpleName}, nil
}

func NewQualifiedDoguVersion(name QualifiedDoguName, version core.Version) QualifiedDoguVersion {
	return QualifiedDoguVersion{
		Name:    name,
		Version: version,
	}
}

func (doguName QualifiedDoguName) Validate() error {
	var errorList []error
	if doguName.Namespace == "" {
		errorList = append(errorList, fmt.Errorf("namespace of dogu %q must not be empty", doguName.SimpleName))
	}
	if doguName.SimpleName == "" {
		errorList = append(errorList, fmt.Errorf("dogu name must not be empty: '%s/%s'", doguName.SimpleName, doguName.Namespace))
	}
	return errors.Join(errorList...)
}

// String returns the dogu name with namespace, e.g. official/postgresql
func (doguName QualifiedDoguName) String() string {
	return fmt.Sprintf("%s/%s", doguName.Namespace, doguName.SimpleName)
}

// QualifiedDoguNameFromString converts a qualified dogu as a string, e.g. "official/nginx", to a dedicated QualifiedDoguName or raises an error if this is not possible.
func QualifiedDoguNameFromString(qualifiedName string) (QualifiedDoguName, error) {
	splitName := strings.Split(qualifiedName, "/")
	if len(splitName) != 2 {
		return QualifiedDoguName{}, fmt.Errorf("dogu name needs to be in the form 'namespace/dogu' but is '%s'", qualifiedName)
	}
	return NewQualifiedDoguName(
		DoguNamespace(splitName[0]),
		SimpleDoguName(splitName[1]),
	)
}
