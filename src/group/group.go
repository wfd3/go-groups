// Package group allows group account lookups by name or id.
package group

import (
	"strconv"
)

var implemented = true // Set false by lookup_stub.go

// Group represents a group account and includes a list of member usernames.
type Group struct {
	Name    string
	Gid     int
	Members []string
}

// UnknownGroupIdError is returned by LookupId when a group ID cannot be found.
type UnknownGroupIdError int

func (e UnknownGroupIdError) Error() string {
	return "group: unknown gid " + strconv.Itoa(int(e))
}

// UnknownGroupError is returned by Lookup when a group name cannot be found.
type UnknownGroupError string

func (e UnknownGroupError) Error() string {
	return "group: unknown group " + string(e)
}
