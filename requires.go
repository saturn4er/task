package task

import (
	"context"

	"github.com/saturn4er/task/v3/errors"
	"github.com/saturn4er/task/v3/taskfile/ast"
)

func (e *Executor) areTaskRequiredVarsSet(ctx context.Context, t *ast.Task, call *ast.Call) error {
	if t.Requires == nil || len(t.Requires.Vars) == 0 {
		return nil
	}

	vars, err := e.Compiler.GetVariables(t, call)
	if err != nil {
		return err
	}

	var missingVars []string
	for _, requiredVar := range t.Requires.Vars {
		if !vars.Exists(requiredVar) {
			missingVars = append(missingVars, requiredVar)
		}
	}

	if len(missingVars) > 0 {
		return &errors.TaskMissingRequiredVars{
			TaskName:    t.Name(),
			MissingVars: missingVars,
		}
	}

	return nil
}
