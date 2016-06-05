// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"os"
)

func main() {
	conf, _ := parseArgs()
	fmt.Printf("Hello, %s\n", conf.Arg)
}

type Config struct {
	Arg string
}

func parseArgs() (Config, error) {
	arg1 := os.Args[1]

	return Config{
		Arg: arg1,
	}, nil
}
