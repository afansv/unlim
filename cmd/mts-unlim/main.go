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
			fmt.Println("Error editing the Windows registry. Please make sure that the program is running as administrator.")
		}
	default:
		fmt.Println("Windows only yet")
	}
}
