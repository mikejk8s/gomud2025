package main

import "gomud2025/lib"

func main() {
	err := lib.NewMudServer().Run()
	if err != nil {
		panic(err)
	}
}
