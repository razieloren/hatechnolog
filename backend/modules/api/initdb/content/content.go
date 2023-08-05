package content

import (
	"bytes"
	"compress/zlib"
	"crypto/sha256"
	"fmt"
	"os"

	"github.com/adrg/frontmatter"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
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
	return markdown.ToHTML([]byte(md), p, renderer)
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
