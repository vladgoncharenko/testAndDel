package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	c := http.Client{
		Timeout: 20 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:9897", nil)

	if err != nil {
		fmt.Println(err)
	}

	resp, err := c.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	response, err := ioutil.ReadAll(resp.Body)

	defer func() {
		err = resp.Body.Close()
	}()

	fmt.Println(string(response))
}
