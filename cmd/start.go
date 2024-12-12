package cmd

import (
	"GaMachine/internal/server"
	"GaMachine/pkg/logger"
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const HTTPPort = "8080"

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start",
	Run: func(cmd *cobra.Command, args []string) {
		routerHandler := server.NewHTTPRouter()

		httpSrv := &http.Server{
			Addr:         ":" + HTTPPort,
			Handler:      routerHandler,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
		}

		go func() {
			if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatal("ListenAndServe error:", err)
			}
		}()
		shutdown(httpSrv)

		fmt.Println("http_start called")

	},
}

func init() {
	RootCmd.AddCommand(startCmd)
}

func shutdown(httpSrv *http.Server) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	select {
	case sig := <-sigs:
		fmt.Printf("%s|||%s \r\n", logger.HTTPPort, fmt.Sprintf("捕获信号signal.Notify,sigs:%v", sig))
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) // 3s没有处理完，则强制关闭
		defer cancel()
		if err := httpSrv.Shutdown(ctx); err != nil {
			fmt.Printf("%s|||%s \r\n", logger.HTTPPort, fmt.Sprintf("捕获信号signal.shutdown,err::%v", err))
		}
		fmt.Printf("%s|||%s \r\n", logger.HTTPPort, "http shutdown...")
	}
	// 设置一个关机信号,程序执行时,如果值存在,则表示即将关机
	// td: 待完成
	time.Sleep(3 * time.Second)
}
