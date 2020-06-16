package main

import (
	"fmt"
	"net/http"
)

func helloHandle(w http.ResponseWriter, r *http.Request) {
	h := `
		<html>
		  <head><title>Hello</title></head>
			<body>Hello</body>
		</html>
	`

	fmt.Fprint(w, h)
}

func goodbyeHandle(w http.ResponseWriter, r *http.Request) {
	h := `
		<html>
		  <head><title>Goodbye</title></head>
			<body>Goodbye</body>
		</html>
	`

	fmt.Fprint(w, h)
}

func main() {
	http.HandleFunc("/hello", helloHandle)
	http.HandleFunc("/goodbye", goodbyeHandle)

	if err := http.ListenAndServe(":4000", nil); err != nil {
		fmt.Println(err)
	}
}
