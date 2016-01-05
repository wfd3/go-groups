package main

import (
	"group"
	"fmt"
	"flag"
)


func main() {

	var gid int
	var gname string
	var G *group.Group
	var e error

	flag.IntVar(&gid, "g", -1, "GID")
	flag.StringVar(&gname, "n", "", "Group name")
	flag.Parse()

	if gid != -1 {
		G, e = group.LookupId(gid)
	} 
	if gname != "" {
		G, e = group.Lookup(gname)
	}
	if G == nil { 
		fmt.Printf("Error: %v\n", e)
	} else {
		fmt.Printf("Group '%s' id %d\n", G.Name, G.Gid)
		fmt.Printf("Members %v\n", G.Members)
	}
}