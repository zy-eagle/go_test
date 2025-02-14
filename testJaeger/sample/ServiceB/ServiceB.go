package main

import (
	"fmt"
	"net/http"
	"time"

	"go-test/testJaeger/tracing"
)

func main() {
	fmt.Println("===service B start===")
	tracing.InitTracer("service-B")

	ListenHTTPB()
}

func ListenHTTPB() {
	http.HandleFunc("/serviceB/api/test", func(w http.ResponseWriter, r *http.Request) {
		span, traceId, _ := tracing.StartSpan(r.RequestURI, r.Header.Get("traceid"), false)

		go callServiceC(traceId)

		time.Sleep(300 * time.Millisecond)
		tracing.FinishSpan(span)
		w.Write([]byte("serviceB done"))
	})
	fmt.Println(http.ListenAndServe(":9992", nil))
}

func callServiceC(traceId string) {
	req, _ := http.NewRequest("GET", "http://localhost:9993/serviceC/api/test?traceid="+traceId, nil)
	http.DefaultClient.Do(req)
}
