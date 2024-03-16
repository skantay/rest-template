package main

import "github.com/skantay/service/internal/apps/app"

func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}
