package app

import (
	"net/http"
	"fmt"
)

func init() {
	http.HandleFunc("/", handlePata)
}

func handlePata(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// fmt.Fprintf(w, "Hello world!\n")
	a := r.FormValue("a")
	b := r.FormValue("b")
	result := ""
	// count := 0
	// if a >= b {
	// 	count = b
	// } else {
	// 	count = a
	// }
	// var result bytes.Buffer
	for pos, c:= range a {
		result += string(c)
		result += string(b[pos])
	}

	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html>
		<head>
		  <title>パタトクカシーー</title>
		</head>
		<body>`)

		fmt.Fprintf(w, `
			<font size="+10">%s</font>
			`, result)

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
