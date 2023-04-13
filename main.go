package main

import (
	"log"

	"github.com/BautistaBianculli/HelloFiber/bd"
	"github.com/BautistaBianculli/HelloFiber/handlers"
)


func main (){
	if !bd.ConenectionCheck() {
		log.Fatal("No Conecction")
		return
	}
	handlers.Handlers()
}