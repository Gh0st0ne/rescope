package doerror

import (
	"fmt"
	color "github.com/gookit/color"
	"log"
)

// BadJSON fatal error when JSON fails to unmarshal
func BadJSON() {
	log.Fatalf("\n%s Failed to parse JSON.", color.FgRed.Text("[!]"))
}

// NoScope fatal error when no scope tables was found
func NoScope(url string) {
	log.Fatalf("\n%s Unable to grab scope from %s. Is it public?", color.FgRed.Text("[!]"), url)
}

// NoResponse fatal error when http response failed
func NoResponse(url string) {
	log.Fatalf("\n%s No HTTP response from %s. Please make sure the URI is correct", color.FgRed.Text("[!]"), url)
}

// BadStatusCode fatal error when bad status code
func BadStatusCode(url string, c int) {
	code := fmt.Sprintf("%d", c)
	log.Fatalf("\n%s %s responded with status code %s. Make sure it's correct", color.FgRed.Text("[!]"), color.FgYellow.Text(url), color.FgRed.Text(code))
}
