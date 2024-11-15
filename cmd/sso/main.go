package main

import (
	"fmt"

	"github.com/kirinyoku/grpc-auth-sso/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
}
