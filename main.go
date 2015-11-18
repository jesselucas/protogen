package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jesselucas/executil"
)

// Semantic Version
const VERSION = "0.0.2"

const space = "  " // Used for pretty printing
const br = "\n\n"

func main() {
	// Check for our command line configuration flags
	var (
		versionUsage = "Prints current version" + " (v. " + VERSION + ")"
		versionPtr   = flag.Bool("version", false, versionUsage)
	)

	// Set up short hand flags
	flag.BoolVar(versionPtr, "v", false, versionUsage+" (shorthand)")
	flag.Parse()

	if *versionPtr {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	// usage description
	usage := "protogen /path/to/myService.proto"
	options := "--version (-v)"
	optionsDesc := versionUsage

	// Create help message with usage messaging
	helpMessage := fmt.Sprintf(bold("USAGE:")+"\n%s%v", space, usage)
	// Break between messages
	helpMessage += br
	// Add options messaging
	helpMessage += fmt.Sprintf(bold("OPTIONS:")+"\n%v\n%s%v", options, space, optionsDesc)

	// Check arg for appname to load
	if len(os.Args) != 2 {
		fmt.Println(helpMessage)
		os.Exit(0)
	}

	protoPath := os.Args[1]

	// run protoc command (protoc --go_out=plugins=grpc:. $proto)
	err := executil.CmdStart("protoc", "--go_out=plugins=grpc:.", protoPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(bold("SUCCESS: ") + protoPath + " pb.go successfully created.")

}

func bold(s string) string {
	return "\033[1m" + s + "\033[0m"
}
