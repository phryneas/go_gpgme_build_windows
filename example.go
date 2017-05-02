package main

import (
	"github.com/proglottis/gpgme"
	"fmt"
)

func main() {
	ctx, err := gpgme.New()
	if err != nil {
		panic(err)
	}
	defer ctx.Release()

	fmt.Println(ctx.EngineInfo().Version())
}