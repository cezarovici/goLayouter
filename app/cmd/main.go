package main

import (
	"fmt"
	"os"

	"github.com/cezarovici/goLayouter/app/service"
	"github.com/cezarovici/goLayouter/app/service/render"
	"github.com/cezarovici/goLayouter/domain/line"
	"github.com/cezarovici/goLayouter/helpers"
)

// It reads the file source from the command line
// and then creates a service that will render the
//	contents of the file source	.

// If any error occurs, the application will exit
//	with a non-zero exit code.

// The exit codes are:
// 1 - no file source provided
// 2 - failed to read file
// 3 - failed to create lines from file
// 4 - failed to create service
// 5 - failed to render service

const (
	ExitCodeSuccess = iota
	ExitCodeNoFileSource
	ExitCodeInvalidArgs
	ExitCodeParsingContent
	ExitCodeCreateService
	ExitCodeRender
)

const (
	MinimumArgs = 2
	FirstArg    = 1
)

func main() {
	if len(os.Args) < MinimumArgs {
		fmt.Println("Error: no file source provided")
		os.Exit(ExitCodeNoFileSource)
	}

	fileSource := os.Args[FirstArg]
	content, err := helpers.ReadFile(fileSource)
	if err != nil {
		fmt.Printf("Error: failed to read file: %v\n", err)
		os.Exit(ExitCodeInvalidArgs)
	}

	lines, err := line.NewLines(content)
	if err != nil {
		fmt.Printf("Error: failed to create lines from file: %v\n", err)
		os.Exit(ExitCodeParsingContent)
	}

	items := lines.ToItems()
	serv, errNewService := service.NewService(*items, render.Funcs)
	if errNewService != nil {
		fmt.Printf("Error: failed to create service: %v\n", err)
		os.Exit(ExitCodeCreateService)
	}

	if err := serv.Render(); err != nil {
		fmt.Printf("Error: failed to render service: %v\n", err)
		os.Exit(ExitCodeRender)
	}
}
