package torgo

import (
	"regexp"
	"strings"
)

func IsValidHiddenServiceId(service string) bool {
	service = strings.TrimSuffix(service, ".onion")
	if match, _ := regexp.Match("^[a-z2-7]{16}$", []byte(service)); match {
		return true
	} else if match, _ := regexp.Match("^[a-z2-7]{56}$", []byte(service)); match {
		return true
	}
	return false
}
