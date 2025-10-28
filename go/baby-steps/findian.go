package main

import (
	"bufio"
	"fmt"
	"os"
	s "strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var input, _ = in.ReadString('\n')

	var trimmed = s.TrimSpace(s.ReplaceAll(input, " ", ""))
	var str = s.ToLower(trimmed)

	switch {
	case s.HasPrefix(str, "i") && s.HasSuffix(str, "n") && s.Contains(str, "a"):
		fmt.Println("Found!")
	default:
		fmt.Println("Not found!")
	}
}
