package paths

import (
	"creatures/creature"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func GetRandomCreature(w http.ResponseWriter, r *http.Request) {
	creature := creature.NewCreature()
	fmt.Printf("/random_creature request received\n")
	htmxFile, err := os.ReadFile("../static/components/creature.html")
	if err != nil {
		fmt.Printf("error reading file")
	}
	htmlTemplate := string(htmxFile)
	template := template.Must(template.New("hello").Parse(htmlTemplate))
	builder := &strings.Builder{}
	template.Execute(builder, creature)
	s := builder.String()
	io.WriteString(w, s)
}

func GetAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("/random_creature request received\n")
	htmlFile, err := os.ReadFile("../static/components/about.html")
	if err != nil {
		fmt.Printf("error reading file")
	}
	htmlTemplate := string(htmlFile)
	io.WriteString(w, htmlTemplate)
}
