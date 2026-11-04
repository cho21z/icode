package main

import (
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	logger, _ = zap.NewProduction()
}

func fastHTTPHandle(ctx *fasthttp.RequestCtx) {
	logger.Info("Hello go module", zap.ByteString("uri", ctx.RequestURI()))
}

func main()  {
	fasthttp.ListenAndServe(":8081", fastHTTPHandle)
}