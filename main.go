//main.go
package main

import (
	"os"
	"qanly_chain/cli"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
}
