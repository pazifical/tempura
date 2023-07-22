package tempura

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type TemplateEngine struct {
	templateDirectory string
}

func NewTemplateEngine(templateDirectory string) TemplateEngine {
	return TemplateEngine{
		templateDirectory: templateDirectory,
	}
}

func (te *TemplateEngine) Prepare(filePath string) (string, error) {
	b, err := os.ReadFile(filepath.Join(te.templateDirectory, filePath))
	if err != nil {
		return "", err
	}

	content := string(b)
	r, err := regexp.Compile("{{T.*T}}")
	if err != nil {
		return "", err
	}

	matches := r.FindAllString(content, -1)
	if matches == nil {
		return content, nil
	}

	for _, match := range matches {
		p := replaceAll(match, []string{"{{T", "T}}", " "}, "")

		sub, err := te.Prepare(p)
		if err != nil {
			log.Printf("ERROR: %v", err)
			continue
		}

		content = strings.ReplaceAll(content, match, sub)
	}

	return content, nil
}

func replaceAll(text string, inputs []string, output string) string {
	p := text
	for _, input := range inputs {
		p = strings.ReplaceAll(p, input, output)
	}
	return p
}
