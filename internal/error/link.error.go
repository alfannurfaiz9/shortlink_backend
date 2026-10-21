package custom_error

import "errors"

var EmptyField = errors.New("all field must be filled")
var MinimumSlug = errors.New("minimum slug is 3 character")
var MaximumSlug = errors.New("maximum slug is 50 character")
var ReservedKeyword = errors.New("cannot use api, login, register, dashboard as a slug")
