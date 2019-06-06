package algorithms

import (
	"regexp"

	"github.com/Jeffail/gabs"
)

// TextStripBackSlash algorithm
func TextStripBackSlash(doc *gabs.Container) (*gabs.Container, bool) {
	text, ok := doc.Path("text").Data().(string)
	if text == "" || !ok {
		return nil, false
	}
	var re = regexp.MustCompile(`\\`)
	doc.Set(re.ReplaceAllString(text, ``), "text")
	return doc, true
}
