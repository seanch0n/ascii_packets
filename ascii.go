package main

import (
	"strings"
	"unicode/utf8"
)

/*
	this function generates a border around a string. The string can be multiline
	or single line, with varying line lengths. See unit test for examples.
	return is the fully formatted string
*/
func generateBorder(text string) string {
	lines := strings.Split(text, "\n")
	width := getMaxLengthString(lines)
	corner := "*"
	side := "|"
	c := "-"
	// we do (width + 2) to account for single space on either side of the line
	top := corner + genStringOfLen(c, (width+2)) + corner
	var bordered []string
	bordered = append(bordered, top)
	for _, line := range lines {
		pad := genStringOfLen(" ", width-len(line))
		bordered = append(bordered, (side + " " + line + " " + pad + side))
	}
	bordered = append(bordered, top)

	return strings.Join(bordered[:], "\n")
}

func buildWall(text string, insert bool) string {
	pad := genStringOfLen(" ", (len(text) + 2))
	if insert {
		pad = " " + text + " "
	}
	final := "|" + pad + "|"
	return final
}

// refactor this to take the string for the middle of the arrow:
// ----SYN---->
// Note that it will truncate for odds.
func genArrow(arrowType string, len int, text string) string {
	var base string
	genLength := len - utf8.RuneCountInString(text)
	// When we have an odd number length, we don't want to truncate, but we also don't want to over add
	// by adding the same amount to either side. We need to balance the offset. To do this, we put the extra
	// character at the end of the arrow
	endLength := (genLength - 1) / 2
	if (genLength-1)%2 != 0 {
		endLength = ((genLength - 1) / 2) + 1
	}
	switch arrowType {
	case "right":
		base = genStringOfLen("-", ((genLength-1)/2)) + text + genStringOfLen("-", endLength)
		base += ">"
	case "left":
		base = genStringOfLen("-", ((genLength-1)/2)) + text + genStringOfLen("-", endLength)
		base = "<" + base
	case "bi":
		base = genStringOfLen("-", ((genLength-2)/2)) + text + genStringOfLen("-", endLength)
		base = "<" + base + ">"
	}
	return base
}
