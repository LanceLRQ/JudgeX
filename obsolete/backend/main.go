package main

import (
    "backend/core"
	"log"
	"os"
	"path"
)

func main() {
	wDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	cfile := "config/system.yaml"
	if err := core.LoadServerConfig(path.Join(wDir, cfile)); err != nil {
		log.Fatal(err)
	}
	if err := core.InitJudgeXServer(); err != nil {
		log.Fatal(err)
	}
}