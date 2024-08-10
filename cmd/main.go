package main 

import (

	"log"
	"net/http"

)

func main(){
	configs.Loadenv()
	client := configs.ConnectDB()

	log.Println("Server started at: 8080")
	log.Fatal(http.ListenAndServe(":8080",router))
}