package main

import (
	"fmt"
	//"github.com/haisum/recaptcha"
	""
	"log"
	"net/http"
)

func main() {
	sitekey := "6LcKjGYUAAAAAJT786TtMrlWSNWNFvD4Z1y1Vm04"
	re := recaptcha.R{
		Secret: "6LcKjGYUAAAAAGlp3EfFvTxExxIu9t7J3udI9NjZ",
	}

	//HTML記述
	form := fmt.Sprintf(`
		<html>
			<head>
			<title>reCAPTCHA</title>
				<script src='https://www.google.com/recaptcha/api.js'></script>
			</head>
			<body>
				<form action="/submit" method="post">
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
	//認証結果画面だすハンドラ
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		//実際の認証結果はisValidの中（true|false）
		isValid := re.Verify(*r)
		log.Printf("CAPTCHA AUTH RESULT: ", isValid)
		if isValid {
			fmt.Fprintf(w, "成功")
		} else {
			fmt.Fprintf(w, "失敗: %v", re.LastError())
		}
	})

	log.Printf("\n Starting server on http://localhost:8100 . Check example by opening this url in browser.\n")

	err := http.ListenAndServe(":8100", nil)

	if err != nil {
		log.Fatalf("Could not start server. %s", err)
	}
}
