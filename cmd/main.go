package main

import (
	"fmt"
	"github.com/tuusuario/prj-go/internal/macparser"
)

func main() {
	mac := "00:1A:2B:3C:4D:5E"
	info, err := macparser.DetectAndNormalize(mac)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Tipo:", info.MACType)
}
