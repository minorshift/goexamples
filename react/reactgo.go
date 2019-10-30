package main

import (
	"html/template"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
	const header = `<!DOCTYPE html>
    <head>
        <meta charset="UTF=8">
        <title>{{.Title}}</title>
        <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap.min.css" integrity="sha384-1q8mTJOASx8j1Au+a5WDVnPi2lkFfwwEAa8hDDdjZlpLegxhjVME1fgjWPGmkzs7" crossorigin="anonymous">
        <script src="https://unpkg.com/react@16/umd/react.development.js"></script>
        <script src="https://unpkg.com/react-dom@16/umd/react-dom.development.js"></script>
        <script src="https://unpkg.com/babel-standalone@6.15.0/babel.min.js"></script>
        <script async src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/js/bootstrap.min.js" integrity="sha384-0mSbJDEHialfmuBBQP6A4Qrprq5OVfW37PRR3j5ELqxss1yVqOtnepnHVP9aJ7xS" crossorigin="anonymous"></script>
    </head>
    <body>`
	const menu = `<nav class="navbar navbar-default">
        <div class="container">
            <div class="navbar-header">
                {{range .Menu}}<a class="navbar-brand" href="/{{ . }}">{{ . }}</a>{{else}}<div><strong>Data Missing</strong></div>{{end}}
            </div>
        </div>
    </nav>`
	const content = `<div id="root"></div>
    <script type="text/babel">  
        class Hello extends React.Component {
            render() {
                return <h1>Hello world!</h1>;
            }
        }
        ReactDOM.render(
            <Hello />,
            document.getElementById("root")
        );
    </script>`
	const footer = `</body>
</html>`
	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	tpl := header + menu + content
	t, err := template.New("webpage").Parse(tpl)
	check(err)

	data := struct {
		Title   string
		Menu    []string
		Content []string
	}{
		Title: "Go react example",
		Menu: []string{
			"Home",
			"Mission",
			"About",
		},
		Content: []string{
			"Welcome",
		},
	}

	err = t.Execute(w, data)
	check(err)
}

func main() {
	http.HandleFunc("/", handler)

	log.Printf("About to listen on port 80. Go to http:127.0.0.1:80/")
	err := http.ListenAndServe(":80", nil)
	log.Fatal(err)
}
