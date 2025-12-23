package Task_9

import "errors"

var ErrNotFound = errors.New("not found")

func existence(id int) error {
	if id <= 0 {
		return ErrNotFound
	}
	return nil
}
