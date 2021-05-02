package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kratos/kratos/pkg/sync/errgroup"
)

func main() {

	g := errgroup.WithContext(context.Background())

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "pong")
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	g.Go(func(ctx context.Context) error {
		err := server.ListenAndServe()
		fmt.Println("goroutine1 exit")
		return err
	})

	g.Go(func(ctx context.Context) error {
		sign := make(chan os.Signal)
		signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)
		var err error
		select {
		case <-ctx.Done():
			err = ctx.Err()
		case <-sign:
			err = server.Shutdown(ctx)
		}
		fmt.Println("goroutine2 exit")
		return err
	})

	fmt.Println(g.Wait())
}
