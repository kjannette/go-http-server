package main

import (
    "context"
    "github.com/kjannette/go-server/internal/config"
    "github.com/kjannette/go-server/internal/infrastructure/os"
    "go.uber.org/zap"
)

var (
    appName = ""
    appVarsion = ""
)

func main() {
    defer func() {_ = zap.L().Sync() }()
}
