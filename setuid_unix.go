// +build linux darwin

package main

import (
	"errors"
	"syscall"
)

func setUidGid(syscallID uint, uidgid uint16) error {
	if uidgid == 0 {
		return nil
	}
	_, _, errno := syscall.Syscall(uintptr(syscallID), uintptr(uidgid), 0, 0)
	if errno != 0 {
		return errors.New(errno.Error())
	}
	return nil
}

func setUID(uid uint16) error {
	return setUidGid(syscall.SYS_SETUID, uid)
}

func setGID(gid uint16) error {
	return setUidGid(syscall.SYS_SETGID, gid)
}
