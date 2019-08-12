package main

import (
	"fmt"
	"runtime"

	"github.com/afansv/unlim/pkg/win"
)

func main() {
	switch runtime.GOOS {
	case "windows":
		if err := win.SetIPv4TTL(); err != nil {
			panic(err)
		}

		if err := win.DisableIPv6(); err != nil {
			panic(err)
		}
		fmt.Println("OK! Now restart your computer")
	default:
		panic("Windows only yet")
	}
}
