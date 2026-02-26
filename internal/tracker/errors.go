package tracker

import "errors"

var ErrNotFound = errors.New("not found")
var ErrItemExist = errors.New("Item with such Id already exists ")
