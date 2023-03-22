package main

import (
	"fmt"
	"os"

	"github.com/cezarovici/goLayouter/app/services"
	"github.com/cezarovici/goLayouter/app/services/renders"
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

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: no file source provided")
		os.Exit(1)
	}

	fileSource := os.Args[1]
	content, err := helpers.ReadFile(fileSource)
	if err != nil {
		fmt.Printf("Error: failed to read file: %v\n", err)
		os.Exit(2)
	}

	lines, err := line.NewLines(content)
	if err != nil {
		fmt.Printf("Error: failed to create lines from file: %v\n", err)
		os.Exit(3)
	}

	items := lines.ToItems()
	serv, errNewService := services.NewService(*items, renders.RenderFuncs)
	if errNewService != nil {
		fmt.Printf("Error: failed to create service: %v\n", err)
		os.Exit(4)
	}

	if err := serv.Render(); err != nil {
		fmt.Printf("Error: failed to render service: %v\n", err)
		os.Exit(5)
	}
}
