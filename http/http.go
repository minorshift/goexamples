package main

import (
	"html/template"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {

	//HTML Template
	const tpl = `<!DOCTYPE html>
              <html>
              <head>
                <meta charset="UTF-8">
                <title>{{.Title}}</title>
                <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap.min.css" integrity="sha384-1q8mTJOASx8j1Au+a5WDVnPi2lkFfwwEAa8hDDdjZlpLegxhjVME1fgjWPGmkzs7" crossorigin="anonymous">
                <script async src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/js/bootstrap.min.js" integrity="sha384-0mSbJDEHialfmuBBQP6A4Qrprq5OVfW37PRR3j5ELqxss1yVqOtnepnHVP9aJ7xS" crossorigin="anonymous"></script>
              </head>
                <body>
                  <nav class="navbar navbar-default">
                    <div class="container">
                      <div class="navbar-header">
                            {{range .Menu}}<a class="navbar-brand" href="/{{ . }}">{{ . }}</a>{{else}}<div><strong>Data Missing</strong></div>{{end}}
					  </div>
					  <div>
                    </div>
                  </nav>
                </body>
			  </html>`

	//Function to check for errors
	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	//Template Parser
	t, err := template.New("webpage").Parse(tpl)
	check(err)

	data := struct {
		Title string
		Menu  []string
		Body  []string
	}{
		Title: "Go http example",
		Menu: []string{
			"Home",
			"Mission",
			"About",
		},
	}

	err = t.Execute(w, data)
	check(err)

}

func main() {

	http.HandleFunc("/", handler)

	log.Printf("About to listen on port 80. Go to http://127.0.0.1:80/")
	err := http.ListenAndServe(":80", nil)
	log.Fatal(err)
}
