package main

import (
	"github.com/cisordeng/beego/xenon"

	_ "leo/model"
	_ "leo/rest"
)

func main() {
	xenon.Run()
}
