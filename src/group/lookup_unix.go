package group

// +build darwin freebsd linux
// +build cgo

import (
	"fmt"
	"syscall"
	"unsafe"
)

/*
#include <unistd.h>
#include <sys/types.h>
#include <grp.h>
#include <stdlib.h>
*/
import "C"

// Convert (**char)clist to []string
func convert(clist **C.char) []string {
	var members []string

	p := (*[1 << 30]*C.char)(unsafe.Pointer(clist))
	for i := 0; p[i] != nil; i++ {
		members = append(members, C.GoString(p[i]))
	}

	return members
}

// Current returns the group information
func Current() (*Group, error) {
	return lookup(syscall.Getgid(), "", false)
}

// LookupId returns the group information by GID.  If the group cannot be found,
// the error UnknownGroupIdError is returned.
func LookupId(gid int) (*Group, error) {
	return lookup(gid, "", false)
}

// Lookup returns the group information by group name.  If the group cannot be
// found, the error UnknownGroupError is returned.
func Lookup(groupname string) (*Group, error) {
	return lookup(-1, groupname, true)
}

func lookup(gid int, groupname string, lookupByName bool) (*Group, error) {
	var grp C.struct_group
	var result *C.struct_group
	var bufsize C.long

	bufsize = C.sysconf(C._SC_GETGR_R_SIZE_MAX)
	if bufsize == -1 {
		bufsize = 1024
	}
	buf := C.malloc(C.size_t(bufsize))
	defer C.free(buf)

	var rv C.int
	if lookupByName {
		CGroup := C.CString(groupname)
		defer C.free(unsafe.Pointer(CGroup))
		rv = C.getgrnam_r(CGroup, &grp, (*C.char)(buf),
			C.size_t(bufsize), &result)
		if rv != 0 {
			return nil,
				fmt.Errorf("group: lookup group name %s: %s",
					groupname, syscall.Errno(rv))
		}
		if result == nil {
			return nil, UnknownGroupError(groupname)
		}
	} else {
		rv = C.getgrgid_r(C.gid_t(gid), &grp, (*C.char)(buf),
			C.size_t(bufsize), &result)
		if rv != 0 {
			return nil, fmt.Errorf("group: lookup gid %d: %s",
				gid, syscall.Errno(rv))
		}
		if result == nil {
			return nil, UnknownGroupIdError(gid)
		}
	}

	g := &Group{
		Gid:     int(grp.gr_gid),
		Name:    C.GoString(grp.gr_name),
		Members: convert(grp.gr_mem),
	}

	return g, nil
}
