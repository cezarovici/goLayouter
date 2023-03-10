package main

import (
	"fmt"
	"os"

	"github.com/cezarovici/goLayouter/app/services"
	"github.com/cezarovici/goLayouter/app/services/renders"
	"github.com/cezarovici/goLayouter/helpers"
	"github.com/cezarovici/goLayouter/line"
)

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
