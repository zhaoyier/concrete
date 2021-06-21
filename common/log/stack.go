package log

import (
	"fmt"
	"path/filepath"
	"runtime"
)

type FieldMap map[string]interface{}

func String(depth int) string {
	pc, _, n, ok := runtime.Caller(1 + depth)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%v:%v",
		filepath.Base(runtime.FuncForPC(pc).Name()), n)
}
