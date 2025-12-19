package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"tron-signal/internal/app"
	"tron-signal/internal/logx"
)

func main() {
	logx.Init()

	log.Println("[BOOT] tron-signal starting")

	a, err := app.New()
	if err != nil {
		log.Fatalf("[FATAL] app init failed: %v", err)
	}

	if err := a.Start(); err != nil {
		log.Fatalf("[FATAL] app start failed: %v", err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	log.Println("[SHUTDOWN] signal received")
	a.Stop()
	log.Println("[SHUTDOWN] clean exit")
}
