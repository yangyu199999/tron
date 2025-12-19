package app

import (
	"log"
	"os"
	"path/filepath"

	"tron-signal/internal/config"
	"tron-signal/internal/logx"
	"tron-signal/internal/rule"
)

type App struct{}

func New() (*App, error) {
	initDirs()

	// 每次启动都记录一次重启（手动/异常不区分）
	logx.LogRestart()

	if err := config.Load(); err != nil {
		return nil, err
	}

	// 强制重置 ON / OFF 计数器
	rule.ResetAll()

	return &App{}, nil
}

func (a *App) Start() error {
	log.Println("[APP] started")
	return nil
}

func (a *App) Stop() {
	log.Println("[APP] stopping")
}

func initDirs() {
	dirs := []string{
		"data",
		"data/config",
		"data/storage",
		"data/logs",
	}

	for _, d := range dirs {
		if err := os.MkdirAll(filepath.Clean(d), 0755); err != nil {
			panic(err)
		}
	}
}
