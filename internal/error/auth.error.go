package custom_error

import "errors"

var UserAlreadyExist = errors.New("user already exist")
var EmptyLoginField = errors.New("all field must be filled")
var InvalidEmailOrPassword = errors.New("invalid email or password")
var TokenExpired = errors.New("please login first")
