/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/0ya-sh0/keyforge/cmd"
	"github.com/0ya-sh0/keyforge/internal/logger"
)

func main() {
	defer logger.Sync()
	logger.Info("Starting KeyForge")
	cmd.Execute()
}
