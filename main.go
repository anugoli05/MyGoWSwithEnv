package main

import (
	
  "fmt"
  "log"
  "os"
  // "strings"
   "github.com/anugoli05/MyGoWSwithEnv/.MyPackages"
)

func main() {
  // os.Setenv("TEST_ENVIRONMENT", "VALQA")
  // for _, env := range os.Environ() {
	// 	// env is
	// 	envPair := strings.SplitN(env, "=", 2)
	// 	key := envPair[0]
	// 	value := envPair[1]

	// 	fmt.Printf("%s : %s\n", key, value)
  // }
  // driver, ok := os.LookupEnv("TEST_ENVIRONMENT")
	// if !ok {
	// 	fmt.Println("TEST_ENVIRONMENT is not present")
	// } else {
	// 	fmt.Printf("TEST_ENVIRONMENT: %s\n", driver)
	// }

  config, err := util.LoadConfig(".")

 if err!=nil{
 log.Fatal("cannot  load config: ", err)
}
fmt.Printf("\nTestenvironment: %s", config.TestEnv)
fmt.Printf("\nURL: %s", config.URL)
fmt.Printf("\nTEST_ENVIRONMENT from main file: %s\n", os.Getenv("TEST_ENVIRONMENT"))

}