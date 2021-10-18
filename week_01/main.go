package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/berryhe/cloud-native-curriculum/week_01/transport"
)

func main() {
	transport.Version = os.Getenv("VERSION")
	mux := http.DefaultServeMux
	mux.HandleFunc("/", transport.HandleRootPath)
	mux.HandleFunc("/healthz", transport.HandleHealthz)
	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				panic(err)
			}
		}
	}()

	fmt.Println("server start.....")
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	<-ctx.Done()

	stop()

	fmt.Println("server shutting down gracefully")

	// 给程序最多 5 秒时间处理正在服务的请求
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(timeoutCtx); err != nil {
		fmt.Println(err)
	}

}
