package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// encrypted := "lIR3JIHpomC5Zm8sjy29D/xFcXUX0c/4vQ=="

	// fmt.Println(encrypted)
	// fmt.Println(keyDecrypt("SecretKey", encrypted))

	http.HandleFunc("/json",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("Server get a request \n")

			responseJson := string("{\"a\":1}")
			responseStr := KeyEncrypt("SecretKey", responseJson)
			fmt.Println("send response: ", responseStr)
			fmt.Println("--------------------------------------------------")
			io.WriteString(w, responseStr)
		})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("err: " + err.Error())
	}
}
