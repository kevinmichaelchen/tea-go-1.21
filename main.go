package main

import (
	"fmt"
	"github.com/kevinmichaelchen/fx-libs/gqlgen"
	"github.com/rs/zerolog"
)

func main() {
	fmt.Println("hi")
	gqlgen.NewRelicMiddleware()
	l := zerolog.Logger{}
	l.Info().Msg("hi")
}
