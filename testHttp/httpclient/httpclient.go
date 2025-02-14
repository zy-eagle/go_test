package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://news.baidu.com/", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	req.Header.Add("User-Agent", "myClient")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	data, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(data))

	defer resp.Body.Close()
}
