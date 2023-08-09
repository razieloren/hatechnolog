package content

import (
	"bytes"
	"compress/zlib"
	"crypto/sha256"
	"fmt"
	"os"
	"regexp"

	"github.com/adrg/frontmatter"
	chroma_html "github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

var (
	CodeRe = regexp.MustCompile("<pre><code class=\"language-(?P<Lang>.+)\">(?P<Source>(.|\n|\t)+)</code></pre>")
)

const (
	CodeStyle = "github-dark"
)

type ContentMetadata struct {
	Slug      string `yaml:"slug"`
	Title     string `yaml:"title"`
	Author    string `yaml:"author"`
	Ltr       bool   `yaml:"ltr"`
	Monetized bool   `yaml:"monetized"`
	Category  string `yaml:"category"`
	Type      string `yaml:"-"`
	Hash      []byte `yaml:"-"`
}

func (metadata *ContentMetadata) CalcHash(content []byte) {
	hash := sha256.New()
	hash.Write([]byte(metadata.Slug))
	hash.Write([]byte(metadata.Title))
	hash.Write([]byte(metadata.Author))
	if metadata.Monetized {
		hash.Write([]byte{1})
	} else {
		hash.Write([]byte{0})
	}
	if metadata.Ltr {
		hash.Write([]byte{1})
	} else {
		hash.Write([]byte{0})
	}
	hash.Write([]byte(metadata.Category))
	hash.Write([]byte(metadata.Type))
	hash.Write(content)
	metadata.Hash = hash.Sum(nil)
}

func markdownToHtml(md []byte) []byte {
	p := parser.NewWithExtensions(parser.CommonExtensions)
	opts := html.RendererOptions{Flags: html.CommonFlags}
	renderer := html.NewRenderer(opts)
	htmlContent := markdown.ToHTML([]byte(md), p, renderer)
	for _, match := range CodeRe.FindAllSubmatch(htmlContent, -1) {
		orig := match[0]
		language := match[1]
		sourceCode := match[2]
		lexer := lexers.Get(string(language))
		if lexer == nil {
			lexer = lexers.Fallback
		}
		style := styles.Get(CodeStyle)
		if style == nil {
			style = styles.Fallback
		}
		iterator, err := lexer.Tokenise(nil, string(sourceCode))
		if err != nil {
			fmt.Println("error tokenzing source code: ", err)
			continue
		}
		formatter := chroma_html.New(
			chroma_html.TabWidth(4), chroma_html.WithLineNumbers(true),
			chroma_html.WithLineNumbers(true), chroma_html.WithClasses(true),
			chroma_html.ClassPrefix("chr-"))

		var htmlBuffer bytes.Buffer
		if err := formatter.Format(&htmlBuffer, style, iterator); err != nil {
			fmt.Println("error formatting source code: ", err)
			continue
		}
		formattedCode := htmlBuffer.Bytes()
		formattedCode = bytes.ReplaceAll(formattedCode, []byte("&amp;"), []byte("&"))
		htmlContent = bytes.ReplaceAll(htmlContent, orig, formattedCode)
	}
	// Code sections are always left-to-right.
	htmlContent = bytes.ReplaceAll(htmlContent, []byte("<code"), []byte("<code dir=\"ltr\""))
	return htmlContent
}

func LoadContent(contentPath string, contentType string) (*ContentMetadata, []byte, error) {
	fd, err := os.Open(contentPath)
	if err != nil {
		return nil, nil, fmt.Errorf("open: %w", err)
	}
	var metadata ContentMetadata
	md, err := frontmatter.MustParse(fd, &metadata)
	html := markdownToHtml(md)
	metadata.Type = contentType
	metadata.CalcHash(html)

	var compressedHtml bytes.Buffer
	zlibWriter := zlib.NewWriter(&compressedHtml)
	if _, err := zlibWriter.Write(html); err != nil {
		return nil, nil, fmt.Errorf("zlib write: %w", err)
	}
	if err := zlibWriter.Close(); err != nil {
		return nil, nil, fmt.Errorf("zlib close: %w", err)
	}
	return &metadata, compressedHtml.Bytes(), nil
}
