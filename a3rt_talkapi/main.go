package main

import (
	//"encoding/json"
	"net/url"
	"os"
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
)

const ENDPOINT = "https://api.a3rt.recruit-tech.co.jp/talk/v1/smalltalk"

type Results struct {
	Perplexity string `json:"perplexity"`
	Reply string `json:"reply"`
}

type Responce struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Results Results `json:"results"`
}

func main() {
	apikey := os.Getenv("APIKEY")
	/*
	que := "おはよう"
	q := map[string]string{
		"apikey": apikey,
		"query": que,
	}
	*/

	values := url.Values{}
	values.Add("apikey", apikey)
	values.Add("query", "おはよう")

	// http.PostForm(ENDPOINT, values)の中身
	// https://api.a3rt.recruit-tech.co.jp/talk/v1/smalltalk?apikey=apikey&query=おはよう
	res, err := http.PostForm(ENDPOINT, values)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(string(body))
	/*
	var responce Responce
	err = json.Unmarshal(body, &responce)

	if err != nil {
		log.Fatal("*json.Unmarshal*\n", err)
	}
	*/
}