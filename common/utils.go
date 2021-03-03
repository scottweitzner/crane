package common

import (
	"errors"

	"github.com/fatih/color"
)

// RedWrapError is a utility function to wrap an error in the color red
func RedWrapError(err error) error {
	if err == nil {
		return err
	}
	return errors.New(color.RedString(err.Error()))
}
