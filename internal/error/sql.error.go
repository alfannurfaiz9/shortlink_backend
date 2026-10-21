package custom_error

import "errors"

var ErrNoRows = errors.New("no rows affected")
