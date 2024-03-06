package env

import (
	"fmt"
	"os"

	"github.com/saturn4er/task/v3/taskfile/ast"
)

func Get(t *ast.Task) []string {
	if t.Env == nil {
		return nil
	}

	environ := os.Environ()

	for k, v := range t.Env.ToCacheMap() {
		str, isString := v.(string)
		if !isString {
			continue
		}

		if _, alreadySet := os.LookupEnv(k); alreadySet && !t.Env.Get(k).Overwrite {
			continue
		}

		environ = append(environ, fmt.Sprintf("%s=%s", k, str))
	}

	return environ
}
