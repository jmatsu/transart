package config

import "errors"

var destinationNotSupported = errors.New("destination is not supported")
var sourceNotSupported = errors.New("source is not supported")
