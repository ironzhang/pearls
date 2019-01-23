package tpl

import (
	"os"
	"testing"
)

type Head struct {
	Title string
}

type Blog struct {
	Head    Head
	Toc     string
	Article string
}

func TestParseGlob(t *testing.T) {
	tp, err := ParseGlob("./layouts/include/*.html")
	if err != nil {
		t.Fatalf("parse glob: %v", err)
	}

	_, err = tp.ParseGlob("./layouts/html/*.html")
	if err != nil {
		t.Fatalf("parse glob: %v", err)
	}

	data := Blog{
		Head:    Head{Title: "Title"},
		Toc:     "Toc",
		Article: "Article",
	}
	err = tp.ExecuteTemplate(os.Stdout, "blog.html", data)
	if err != nil {
		t.Fatalf("execute template: %v", err)
	}
}
