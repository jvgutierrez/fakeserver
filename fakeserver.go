package main

import (
	"html/template"
	"math/rand"
	"net/http"
)

type Page struct {
	Title string
	Links []int
}

func Handler(w http.ResponseWriter, r *http.Request) {
	const tpl = `
<html><head><title>{{.Title}}</title></head>
<body>
<h1>You're visiting {{.Title}}</h1>
<h2>This a litle fake page served by fakeserver</h2>
<p>And now some links to be crawled</p>
{{range .Links}}<a href="/{{.}}" />{{.}}</a>{{end}}
</body>
</html>`
	var nlinks = rand.Intn(25)
	p := Page{Title: r.URL.Path}
	p.Links = make([]int, nlinks)
	for i := 0; i < nlinks; i++ {
		p.Links = append(p.Links, rand.Intn(5000))
	}
	t, _ := template.New("page").Parse(tpl)
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}
