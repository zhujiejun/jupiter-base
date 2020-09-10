package main

import (
    "fmt"
	"log"
	"jupiter-base/internal/app/engine"
	"jupiter-base/internal/app/model"
    "jupiter-base/internal/app/service"
    "github.com/douyu/jupiter"
)

func main() {
	eng := engine.NewEngine()
	eng.RegisterHooks(jupiter.StageAfterStop, func() error {
        fmt.Println("exit jupiter app ...")
        return nil
      })

    model.Init()
    service.Init()
    if err := eng.Run(); err != nil {
    	log.Fatal(err)
    }
}

