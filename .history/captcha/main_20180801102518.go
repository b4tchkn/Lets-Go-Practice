package main

import (
	"fmt"
	"github.com/haisum/recaptcha"
	"log"
	"net/http"
)

func main() {
	sitekey := "6Ldui2cUAAAAALg5fTPHG8iv7B1HCa1Zxh4Xf8XG"
	//sitekey := "6LcKjGYUAAAAAJT786TtMrlWSNWNFvD4Z1y1Vm04"
	re := recaptcha.R{
		Secret: "6LcKjGYUAAAAAGlp3EfFvTxExxIu9t7J3udI9NjZ",
	}

	form := fmt.Sprintf(`
		<html>
			<head>
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, form)
	})
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
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
