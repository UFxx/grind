package main

import (
	"fmt"

	"github.com/sunsetsavorer/grind/internal/app"
)

func main() {

	app := app.New()

	if err := app.Run(); err != nil {
		fmt.Printf("failed to run app: %v", err)
	}
}
