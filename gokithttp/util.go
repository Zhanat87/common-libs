package gokithttp

import (
	"github.com/Zhanat87/common-libs/encoders"
	kitlog "github.com/go-kit/kit/log"
	kitoc "github.com/go-kit/kit/tracing/opencensus"
	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
)

func GetServerOptions(logger kitlog.Logger) []kithttp.ServerOption {
	return []kithttp.ServerOption{
		kithttp.ServerErrorHandler(kittransport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encoders.EncodeErrorJSON),
		kitoc.HTTPServerTrace(),
	}
}
