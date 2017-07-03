package app

import (
	"net/http"
	"fmt"
)

func init() {
	http.HandleFunc("/", handlePata)
	http.HandleFunc("/Transfer", handleTransfer)
}

func handlePata(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	a := r.FormValue("a")
	b := r.FormValue("b")
	result := ""

	b_char := []string{}
	for _, c := range b {
		b_char = append(b_char, string(c))
	}

	index := 0
	for _, c:= range a {
		result += string(c)
		if index < len(b_char) {
			result += b_char[index]
			index += 1
		}
	}
	if index < len(b_char) {
		for i := index; i < len(b_char); i++ {
			result += b_char[index]
		}
	}

	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html>
		<head>
		  <title>パタトクカシーー</title>
		</head>
		<body>`)

	if result != "" {
		fmt.Fprintf(w, `
			<img src="images/pata.jpg">
			<font size="+10">%s</font>
			<hr>
			`, result)
	}

	fmt.Fprintf(w, `
		  <form>
		    <input type=text name=a><br>
		    <input type=text name=b><br>
		    <input type=submit>
		    </form>
		</body>
		</html>
		`)
}

func handleTransfer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `Hello, World!`)
}
