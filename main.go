package main

import (
	_ "flower/controller"
	"flower/mysql"
	"flower/router"
	"flower/server"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("result ")
		}
	}()
	// 初始化配置
	initMysqlDb()
	// 初始化路由
	router.Do()
	// 启动服务
	StartServer()
}

func initMysqlDb() {
	err := mysql.Init("root:linym6303763!@tcp(123.207.1.119:3306)/flower?charset=utf8mb4")
	if err != nil {
		// TODO
	}
}

func StartServer() {
	s := server.Server{
		Port: 8089,
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	signalCh := make(chan os.Signal, 1)
	go func() {
		for {
			sig := <-signalCh
			if sig == syscall.SIGINT || sig == syscall.SIGTERM {
				err := s.Shutdown()
				if err != nil {
					// TODO
				}
				Shutdown()
				wg.Done()
			}
		}
	}()
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	s.Start()
	wg.Wait()
}

// 关闭资源
func Shutdown() {
	fmt.Printf("close")
}
