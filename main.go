package main

import (
	"fmt"
	"log"
	"os"
)

func  main()  {



	 portString:= os.Getenv("PORT")


	 if portString==""{
		log.Fatal("PORT Not Found")
	 }
	fmt.Println("Hello Go")
}