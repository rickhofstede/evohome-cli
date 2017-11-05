package main

import (
    "regexp"
    "testing"
)

func TestReSubMatchMap(t *testing.T) {  
    re := regexp.MustCompile(`(?P<group1>[a-z]+) (?P<group2>[0-9]+)`)
    matches := RegexSubMatchMap(re, "abc 123")

    if len(matches) != 2 {
        t.Errorf("Number of match groups incorrect; got: %d, expected: %d", len(matches), 2)
    }
    if matches["group1"] != "abc" {
        t.Errorf("Matched value incorrect; got: %s, expected: %s", matches["group1"], "abc")   
    }
    if matches["group2"] != "123" {
        t.Errorf("Matched value incorrect; got: %s, expected: %s", matches["group2"], "123")   
    }
}

func TestReSubMatchMapEmpty(t *testing.T) {  
    re := regexp.MustCompile(`([a-z]+)`)
    matches := RegexSubMatchMap(re, "abc 123")

    if len(matches) != 0 {
        t.Errorf("Number of match groups incorrect; got: %d, expected: %d", len(matches), 0)
    }
}

func TestReSubMatchMapNonMatchingRegex(t *testing.T) {  
    re := regexp.MustCompile(`(?P<group1>[a-z]+) (?P<group2>[0-9]+)`)
    matches := RegexSubMatchMap(re, "abc")

    if len(matches) != 0 {
        t.Errorf("Number of match groups incorrect; got: %d, expected: %d", len(matches), 0)
    }
}
