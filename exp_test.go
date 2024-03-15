package expronaut

import (
	"bytes"
	"html/template"
	"testing"
)

func TestExp(t *testing.T) {
	content := `{{ if exp "foo == 5" "foo" .foo }} OK {{ else }} NOT OK {{ end }}`

	templ, err := template.New("test").Funcs(map[string]any{
		"exp": Exp,
	}).Parse(content)
	if err != nil {
		t.Fatal(err)
	}

	data := map[string]any{
		"foo": 5,
	}

	var wr bytes.Buffer
	err = templ.ExecuteTemplate(&wr, "test", data)
	if err != nil {
		t.Fatal(err)
	}

	if wr.String() != " OK " {
		t.Fatalf("expected OK, got %s", wr.String())
	}
}
