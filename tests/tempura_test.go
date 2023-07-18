package tests

import (
	"os"
	"strings"
	"testing"

	"github.com/TwoWaySix/tempura"
)

func TestTempest(t *testing.T) {
	engine := tempura.NewTemplateEngine("testdata")

	content, err := engine.Prepare("index.html")
	if err != nil {
		t.Errorf(err.Error())
	}
	contentStripped := stripDown(content)

	b, err := os.ReadFile("testdata/result.html")
	if err != nil {
		t.Errorf(err.Error())
	}

	want := stripDown(string(b))
	if contentStripped != want {
		t.Errorf("%s != %s", contentStripped, want)
	}
}

func stripDown(content string) string {
	stripped := strings.ReplaceAll(content, "\n", "")
	return strings.ReplaceAll(stripped, " ", "")
}
