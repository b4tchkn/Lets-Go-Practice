package main

import (
	"net/url"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"log"
)

type handler struct {}

type recaptchaResponce struct {
	Success    bool `json:"success"`
	ChallengeTs time.Time `json:"challenge_ts"`
	Hostname string `json:"hostname"`
	ErrorCodes []string `json:"error-codes"`
}

var postURL = "https://www.google.com/recaptcha/api/siteverify"

var (
	sitekey = "6LcKjGYUAAAAAJT786TtMrlWSNWNFvD4Z1y1Vm04"
	secrect = "6LcKjGYUAAAAAGlp3EfFvTxExxIu9t7J3udI9NjZ"
)

//requestポインタをstringにcast
func castResponce(req http.Request) string {
	responce := req.FormValue("g-recaptcha-response")
	return responce
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	<-time.After(20 * time.Second)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello!"))
	log.Println("ok")
}

func main() {
//HTML記述
	form := fmt.Sprintf(`
		<html>
			<head>
			<title>reCAPTCHA</title>
				<script src='https://www.google.com/recaptcha/api.js'></script>
			</head>
			<body>
				<form action="/result" method="post">
					<div class="g-recaptcha" data-sitekey="%s"></div>
					<input type="submit">
				</form>
			</body>
		</html>
	`, sitekey)

	//認証画面出すハンドラ
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, form)
	})

	//timepout
	var h handler
	http.Handle("/timeout", http.TimeoutHandler(&h, 15 * time.Second, "timeout!!"))
	//http.Handle("/timeout", &h)

	//認証結果画面だすハンドラ
	http.HandleFunc("/result", func(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{Timeout: time.Duration(15) * time.Second}
	responce := castResponce(*r)
	res, err := client.PostForm(postURL, url.Values{"secret": {secrect}, "response": {responce}})

	if err != nil {
		log.Printf("Error.")
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error.")
	}
	rr := new(recaptchaResponce)
	err = json.Unmarshal(body, rr)
	if err != nil {
	log.Printf("Error.")
}

//出力用のデータ準備
year, month, day := time.Now().Date()
hour, min, sec := time.Now().Clock()

fmt.Fprintf(w, "RESULT → %v", rr.Success)
fmt.Println("--------------------------------")
fmt.Printf("Success = %v\n", rr.Success)
fmt.Printf("Date = %v-%v-%v\n", year, month, day)
fmt.Printf("Time = %v:%v:%v\n",hour, min, sec)
fmt.Printf("ErrorCodes = %v\n", rr.ErrorCodes)
fmt.Printf("Hostname = %v\n", rr.Hostname)
fmt.Println("--------------------------------")
})

	log.Printf("\n Starting server on http://localhost:8100 . Check example by opening this url in browser.\n")

	err := http.ListenAndServe(":8100", nil)

	if err != nil {
		log.Fatalf("Could not start server. %s", err)
	}
}