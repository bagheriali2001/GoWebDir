package templates

import (
	"html/template"
	"io"
	"regexp"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"

	_ "embed"
)

var minifier *minify.M

func GetMinifiedWriter(contentType string, w io.Writer) io.WriteCloser {
	if minifier == nil {
		minifier = minify.New()
		minifier.Add("text/html", &html.Minifier{
			KeepDefaultAttrVals: true,
			KeepDocumentTags:    true,
			KeepEndTags:         true,
			KeepQuotes:          true,
		})
		minifier.AddFunc("text/css", css.Minify)
		minifier.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
	}
	return minifier.Writer(contentType, w)
}

var htmlTemplate *template.Template

//go:embed index.html
var rawHTMLTemplate string

func GetHTMLTemplate() *template.Template {
	if htmlTemplate == nil {
		htmlTemplate = template.Must(template.New("index.html").Parse(rawHTMLTemplate))
		// htmlTemplate = template.Must(template.ParseFS(content, "templates/index.html"))
	}
	return htmlTemplate
}
