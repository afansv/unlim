package main

import (
	"fmt"
	"runtime"

	"github.com/afansv/mts-unlim/pkg/win"
)

func main() {
	switch runtime.GOOS {
	case "windows":
		if err := win.SetTTL(); err != nil {
			panic(err)
		}

		if err := win.DisableIPv6(); err != nil {
			panic(err)
		}
		fmt.Println("OK!")
	default:
		panic("Windows only yet")
	}
}