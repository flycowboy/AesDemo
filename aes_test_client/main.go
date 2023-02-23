package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	go func() {
		for {
			testApi()
			time.Sleep(time.Second * 2)
		}
	}()
	select {}
}

func testApi() {
	url := "http://localhost:8080/json"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	fmt.Println("Client get url: ", url)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	dataStr := string(body)
	fmt.Println(res.Status)
	fmt.Println("get response:", dataStr)
	fmt.Println("decrypt data: " + KeyDecrypt("SecretKey", dataStr))
	fmt.Println("--------------------------------------------------")
}
