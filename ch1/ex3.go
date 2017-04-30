package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("print by strings.Join")
	printByJoin(os.Args[:])
	fmt.Println("print by fmt.Println")
	printByPrintln(os.Args[:])
}

func printByPrintln(args []string) {
	fmt.Println(args[:])
}

func printByJoin(args []string) {
	fmt.Println(strings.Join(args[:], " "))
}
