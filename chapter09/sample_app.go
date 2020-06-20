package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonData struct {
	A string
	B string `json:"code"`
	C int    `json:",string"`
	D string `json:",omitempty"`
	E string `json:"-"`
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	data := JsonData{A: "データ1", B: "データ2", C: 2, D: "", E: "データ5"}
	j, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprint(w, string(j))
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	h := `
	<html>
		<head><title>JSON</title></head>
		<body>
			<a href="#" onclick="reqJson()">JSON取得</a>
			<script>
			  const reqJson = () => {
					let xhr = new XMLHttpRequest()
					xhr.open("GET", "/json")
					xhr.onload = () => { alert(xhr.responseText) }
					xhr.send()
				}
			</script>
		</body>
	</html>
	`

	fmt.Fprint(w, h)
}

func main() {
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/html", htmlHandler)

	if err := http.ListenAndServe(":4000", nil); err != nil {
		fmt.Println(err)
	}
}
