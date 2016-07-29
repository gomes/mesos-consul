package mesos

import (
	"sort"
	"strings"
	"testing"
)

func TestLeaderIP(t *testing.T) {
	l := "master@124.123.123.121:5050"

	ip := leaderIP(l)

	t.Log("ip: ", ip)
}

func TestSliceEq(t *testing.T) {
	for _, tt := range []struct {
		a []string
		b []string
		r bool
	}{
		{[]string{}, []string{}, true},
		{[]string{"one"}, []string{}, false},
		{[]string{"one"}, []string{"one"}, true},
		{[]string{"one"}, []string{"two"}, false},
		{[]string{"one"}, []string{"one", "two"}, false},
	} {
		r := sliceEq(tt.a, tt.b)
		if r != tt.r {
			t.Errorf("sliceEq(%v, %v) => %t, want %t", tt.a, tt.b, r, tt.r)
		}
	}
}

func TestSliceContainsString(t *testing.T) {
	for _, tt := range []struct {
		s []string
		b string
		r bool
	}{
		{[]string{}, "one", false},
		{[]string{}, "", false},
		{[]string{"one"}, "one", true},
		{[]string{"one"}, "", false},
		{[]string{"one"}, "two", false},
		{[]string{"one", "two"}, "one", true},
		{[]string{"one", "two"}, "three", false},
	} {
		r := sliceContainsString(tt.s, tt.b)
		if r != tt.r {
			t.Errorf("sliceContainsString(%v, %s) => %t, want %t", tt.s, tt.b, r, tt.r)
		}
	}
}

func TestEscapedComma(t *testing.T) {
	cases := []struct {
		Tag      string
		Expected []string
	}{
		{
			Tag:      "",
			Expected: []string{},
		},
		{
			Tag:      "foobar",
			Expected: []string{"foobar"},
		},
		{
			Tag:      "foo,bar",
			Expected: []string{"foo", "bar"},
		},
		{
			Tag:      "foo\\,bar",
			Expected: []string{"foo,bar"},
		},
		{
			Tag:      "foo,bar\\,baz",
			Expected: []string{"foo", "bar,baz"},
		},
		{
			Tag:      "\\,foobar\\,",
			Expected: []string{",foobar,"},
		},
		{
			Tag:      ",,,,foo,,,bar,,,",
			Expected: []string{"foo", "bar"},
		},
		{
			Tag:      ",,,,",
			Expected: []string{},
		},
		{
			Tag:      ",,\\,,",
			Expected: []string{","},
		},
	}

	for _, c := range cases {
		results := recParseEscapedComma(c.Tag)
		sort.Strings(c.Expected)
		sort.Strings(results)
		if !sliceEq(results, c.Expected) {
			t.Error("Expected [" + strings.Join(c.Expected, ",") + "] but result was [" + strings.Join(results, ",") + "]")
		}
	}
}
