package output

import (
	"fmt"
	"io"

	"github.com/saturn4er/task/v3/taskfile/ast"
)

// Templater executes a template engine.
// It is provided by the templater.Templater package.
type Templater interface {
	// Replace replaces the provided template string with a rendered string.
	Replace(tmpl string) string
}

type Output interface {
	WrapWriter(stdOut, stdErr io.Writer, prefix string, tmpl Templater) (io.Writer, io.Writer, CloseFunc)
}

type CloseFunc func(err error) error

// Build the Output for the requested ast.Output.
func BuildFor(o *ast.Output) (Output, error) {
	switch o.Name {
	case "interleaved", "":
		if err := checkOutputGroupUnset(o); err != nil {
			return nil, err
		}
		return Interleaved{}, nil
	case "group":
		return Group{
			Begin:     o.Group.Begin,
			End:       o.Group.End,
			ErrorOnly: o.Group.ErrorOnly,
		}, nil
	case "prefixed":
		if err := checkOutputGroupUnset(o); err != nil {
			return nil, err
		}
		return Prefixed{}, nil
	default:
		return nil, fmt.Errorf(`task: output style %q not recognized`, o.Name)
	}
}

func checkOutputGroupUnset(o *ast.Output) error {
	if o.Group.IsSet() {
		return fmt.Errorf("task: output style %q does not support the group begin/end parameter", o.Name)
	}
	return nil
}
