package main

import (
	"fmt"
	"strconv"
)

func main() {
	var floatNumStr string
	var _, err = fmt.Scan(&floatNumStr)
	
	if err == nil {
		var floatNum, _  = strconv.ParseFloat(floatNumStr, 8)
		fmt.Println(int(floatNum))
	} else {
		fmt.Println("Wrong input data :) ")
	}
}
