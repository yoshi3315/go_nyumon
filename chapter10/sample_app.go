package main

import (
	"fmt"
	"net/http"
)

type serverHTML struct {
	title string
	body  string
}

func (s *serverHTML) ServeHTTP(w http.ResponseWriter, h *http.Request) {
	h := `
	<html>
		<head><title>` + s.title + `</title></head>
		<body>` + s.body + `</body>
	</html>
	`

	fmt.Fprint(w, h)
}

func main() {
	helloHTML := &serverHTML{"Hello Title", "Hello"}
	goodbyeHTML := &serverHTML{"Goodbye Title", "Goodbye"}

	http.Handle("/hello", helloHTML)
	http.Handle("/goodbye", goodbyeHTML)

	if err := http.ListenAndServe(":4000", nil); err != nil {
		fmt.Println(err)
	}
}
