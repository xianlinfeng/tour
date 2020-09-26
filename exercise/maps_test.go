package exercise

import (
	"fmt"
	"strings"
	"testing"
)

func WordCount(s string) map[string]int {
	results := strings.Fields(s)

	fmt.Println(results)
	m := make(map[string]int)
	for _, v := range results {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	return m
}

func TestMap(t *testing.T) {
	fmt.Println(WordCount("Shall I compare thee to a summer’s day?  So long lives this, and this gives life to thee.     So long as men can breathe, or eyes can see, When in eternal lines to Time thou grow'st. Nor shall death brag thou wand'rest in his shade, Nor lose possession of that fair thou ow’st,"))
}
