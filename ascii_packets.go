package ascii_packets

import (
	"io"
	"io/ioutil"
	"log"
	"os"
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
		log.Fatal(err)
		return nil, err
	}
	return strings.Split(string(lines), "\n"), nil
}

func readDataFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	lines, err := readFile(file)
	if err != nil {
		return nil, err
	}
	return lines, nil
}
