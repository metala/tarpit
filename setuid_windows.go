package main

import (
	"errors"
)

func setUID(uid uint16) error {
	if uid != 0 {
		return errors.New("unable to setuid on Windows")
	}
	return nil
}

func setGID(uid uint16) error {
	if uid != 0 {
		return errors.New("unable to setgid on Windows")
	}
	return nil
}
