package main

import (
    "strings"
    "testing"
    "github.com/c-bata/go-prompt"
)

type CompleterTest struct {
    command string
    expected []string
}

// Verify whether two lists are equal.
func (t *CompleterTest) equal(actual []string) (bool) {
    if len(t.expected) != len(actual) {
        return false
    }
    for i, elem := range actual {
        if elem != t.expected[i] {
            return false
        }
    }
    return true
}

// Get all command texts from list of commands.
func commandTexts(commands []prompt.Suggest) (texts []string) {
    for _, cmd := range commands {
        texts = append(texts, cmd.Text)
    }
    return texts
}

func TestCompleterShowCmd(t *testing.T) {
    mock := MockSystem {
        Id: "12345",
        Type: "MockSystem",
        Zones: nil,
    }

    tests := []CompleterTest{
        // General
        CompleterTest{command: "", expected: []string {"cancel", "set", "show", "exit"}},
        CompleterTest{command: " ", expected: []string {}},
        CompleterTest{command: "\t", expected: []string {}},

        // show
        CompleterTest{command: "show ", expected: []string {"zone", "zones"}},
        CompleterTest{command: "show zone", expected: []string {"zone", "zones"}},
        CompleterTest{command: "show zone ", expected: []string {"TestZone"}},
        CompleterTest{command: "show zone TestZone", expected: []string {"TestZone"}},
        CompleterTest{command: "show zone TestZone ", expected: []string {"schedule"}},
        CompleterTest{command: "show zone TestZone schedule ", expected: []string {}},
        CompleterTest{command: "show zones", expected: []string {"zones"}},
        CompleterTest{command: "show zones ", expected: []string {}},

        // set
        CompleterTest{command: "set ", expected: []string {"zone"}},
        CompleterTest{command: "set zone", expected: []string {"zone"}},
        CompleterTest{command: "set zone ", expected: []string {"TestZone"}},
        CompleterTest{command: "set zone TestZone ", expected: []string {"temperature"}},
        CompleterTest{command: "set zone TestZone temperature ", expected: []string {}},
        CompleterTest{command: "set zone TestZone temperature 15.0", expected: []string {}},
        CompleterTest{command: "set zone TestZone temperature -15.0", expected: []string {}},
        CompleterTest{command: "set zone TestZone temperature -1a.0", expected: []string {}},
        CompleterTest{command: "set zone TestZone temperature 15.0 ", expected: []string {"until"}},
        CompleterTest{command: "set zone TestZone temperature 15.0 until ", expected: []string {"yyyy/mm/dd hh:mm"}},
        CompleterTest{command: "set zone TestZone temperature 15.0 until 2017/11/24 15:00", expected: []string {}},
        CompleterTest{command: "set zone TestZone temperature 15.0 until 2017/11/24 15:00 ", expected: []string {}},

        // cancel
        CompleterTest{command: "cancel ", expected: []string {"zone"}},
        CompleterTest{command: "cancel zone", expected: []string {"zone"}},
        CompleterTest{command: "cancel zone ", expected: []string {"TestZone"}},
        CompleterTest{command: "cancel zone TestZone ", expected: []string {"temperature"}},
        CompleterTest{command: "cancel zone TestZone temperature ", expected: []string {"override"}},
        CompleterTest{command: "cancel zone TestZone temperature override ", expected: []string {}},
    }

    for _, test := range tests {
        actual := commandTexts(argumentsCompleter(strings.Split(test.command, " "), &mock))
        if !test.equal(actual) {
            t.Errorf("Completion list mismatch; expected: %+v, got: %+v", test.expected, actual)
        }
    }
}
