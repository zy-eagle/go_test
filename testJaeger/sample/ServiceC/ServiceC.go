package main

import (
	"fmt"
	"net/http"

	"myProject/testJaeger/tracing"
)

func main() {
	fmt.Println("===service C start===")
	tracing.InitTracer("service-C")

	http.HandleFunc("/serviceC/api/test", func(w http.ResponseWriter, r *http.Request) {
		traceid := r.URL.Query().Get("traceid")
		span, _, _ := tracing.StartSpan(r.RequestURI, traceid, false)

		tracing.FinishSpan(span)
	})

	fmt.Println(http.ListenAndServe(":9993", nil))
}
