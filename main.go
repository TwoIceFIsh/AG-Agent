package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	for {
		fmt.Println("aa")
		resp, err := http.Get("https://www.naver.com")
		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()

		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println("OK")

		time.Sleep(5 * time.Second)

	}
}
