package main

import (
	
  "fmt"
  "log"
   "github.com/anugoli05/MyGoWSwithEnv/.MyPackages"
)

func main() {
 
  config, err := util.LoadConfig(".")

 if err!=nil{
 log.Fatal("cannot  load config: ", err)
}
fmt.Printf("Testenvironment: %s", config.TestEnv)

}