package echovalidate

import (
	"encoding/hex"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

func Required(key string, value interface{}) error {
	hasFailed := false
	switch value.(type) {
	case string:
		if value == "" {
			hasFailed = true
		}
	}

	if hasFailed {
		return errors.New(fmt.Sprintf("%v is required", strings.ReplaceAll(key, "_", " ")))
	}
	return nil
}

func ValidEmail(key, value string) error {
	emailRegexp := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)+$")
	if !emailRegexp.MatchString(value) {
		return errors.New(fmt.Sprintf("%v is not a valid email", value))
	}
	return nil
}

func In(key, needle string, haystack []string) error {
	for _, v := range haystack {
		if needle == v {
			return nil
		}
	}
	return errors.New(fmt.Sprintf(
		"Supplied Value for %v is not supported - Allowed Values: %v",
		key,
		strings.Join(haystack, " | "),
	))
}

func MinLen(key string, valueLen, minLen int) error {
	if valueLen < minLen {
		return errors.New(fmt.Sprintf("%v must have at least %v items | characters | keys ", key, minLen))
	}
	return nil
}

func MaxLen(key string, valueLen, maxLen int) error {
	if valueLen > maxLen {
		return errors.New(fmt.Sprintf("%v must have less than %v items | characters | keys ", key, maxLen))
	}
	return nil
}

func ValidMongoObjectID(key, value string) error {
	//Slightly modified code from the go-mongo driver
	b, err := hex.DecodeString(value)
	if err != nil {
		return errors.New(fmt.Sprintf("invalid ObjectID provided for %v", key))
	}

	if len(b) != 12 {
		return errors.New(fmt.Sprintf("invalid ObjectID provided for %v", key))
	}

	return nil
}

func ValidDateTime(key, value, layout string) error {
	if _, err := time.Parse(layout, value); err != nil {
		return errors.New(fmt.Sprintf("invalid date/time format provided for %v - Sample: %v", key, value))
	}
	return nil
}
