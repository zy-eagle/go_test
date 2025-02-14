package main

import (
	"fmt"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	otlog "github.com/opentracing/opentracing-go/log"
	"log"
	Jaeger "myProject/testJaeger"
	"net/http"
)

func main() {

	tracer, closer := Jaeger.InitJaeger("formatter")
	defer closer.Close()

	http.HandleFunc("/format", func(w http.ResponseWriter, r *http.Request) {
		spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		span := tracer.StartSpan("format", ext.RPCServerOption(spanCtx))
		defer span.Finish()

		greeting := span.BaggageItem("greeting")
		if greeting == "" {
			greeting = "Hello"
		}

		helloTo := r.FormValue("helloTo")
		helloStr := fmt.Sprintf("%s, %s!", greeting, helloTo)
		span.LogFields(
			otlog.String("event", "string-format"),
			otlog.String("value", helloStr),
		)
		w.Write([]byte(helloStr))
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
