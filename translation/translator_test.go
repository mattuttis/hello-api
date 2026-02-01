package translation_test

import (
	"testing"

	"github.com/mattuttis/hello-api/translation"
)

func TestTranslate(t *testing.T) {
	// Arrange
	tt := []struct {
		Word        string
		Language    string
		Translation string
	}{
		{
			Word:        "hello",
			Language:    "english",
			Translation: "hello",
		},
		{
			Word:        "Hello",
			Language:    "german",
			Translation: "hallo",
		},
		{
			Word:        "hello",
			Language:    "german",
			Translation: "hallo",
		},
		{
			Word:        "hello",
			Language:    "finnish",
			Translation: "hei",
		},
		{
			Word:        "hello",
			Language:    "dutch",
			Translation: "",
		},
		{
			Word:        "bye",
			Language:    "dutch",
			Translation: "",
		},
		{
			Word:        "bye",
			Language:    "german",
			Translation: "",
		},
	}

	for _, test := range tt {
		// Act
		res := translation.Translate(test.Word, test.Language)

		// Assert
		if res != test.Translation {
			t.Errorf(
				`expected "%s" to be "%s" but received "%s"`,
				test.Word, test.Translation, res)
		}
	}
}
