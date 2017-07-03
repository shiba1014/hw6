package app

import (
	"net/http"
	"fmt"
	"appengine"
	"appengine/urlfetch"
	"encoding/json"
	"io/ioutil"
)

type Track struct {
	Name string `json:"Name"`
	Stations []string `json:"Stations"`
}

func init() {
	http.HandleFunc("/pata", handlePata)
	http.HandleFunc("/transfer", handleTransfer)
	http.HandleFunc("/search", handleSearch)
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

	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)
	resp, err := client.Get("http://fantasy-transit.appspot.com/net?format=json");
	if err != nil {
		http.Error(w,  err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err:= ioutil.ReadAll(resp.Body);
	if err != nil {
		return
	}

	var tracks []Track
	if err := json.Unmarshal(body, &tracks); err != nil {
		return;
	}

	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html>
		<head>
		  <title>乗り換え案内</title>
		</head>
		<body>
		<form action="/search">
			出発:
			<select name="from"><br>
			`)

	for _, t := range tracks {
		fmt.Fprintf(w, `<option disabled>%s</option>`, t.Name)
		for _, s := range t.Stations {
			fmt.Fprintf(w, `<option>%s</option>`, s)
		}
	}

	fmt.Fprintf(w, `
			</select>
			<br>
			到着:
			<select name="to"><br>
			`)
	for _, t := range tracks {
		fmt.Fprintf(w, `<option disabled>%s</option>`, t.Name)
		for _, s := range t.Stations {
			fmt.Fprintf(w, `<option>%s</option>`, s)
		}
	}

	fmt.Fprintf(w, `
			</select>
			<br>
		  <input type=submit value"乗り換え案内">
		</form>
		</body>
		</html>
		`)
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `result`)
}
