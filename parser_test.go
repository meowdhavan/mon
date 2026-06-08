package moon

import (
	"slices"
	"testing"
)

func TestLongStringFlagParse(t *testing.T) {
	var targetA string
	var targetB string

	c := Command{}
	c.AddStringFlag(&targetA, []string{"test-flag-a"}, "", "", false)
	c.AddStringFlag(&targetB, []string{"test-flag-b"}, "", "", false)

	p := newParser(&c, []string{"app", "--test-flag-a", "target_value_1", "--test-flag-b", "target_value_2"})
	p.parseFlags()

	var gotString string

	gotString = "target_value_1"
	if targetA != gotString {
		t.Errorf("targetA mismatch; got=%s, want %s", targetA, gotString)
	}

	gotString = "target_value_2"
	if targetB != gotString {
		t.Errorf("targetB mismatch; got=%s, want %s", targetB, gotString)
	}
}

func TestShortStringFlagParse(t *testing.T) {
	var targetA string
	var targetB string

	c := Command{}
	c.AddStringFlag(&targetA, []string{}, "a", "", false)
	c.AddStringFlag(&targetB, []string{}, "b", "", false)

	p := newParser(&c, []string{"app", "-a", "target_value_1", "-btarget_value_2"})
	p.parseFlags()

	var gotString string

	gotString = "target_value_1"
	if targetA != gotString {
		t.Errorf("targetA mismatch; got=%s, want %s", targetA, gotString)
	}

	gotString = "target_value_2"
	if targetB != gotString {
		t.Errorf("targetB mismatch; got=%s, want %s", targetB, gotString)
	}
}

func TestStringPosArgParse(t *testing.T) {
	var targetA string
	var targetB string

	c := Command{}
	c.AddStringPosArg(&targetA, "a", "", false)
	c.AddStringPosArg(&targetB, "b", "", false)

	p := newParser(&c, []string{"app", "target_value_1", "target_value_2"})
	p.parseFlags()

	var gotString string

	gotString = "target_value_1"
	if targetA != gotString {
		t.Errorf("targetA mismatch; got=%s, want %s", targetA, gotString)
	}

	gotString = "target_value_2"
	if targetB != gotString {
		t.Errorf("targetB mismatch; got=%s, want %s", targetB, gotString)
	}
}

func TestMultitypePosArgParse(t *testing.T) {
	var targetA string
	var targetB int
	var targetSlice []int

	c := Command{}
	c.AddStringPosArg(&targetA, "a", "", true)
	c.AddIntPosArg(&targetB, "b", "", false)
	c.AddIntVarLenArg(&targetSlice, "vla", "")

	p := newParser(&c, []string{"app", "target_value_1", "123", "10", "20", "30"})
	p.parseFlags()

	gotString := "target_value_1"
	if targetA != gotString {
		t.Errorf("targetA mismatch; got=%s; want %s", targetA, gotString)
	}

	gotInt := 123
	if targetB != gotInt {
		t.Errorf("targetB mismatch; got=%d; want %d", targetB, gotInt)
	}

	gotIntSlice := []int{10, 20, 30}
	if !slices.Equal(targetSlice, gotIntSlice) {
		t.Errorf("targetSlice mismatch; got=%v; want %v", targetSlice, gotIntSlice)
	}
}
