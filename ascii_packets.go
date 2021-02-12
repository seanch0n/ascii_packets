package ascii_packets

import (
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"unicode/utf8"
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

/*
	get the length of the longest string in an array
	return: the length of the longest string
*/
func getMaxLengthString(input []string) int {
	max := -1
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

func readFile(reader io.Reader) ([]string, error) {
	lines, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(lines), "\n"), nil
}

func buildWall(text string, insert bool) string {
	pad := genStringOfLen(" ", (len(text) + 2))
	if insert {
		pad = " " + text + " "
	}
	final := "|" + pad + "|"
	return final
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

/*
	Read the file, count the number of lines. The number of lines - 1 is the number of connections
	in the sequence.
	We need to build a rectangle for each actor, the length of num lines - 1
	Then build arrows for the connections.
*/
// func buildSequence(lines string, leftNode string, rightNode string) string {
// 	corner := "*"
// 	side := "|"
// 	c := "-"
// 	var final string
// 	for idx, line := range lines {
// 		insert := false
// 		// | leftNode |-----SYN---->| rightNode |
// 		if idx/len(lines) == 2 {
// 			// we need to insert the word half way through the rectangle
// 			insert = true
// 		}
// 		final += buildWall(leftNode, insert) // + arrow + buildWall(rightNode, insert)
// 	}

// 	return "ahh"
// }

func readDataFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// gets an array of each line of the file.
	lines, err := readFile(file)
	if err != nil {
		return nil, err
	}
	return lines, nil
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
