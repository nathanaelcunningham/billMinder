package assets

import (
	"embed"
	"io/fs"
	"log"
)

type Embed struct {
	Templates fs.FS
	Emails    fs.FS
	Static    fs.FS
}

//go:embed "templates" "emails" "static/*"
var embeddedFiles embed.FS

func New() *Embed {
	emb := Embed{}
	var err error

	emb.Templates, err = fs.Sub(embeddedFiles, "templates")
	if err != nil {
		log.Fatal(err)
	}

	emb.Emails, err = fs.Sub(embeddedFiles, "emails")
	if err != nil {
		log.Fatal(err)
	}
	emb.Static, err = fs.Sub(embeddedFiles, "static")
	if err != nil {
		log.Fatal(err)
	}

	return &emb
}
