package main

import (
	"fmt"
	"group-24-ECE461/internal/logger"
	"group-24-ECE461/internal/parser"
	"os"
)

func main() {

	logger, err := logger.InitLogger()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer logger.Sync()

	logger.Info("Starting Application")
	fmt.Println("Starting Application")
	argsWithOutProg := os.Args[1:]
	parser.ParseArguments(argsWithOutProg, logger)
}
