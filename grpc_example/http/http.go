package main

import (
	"context"
	"flag"
	"net/http"
	"path"
	"strings"

	"github.com/golang/glog"

	gw "../helloworld"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

const (
	port = ":80"
)

var (
	echoEndpoint = flag.String("endpoint", "localhost:9090", "this is gRPC tcp address")
)

func newGw(ctx context.Context, opts ...runtime.ServeMuxOption) (http.Handler, error) {
	mux := runtime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterGreeterHandlerFromEndpoint(ctx, mux, *echoEndpoint, dialOpts)
	if err != nil {
		return nil, err
	}
	return mux, nil
}

func run(address string, opts ...runtime.ServeMuxOption) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := http.NewServeMux()
	mux.HandleFunc("/swagger/", func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, ".swagger.json") {
			glog.Errorf("Not Found: %s", r.URL.Path)
			http.NotFound(w, r)
			return
		}

		glog.Infof("Serving %s", r.URL.Path)
		p := strings.TrimPrefix(r.URL.Path, "/swagger/")
		p = path.Join("../helloworld", p)
		http.ServeFile(w, r, p)
	})
	gw, err := newGw(ctx, opts...)
	if err != nil{
		glog.Fatal(err)
	}
	mux.Handle("/", gw)
	return http.ListenAndServe(address, mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()
	if err := run(port); err != nil {
		glog.Fatal(err)
	}
}
