package main

import (
	"fmt"
	"strings"
)

func RunTextJustification() {
	result := fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16)
	fmt.Println(strings.Join(result, "\n"))
	fmt.Println()
	result = fullJustify([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16)
	// fmt.Println(strings.Join(result, "\n"))
	// fmt.Println()
	// result = fullJustify([]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20)
	// fmt.Println(strings.Join(result, "\n"))
}

func fullJustify(words []string, maxWidth int) []string {
	doc := []string{}

	lineBuffer := []string{}
	currentWidth := 0
	for _, w := range words {
		if (currentWidth + len(w)) > maxWidth {
			doc = append(doc, makeLine(lineBuffer, currentWidth-1, maxWidth))
			currentWidth = 0
			lineBuffer = []string{}
		}
		lineBuffer = append(lineBuffer, w)
		currentWidth += len(w) + (len(lineBuffer) - 1)
	}

	return doc
}

func makeLine(lineBuffer []string, currentWidth, maxWidth int) string {
	needsFilling := maxWidth - currentWidth
	fmt.Printf("needs filling: %d; current width: %d\n", needsFilling, currentWidth)
	oddSpaces := 0
	if (len(lineBuffer) - 1) == 0 {
		oddSpaces = needsFilling
	} else {
		oddSpaces = needsFilling % (len(lineBuffer) - 1)
	}
	fmt.Println("odd spaces", oddSpaces)
	// create a filler.
	fillerLength := 0
	if (len(lineBuffer) - 1) == 0 {
		fillerLength = 0
	} else {
		fillerLength = (needsFilling - oddSpaces) / (len(lineBuffer) - 1)
	}
	fmt.Println("filler length", fillerLength)
	filler := ""
	for fillerLength > 0 {
		filler += " "
		fillerLength--
	}
	// use the filler to fill gaps.
	justifiedLine := []string{}
	for i, w := range lineBuffer {
		if i < len(lineBuffer)-1 {
			justifiedLine = append(justifiedLine, (w + " "))
		} else {
			justifiedLine = append(justifiedLine, w)
		}
	}
	line := strings.Join(justifiedLine, filler)
	fmt.Printf("%d | %s\n\n", len(line), line)
	return line
}
