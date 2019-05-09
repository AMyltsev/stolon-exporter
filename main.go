package main

import (
	"fmt"
	"net/http"

	"github.com/juju/gnuflag"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
)

func init() {
	// regMeteis()
}

func param() params {
	params := params{}
	gnuflag.StringVar(&params.backend, "b", "etcd",
		"Stolon backend type: etcd, consul, k8s")
	gnuflag.StringVar(&params.port, "p", "9779",
		"Exporter port")
	gnuflag.Parse(true)
	return params
}

func main() {
	param := param()

	log.Infoln("Starting openstack-exporter")
	log.Infoln("Path of metrics /metrics")
	log.Infoln("Port to listen on for telemetry", param.port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Use /metrics\n")
	})

	http.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(":"+param.port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
