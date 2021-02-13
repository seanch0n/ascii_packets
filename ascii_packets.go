package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
	Read the file, count the number of lines. The number of lines - 1 is the number of connections
	in the sequence.
	We need to build a rectangle for each actor, the length of num lines - 1
	Then build arrows for the connections.
*/
func buildSequence(lines []string, leftNode string, rightNode string) string {
	corner := "*"
	c := "-"
	arrowLength := getMaxLengthMsg(lines) + 6
	var final []string
	horizontalEdgeLen := arrowLength + len(leftNode) + len(rightNode) + 6
	horizontalEdge := corner + genStringOfLen(c, horizontalEdgeLen) + corner
	final = append(final, horizontalEdge)
	for idx, line := range lines {
		idx++
		insert := false
		// | leftNode |-----SYN---->| rightNode |
		if len(lines)/idx == 2 {
			// we need to insert the word half way through the rectangle
			insert = true
		}
		arrowText, err := parseMsg(line)
		if err != nil {
			return ""
		}
		arrowDirection, err := findDirection(line, leftNode, rightNode)
		if err != nil {
			return ""
		}
		final = append(final, buildWall(leftNode, insert)+genArrow(arrowDirection, arrowLength, arrowText)+buildWall(rightNode, insert))
	}
	final = append(final, horizontalEdge)

	return strings.Join(final[:], "\n")
}

func findDirection(input string, leftNode string, rightNode string) (string, error) {
	nodeLoc := strings.IndexByte(input, '-')
	if nodeLoc == -1 {
		nodeLoc = strings.IndexByte(input, '=')
		if nodeLoc == -1 {
			return "", &errorString{"Failed to find a node in string"}
		}
		return "bi", nil
	}
	node := input[:nodeLoc]
	if node == leftNode {
		return "right", nil
	} else if node == rightNode {
		return "left", nil
	} else {
		return "", &errorString{"Node does not match left or right nodes"}
	}
}

func main() {
	filePathPtr := flag.String("path", "input.txt", "path to file with diagram data")
	node1Ptr := flag.String("node1", "client", "name of the first node")
	node2Ptr := flag.String("node2", "server", "name of the 2nd node")

	flag.Parse()
	file, err := os.Open(*filePathPtr)
	if err != nil {
		log.Fatal("Error opening file")
	}
	defer file.Close()
	// gets an array of each line of the file.
	lines, err := readFile(file)
	if err != nil {
		log.Fatal("Error reading file")
	}
	fmt.Println(buildSequence(lines, *node1Ptr, *node2Ptr))
}
