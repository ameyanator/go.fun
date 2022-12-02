package main

import (
	"net/http"
	"text/template"
	"time"
)

var tpl = `<!DOCTYPE HTML>
<html>
<head>
<meta charset="utf-8">
<title>Date Example</title>
</head>
<body>
<p>{{.Date | dateFormat "2-Jan-2006"}}</p>
</body>
</html>`

var funcMap = template.FuncMap{
	"dateFormat": dateFormat,
}

func dateFormat(layout string, d time.Time) string {
	return d.Format(layout)
}

func serveTemplate(res http.ResponseWriter, req *http.Request) {
	t := template.New("date")
	t.Funcs(funcMap)
	t.Parse(tpl)
	data := struct{ Date time.Time }{
		Date: time.Now(),
	}
	t.Execute(res, data)
}

func main() {
	http.HandleFunc("/", serveTemplate)
	http.ListenAndServe(":8080", nil)
}
