package main

import (
	"log"

	examples "github.com/hiro-o918/go-bazel-examples"
)

func echo(e examples.Echo, msg string) string {
	return e.Echo(msg)
}

func main() {
	e := examples.NewEcho()
	log.Println(echo(e, "foo"))
}
