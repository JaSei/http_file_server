package main

import (
	"fmt"
	"log"

	"github.com/JaSei/pathutil-go"
	"github.com/alecthomas/kingpin"
	"github.com/valyala/fasthttp"
)

var (
	port = kingpin.Flag("port", "server port").Default("8080").Int()
	dir  = kingpin.Flag("dir", "prefix path").Required().ExistingDir()
)

func main() {
	kingpin.Parse()

	log.Fatal(fasthttp.ListenAndServe(fmt.Sprintf(":%d", *port), func(ctx *fasthttp.RequestCtx) {
		path, err := pathutil.NewPath(*dir, string(ctx.Path()))

		if err != nil {
			ctx.Error(err.Error(), 500)
		} else {
			if path.IsDir() {
				ctx.Error(fmt.Sprintf("Requested path %s is directory", path), 406)
			} else {
				fasthttp.ServeFile(ctx, path.Canonpath())
			}
		}
	}))
}
