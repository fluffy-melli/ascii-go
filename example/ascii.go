package main

import (
	"fmt"
	"log"

	"github.com/fluffy-melli/ascii-go"
)

func main() {
	image, err := ascii.ReadImage("./golang.png")
	if err != nil {
		log.Fatalln(err)
	}
	aci := ascii.Render(image, 150)
	fmt.Println(aci.ToStr())
}
