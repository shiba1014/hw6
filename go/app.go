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

var trainList = make(map[string][]string)
var routes = make([][]string, 0)

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

	makeAdjacencyList(tracks)

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
	from := r.FormValue("from")
	to := r.FormValue("to")
	fmt.Fprintf(w, `<font size="+3">%sから%sまでの乗り換え案内</font><br>`, from, to)
	// fmt.Fprintf(w, `%s`,trainList)
	// fmt.Fprintf(w, `%s<br>`,searchRoute(from, to))
	searchRoute(from, to)
	bestWay := searchBestRoute()
	fmt.Fprintf(w, `%s<br>`, bestWay)
	routes = nil
}

func makeAdjacencyList(tracks []Track){
	for _, t := range tracks {
		for pos, s := range t.Stations {
			list := trainList[s]
			if pos > 0 {
				if contains(list, t.Stations[pos - 1]) == true {
					continue
				}
				list = append(list, t.Stations[pos - 1])
			}
			if pos < len(t.Stations) - 1 {
				if contains(list, t.Stations[pos + 1]) == true {
					continue
				}
				list = append(list, t.Stations[pos + 1])
			}
			trainList[s] = list
		}
	}
}

func searchRoute(from, to string) {
	path := make([]string, 0)
	dfs(to, append(path, from))
}

func dfs(to string, path []string) {
	current := path[len(path) - 1]
	if current == to {
		answer := make([]string, len(path))
		copy(answer, path)
		routes = append(routes, answer)
	} else {
		for _, x := range trainList[current] {
			if !contains(path, x) {
				dfs(to, append(path, x))
			}
		}
	}
}

func contains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

func searchBestRoute() []string {
	var minRoute = routes[0]
	for _, r := range routes {
		if len(minRoute) > len(r) {
			minRoute = r
		}
	}
	return minRoute
}
