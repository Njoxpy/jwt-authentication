package initializers

import (
	"log"
    "github.com/joho/godotenv"
)
// start the name of the function in capital letter so that they can be used in other packages
func LoadEnvironmentVariables(){
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
  
}