package utils

import (
	"fmt"
	"os"
	env "github.com/joho/godotenv"
)

func LoadEnv(){
	env.Load()
	fmt.Println(os.Getenv("PORT"))
}