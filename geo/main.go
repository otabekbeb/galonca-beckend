package main

import (
	"context"
	"fmt"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"tsl_server/api"
	"tsl_server/service"
)

func main() {

	e := godotenv.Load()
	if e != nil {
		log.Fatalf("Error read enveroment %v", e)
	}
	grpcPort := os.Getenv("grpc_port")
	webPort := os.Getenv("web_port")
	grpcServer := grpc.NewServer(
	//			grpc.UnaryInterceptor(service.NewAuthInterceptor().Unary()),
	//			grpc.StreamInterceptor(service.NewAuthInterceptor().Stream()),
	)
	wrapServer := grpcweb.WrapServer(grpcServer)
	webServer := &http.Server{Addr: ":" + webPort}
	webServer.Handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.ProtoMajor == 2 {
			wrapServer.ServeHTTP(w, req)
			log.Println("HTTP2 request")
			return
		} else {
			log.Println("HTTP1 request\nSet Header")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-User-Agent, X-Grpc-Web")
			w.Header().Set("grpc-status", "")
			w.Header().Set("grpc-message", "")
			if req.Method == http.MethodOptions {
				return
			}
			if wrapServer.IsGrpcWebRequest(req) {
				log.Println("GRPc call")
				wrapServer.ServeHTTP(w, req)
				return
			}
		}
		http.DefaultServeMux.ServeHTTP(w, req)
	})
	registerServices(grpcServer)
	l, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		log.Printf("gRPC Server started on port: %s\n", grpcPort)
		err := grpcServer.Serve(l)
		if err != nil {
			log.Fatal(err)
		}
	}()
	go func() {
		log.Printf("webServer started on port: %s\n", webPort)
		err := webServer.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
	waitExitSignal(grpcServer, webServer)
	log.Println("bye!")
	fmt.Println("shutting down")
}

func registerServices(s *grpc.Server) {
	api.RegisterGeoServiceServer(s, service.NewGeoService())
}

func waitExitSignal(s *grpc.Server, w *http.Server) {
	sigCh := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigCh
		log.Printf("Signal %s\n", sig)
		s.Stop()
		err := w.Shutdown(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		done <- true
	}()
	<-done
}
