package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/base64"
    "io/ioutil"
)

var latest = command{cmd: "ls -l"}

type command struct {
    cmd string
}

func latestCommandHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Latest Command: %s\n", latest.cmd)
}

func latestCommandHTMLHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<html><head><script>\n" +
    "console.log('sup');\n"+
    "setTimeout(function(){ window.location.reload(true); }, 1000);\n"+
    "</script></head><body>\n"+
    "<div id=\"command\">%s</div>\n"+
    "</body></html>",
    latest.cmd)
}



func setCommandHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
    _body, _ := ioutil.ReadAll(r.Body)
    body := string(_body)
    decoded, err := base64.StdEncoding.DecodeString(body)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
    latest = command{cmd: string(decoded)}
}

func main() {
    http.HandleFunc("/latestCommand", latestCommandHandler)
    http.HandleFunc("/latestCommand.html", latestCommandHTMLHandler)
    http.HandleFunc("/setCommand", setCommandHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
