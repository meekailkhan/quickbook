package main

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)


func main(){
	if _,err := os.Stat(".env"); errors.Is(err,os.ErrNotExist){
		os.Create(".env")
	}
	if err := godotenv.Load(); err != nil {
		panic("could not load the env")
	} 
	



}