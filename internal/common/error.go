package common

import (
	"fmt"
	"runtime"
)

func WrapError(err error) error {

	if err == nil {
		return nil
	}

	_, filename, line, _ := runtime.Caller(1)
	return fmt.Errorf("[error] %s %d: %w", filename, line, err)
}
