// Package main provides ...
package newDir

import (
	"fmt"
	"os"
)

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	fmt.Println(os.Stderr, fmt.Errorf("error msg"))
	println("-----------------------------JustDemo end>>>")
	return
}