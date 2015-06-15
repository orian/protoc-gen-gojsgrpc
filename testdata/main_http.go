package main

import (
	"log"
	"net/http"
	"fmt"

	"github.com/orian/jsgrpc"
	"github.com/orian/protoc-gen-gojsgrpc/testdata/jsgrpc/testing"
	"golang.org/x/net/context"
)

type X struct{}

func (*X) UnaryCall(ctx context.Context, in *testing.SimpleRequest) (*testing.SimpleResponse, error) {
	s := &testing.SimpleResponse{
		X: 1,
	}
	return s, nil
}

func (*X) Math(ctx context.Context, in *testing.Args) (*testing.Result, error){
	x := float32(0.0)
	switch in.Op {
	case testing.Args_NONE:
		return nil, fmt.Errorf("Cannot perform NONE")
	case testing.Args_SUM:
		for _,v:= range in.Nums {
			x += v
		}
	case testing.Args_MULT:
		x = 1.0
		for _,v:= range in.Nums {
			x *= v
		}
	default:
		return nil, fmt.Errorf("NotImplemented.")
	}
	return &testing.Result{Res: x}, nil
}

func main() {
	s := jsgrpc.NewServer()
	testing.RegisterTestServer(s, &X{})

	m := http.NewServeMux()
	hs := &http.Server{
		Handler: m,
	}
	hs.Addr = ":8080"
	m.Handle("/api/", http.StripPrefix("/api", s))
	log.Printf("Starting server on %s", hs.Addr)
	if err := hs.ListenAndServe(); err != nil {
		log.Printf("err: %s", err)
	}
}
