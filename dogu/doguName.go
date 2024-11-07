package dogu

import (
	"errors"
	"fmt"
	"github.com/cloudogu/cesapp-lib/core"
	"strings"
)

type SimpleName string
type Namespace string

// String returns the string representation of the SimpleName.
func (s SimpleName) String() string {
	return string(s)
}

type QualifiedVersion struct {
	Name    QualifiedName
	Version core.Version
}
type QualifiedName struct {
	SimpleName SimpleName
	Namespace  Namespace
}

func NewQualifiedName(namespace Namespace, simpleName SimpleName) (QualifiedName, error) {
	doguName := QualifiedName{Namespace: namespace, SimpleName: simpleName}
	err := doguName.Validate()
	if err != nil {
		return QualifiedName{}, err
	}
	return doguName, nil
}

func NewQualifiedVersion(name QualifiedName, version core.Version) (QualifiedVersion, error) {
	err := name.Validate()
	if err != nil {
		return QualifiedVersion{}, err
	}
	return QualifiedVersion{
		Name:    name,
		Version: version,
	}, nil
}

func (name QualifiedName) Validate() error {
	var errorList []error
	if name.Namespace == "" {
		errorList = append(errorList, fmt.Errorf("namespace of dogu %q must not be empty", name.SimpleName))
	}
	if name.SimpleName == "" {
		errorList = append(errorList, fmt.Errorf("dogu name must not be empty"))
	}
	splitName := strings.Split(string(name.Namespace), "/")
	if len(splitName) > 1 {
		errorList = append(errorList, fmt.Errorf("dogu name needs to be in the form 'namespace/dogu' but is '%s'", name))
	}

	return errors.Join(errorList...)
}

// String returns the dogu name with namespace, e.g. official/postgresql
func (name QualifiedName) String() string {
	return fmt.Sprintf("%s/%s", name.Namespace, name.SimpleName)
}

// QualifiedNameFromString converts a qualified dogu as a string, e.g. "official/nginx", to a dedicated QualifiedName or raises an error if this is not possible.
func QualifiedNameFromString(qualifiedName string) (QualifiedName, error) {
	splitName := strings.Split(qualifiedName, "/")
	if len(splitName) != 2 {
		return QualifiedName{}, fmt.Errorf("dogu name needs to be in the form 'namespace/dogu' but is '%s'", qualifiedName)
	}
	return NewQualifiedName(
		Namespace(splitName[0]),
		SimpleName(splitName[1]),
	)
}
