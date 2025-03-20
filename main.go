package main

import "github.com/MaurogDark/gator/internal/config"

func main() {
	cfg, err := config.Read()
	if err {
		print("Error: ", err)
	}
	cfg.SetName("Maurog")
	err = config.Write(cfg)
	if err {
		print("Error: ", err)
	}
	cfg, err = config.Read()
	if err {
		print("Error: ", err)
	}
	print(cfg)
}
