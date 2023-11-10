package main

import "auth/internal/server"

func main() {
	if err := server.StartGRPCServer(); err != nil {
		panic(err)
	}
}
