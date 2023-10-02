package main

import (
	"fmt"
	"github.com/kevinmichaelchen/simple-go-118-workspace/bar"
	"github.com/rs/zerolog"
)

func main() {
	fmt.Println("hi")
	bar.Bar()
	l := zerolog.Logger{}
	l.Info().Msg("hi")
}
