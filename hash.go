package task

import (
	"fmt"

	"github.com/saturn4er/task/v3/internal/hash"
	"github.com/saturn4er/task/v3/taskfile/ast"
)

func (e *Executor) GetHash(t *ast.Task) (string, error) {
	r := t.Run
	if r == "" {
		r = e.Taskfile.Run
	}

	var h hash.HashFunc
	switch r {
	case "always":
		h = hash.Empty
	case "once":
		h = hash.Name
	case "when_changed":
		h = hash.Hash
	default:
		return "", fmt.Errorf(`task: invalid run "%s"`, r)
	}
	return h(t)
}
