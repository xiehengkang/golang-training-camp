package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
)




func main() {
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		return httpHandle(ctx, ":8000")
	})
	g.Go(func() error {
		return httpHandle(ctx, ":8080")
	})
	if err := g.Wait(); err != nil {
		fmt.Printf("wait over: %v\n", err)
	}
}


func httpHandle(ctx context.Context, addr string) error {
	var err error
	mux := http.NewServeMux()
	done := make(chan int)
	mux.HandleFunc("/test", func(w http.ResponseWriter, req *http.Request) {
		done <- 1
	})
	server := http.Server{Addr: addr, Handler: mux}
	go func() {
		select {
		case <-ctx.Done():
			fmt.Printf("ctx over,err is :%v", ctx.Err())
		case <-done:
			fmt.Printf("server[%s] done", addr)
		}
	}()
	err = server.Shutdown(context.Background())
	return err

}