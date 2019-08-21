package clients

import (
	"bufio"
	"fmt"
	"net/http"
)

func getResponseFromNetwork(channel chan *http.Response) {
	res, err := http.Get("http://gobyexample.com")
	if err != nil {
		panic(err)
	}
	channel <- res
}

func PrintValue() {
	channel := make(chan *http.Response)
	go getResponseFromNetwork(channel)
	res := <-channel
	defer res.Body.Close()

	fmt.Println("Response StatusCode: ", res.StatusCode)

	scanner := bufio.NewScanner(res.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
