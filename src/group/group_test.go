package group

import (
	"reflect"
	"testing"
)

func skip(t *testing.T) bool {
	if !implemented {
		t.Logf("group: no implemented, skipping tests")
		return true
	}
	return false
}

func compare(t *testing.T, want, got *Group) {
	if want.Gid != got.Gid {
		t.Errorf("got Gid=%q; want=%q", got.Gid, want.Gid)
	}
	if want.Name != got.Name {
		t.Errorf("got Group=%q; want=%q", got.Name, want.Name)
	}
	if !reflect.DeepEqual(want.Members, got.Members) {
		t.Errorf("got Members=%q; want=%q", got.Members, want.Members)
	}
}

func TestCurrent(t *testing.T) {
	if skip(t) {
		return
	}
	g, err := Current()
	if err != nil {
		t.Fatalf("Current: %v", err)
	}
	if g.Name == "" {
		t.Fatal("didn't get a group name")
	}
}

func TestLookupId(t *testing.T) {
	if skip(t) {
		return
	}

	want, err := Current()
	if err != nil {
		t.Fatalf("Current: %v", err)
	}
	got, err := LookupId(want.Gid)
	if err != nil {
		t.Fatalf("LookupId: %v", err)
	}
	compare(t, want, got)
}

func TestLookup(t *testing.T) {
	if skip(t) {
		return
	}
	want, err := Current()
	if err != nil {
		t.Fatalf("Current: %v", err)
	}
	got, err := Lookup(want.Name)
	if err != nil {
		t.Fatalf("Lookup: %v", err)
	}
	compare(t, want, got)
}
