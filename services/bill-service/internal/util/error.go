package util

import "fmt"

func WrapError(err error, message string) error {
	return fmt.Errorf("%s: %w", message, err)
}

func UnwrapError(err error) error {
	return fmt.Errorf("%w", err)
}
func IsError(err, target error) bool {
	return err == target
}

func CreateError(message string) error {
	return fmt.Errorf("%s", message)
}
