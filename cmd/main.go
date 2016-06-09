package cmd

import (
	"flag"
	"fmt"
)

func run() {
	file := flag.String("f", "", "YAML file path")
	flag.Parse()

	fmt.Println(file)
}

func main() {
	fmt.Println("yippee")
}
