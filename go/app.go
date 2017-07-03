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
	// a_count := utf8.RuneCountInString(a)
	b := r.FormValue("b")
	// b_count := utf8.RuneCountInString(b)
	// result, pre, post := ""
	// pre_count := 0
	// if a_count >= b_count {
	// 	pre = b
	// 	post = a
	// 	pre_count = b_count
	// } else {
	// 	pre = a
	// 	post = b
	// 	pre_count = a_count
	// }
	// pre_char := []string{}
	// for pos, c:= range pre {
	// 	pre_char.append(pre_char, string(c))
	// }
	result := ""
	b_char := []string{}
	for _, c := range b {
		b_char = append(b_char, string(c))
	}
	// last_index := 0

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
