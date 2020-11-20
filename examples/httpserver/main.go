package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gustavohenrique/gometrics"
	"github.com/gustavohenrique/gometrics/lib/util"
)

var collector *gometrics.Collector

func main() {
	collector = gometrics.New()
	http.HandleFunc("/", getMetrics)
	log.Println("Listening :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getMetrics(w http.ResponseWriter, r *http.Request) {
	metrics, _ := collector.Metrics()
	data := fmt.Sprintf(`{"data": "%s"}`, util.PrettyJSON(metrics))
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"message": "Method not allowed"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(data))
}
