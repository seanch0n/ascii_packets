package ascii_packets

import (
	"io"
	"io/ioutil"
	"os"
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
// 		// | leftNode |-----SYN---->| rightNode |
// 		final += buildWall(leftNode)
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

func genArrow(arrowType string, len int) string {
	var base string
	switch arrowType {
	case "right":
		base = genStringOfLen("-", (len - 1))
		base += ">"
	case "left":
		base = genStringOfLen("-", (len - 1))
		base = "<" + base
	case "bi":
		base = genStringOfLen("-", (len - 2))
		base = "<" + base + ">"
	}
	return base
}
