package main

import (
	"io"
	"io/ioutil"
	"regexp"
	"strings"
)

/*
	generate a string of character @c of a given length @len
	i.e.:
		c = '-', len = 5
		-----
*/
func genStringOfLen(c string, len int) string {
	if len < 1 {
		return ""
	}
	ret := ""
	for i := 0; i < len; i++ {
		ret += c
	}
	return ret
}

/*
	get the length of the longest string in an array
	return: the length of the longest string
*/
func getMaxLengthString(input []string) int {
	max := -2
	for _, s := range input {
		if len(s) < max {
			// string is shorter, move on
			continue
		}
		if len(s) > max {
			// string is longer, time to dethrone the king
			max = len(s)
		}
	}
	return max
}

func getMaxLengthMsg(input []string) int {
	max := -1
	for _, s := range input {
		msg, err := parseMsg(s)
		if err != nil {
			return -1
		}
		if len(msg) < max {
			continue
		}
		if len(msg) > max {
			max = len(msg)
		}
	}
	return max
}

func readFile(reader io.Reader) ([]string, error) {
	lines, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(lines), "\n"), nil
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
func parseMsg(text string) (string, error) {
	// make sure we don't have too many colons. We only want 1
	r := regexp.MustCompile(":")
	count := len(r.FindAllStringIndex(text, -1))
	if count > 2 {
		return "", &errorString{"Too many values matched"}
	}

	match := strings.Split(text, ": ")
	if len(match) == 2 {
		return match[len(match)-1], nil
	}
	return "", &errorString{"Did not find a match"}
}
