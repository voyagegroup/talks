package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tmpl, err := template.New("").Parse(`<html>
<head>
	<meta charset="UTF-8">
	<title>{{.Title}}</title>
</head>
<body>
{{.Body}}
</body>
</html>
`)
	if err != nil {
		log.Fatalf("compile template failed: %s", err)
	}
	if err := tmpl.Execute(os.Stdout, map[string]interface{}{
		"Title": "cool title",
		"Body":  "here is body",
	}); err != nil {
		log.Fatalf("execute template failed: %s", err)
	}
}
