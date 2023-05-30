package main

import (
	"api/env"
	"api/pkg/nats"
)

func main() {
	cfg := env.GetConfig()
	cfg.ConfigSanityCheck()
	_, err := nats.New(cfg.NATS.Url)
	if err != nil {
		panic(err)
	}
}
