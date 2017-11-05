package main

import (
    "regexp"
)

// RegexSubMatchMap returns a slice of the matches, if any, of subexpressions.
func RegexSubMatchMap(r *regexp.Regexp, s string) (matchMap map[string]string) {
    match := r.FindStringSubmatch(s)
    matchMap = make(map[string]string)
    for i, name := range r.SubexpNames() {
        if i > 0 && len(match) >= i && name != "" {
            matchMap[name] = match[i]
        }
    }
    return matchMap
}
