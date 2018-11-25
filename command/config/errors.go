package config

import (
	"errors"
	"fmt"
	"github.com/jmatsu/transart/config"
)

var cannotOverwriteWithoutForce = fmt.Errorf("a config file already exists. cannot overwrite without --%s option", forceOptionKey)
var locationTypeIsRequired = fmt.Errorf("either of --%s or --%s is required", sourceOptionKey, destinationOptionKey)

func sourceNotSupported(t config.LocationType) error {
	return fmt.Errorf("%s as source service is not supported", t)
}

func destinationNotSupported(t config.LocationType) error {
	return fmt.Errorf("%s as destination service is not supported", t)
}

var invalidVersion = errors.New("version must be greater than 0")
