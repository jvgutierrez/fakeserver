package main

import (
	"flag"
	"html/template"
	"math/rand"
	"net/http"
)

const tpl = `
<html><head><title>{{.Title}}</title></head>
<body>
<h1>You're visiting {{.Title}}</h1>
<h2>This a litle fake page served by fakeserver</h2>
<p>And now some links to be crawled</p>
{{range .Links}}<a href="/{{.}}" />{{.}}</a>{{end}}
</body>
</html>`

type Page struct {
	Title string
	Links []int
}

type FakeServer struct {
	addr      string
	linksPage int
	maxLinks  int
	template  *template.Template
}

func (fs *FakeServer) Handler(w http.ResponseWriter, r *http.Request) {
	var nlinks = rand.Intn(fs.linksPage)
	p := Page{Title: r.URL.Path}
	p.Links = make([]int, nlinks)
	for i := 0; i < nlinks; i++ {
		p.Links = append(p.Links, rand.Intn(fs.maxLinks))
	}
	fs.template.Execute(w, p)
}

func (fs *FakeServer) Init() {
	http.HandleFunc("/", fs.Handler)
	http.ListenAndServe(fs.addr, nil)
}

func NewFakeServer(addr string, linksPage, maxLinks int) *FakeServer {
	return &FakeServer{
		addr:      addr,
		linksPage: linksPage,
		maxLinks:  maxLinks,
		template:  template.Must(template.New("page").Parse(tpl)),
	}
}

func main() {
	var addr = flag.String("a", ":8080", "Listening address")
	var linksPage = flag.Int("lp", 25, "Max. links per page")
	var maxLinks = flag.Int("ml", 5000, "Max. different links")
	flag.Parse()
	fakeServer := NewFakeServer(*addr, *linksPage, *maxLinks)
	fakeServer.Init()
}
