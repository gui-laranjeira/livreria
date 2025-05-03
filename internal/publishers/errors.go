package publishers

import "errors"

var ErrPublisherNameRequired = errors.New("publisher name required")
var ErrPublisherNotFound = errors.New("publisher not found")
