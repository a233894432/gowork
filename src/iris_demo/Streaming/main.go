package main

import (
	"bufio"
	"fmt"
	"time"

	"github.com/kataras/iris"
)

func main() {
	iris.Any("/stream", func(ctx *iris.Context) {
		ctx.StreamWriter(stream)
	})

	iris.Listen(":8080")
}

func stream(w *bufio.Writer) {
	for i := 0; i < 10; i++ {
		fmt.Fprintf(w, "this is a message number %d", i)

		// Do not forget flushing streamed data to the client.
		if err := w.Flush(); err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}
