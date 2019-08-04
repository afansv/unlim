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
			panic("Error editing the Windows registry. Please make sure that the program is running as administrator.")
		}
		fmt.Println("OK!")
	default:
		panic("Windows only yet")
	}
}
