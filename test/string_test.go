package test

import (
	"testing"

	"github.com/sweetycode/goal"
)

func TestStringIn(t *testing.T) {
	if goal.StringIn("goal", []string{}) {
		t.Fail()
	}

	if goal.StringIn("goal", []string{"hello", "world"}) {
		t.Fail()
	}

	if !goal.StringIn("goal", []string{"goal"}) {
		t.Fail()
	}
}

func TestSplit2(t *testing.T) {
	if s1, s2 := goal.Split2("www.pystarter.com", "."); s1 != "www" || s2 != "pystarter.com" {
		t.Fail()
	}
	if s1, s2 := goal.Split2("www.pystarter.com", "/"); s1 != "www.pystarter.com" || s2 != "" {
		t.Fail()
	}
	if s1, s2 := goal.Split2("www.pystarter.com", "pystarter"); s1 != "www." || s2 != ".com" {
		t.Fail()
	}
}

func TestRSplit2(t *testing.T) {
	if s1, s2 := goal.RSplit2("www.pystarter.com", "."); s1 != "www.pystarter" || s2 != "com" {
		t.Fail()
	}

	if s1, s2 := goal.RSplit2("www.pystarter.com", "/"); s1 != "www.pystarter.com" || s2 != "" {
		t.Fail()
	}
	if s1, s2 := goal.RSplit2("www.pystarter.com", "pystarter"); s1 != "www." || s2 != ".com" {
		t.Fail()
	}
}
