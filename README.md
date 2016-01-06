# go-groups
Get group information by name or ID in Golang

PACKAGE DOCUMENTATION

package group
    import "group"

Package group allows group account lookups by name or id.

TYPES

    type Group struct {
        Name    string
        Gid     int
        Members []string
    }
Group represents a group account and includes a list of member usernames.

    func Current() (*Group, error)
Current returns the curreng group information (from getgid())

    func Lookup(groupname string) (*Group, error)
Lookup returns the group information for a specific group name. If the group cannot be found, the error UnknownGroupError is returned.

    func LookupId(gid int) (*Group, error)
LookupId returns the group information for a specific GID. If the group cannot be found, the error UnknownGroupIdError is returned.

    type UnknownGroupError string
UnknownGroupError is returned by Lookup when a group name cannot be found.

    func (e UnknownGroupError) Error() string

    type UnknownGroupIdError int
UnknownGroupIdError is returned by LookupId when a group ID cannot be found.

    func (e UnknownGroupIdError) Error() string
