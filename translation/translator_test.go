package translation_test

import (
	"github.com/oreillyross/hello-go/translation"
	"testing"
)

func TestTranslate(t *testing.T) {
	// Arrange
	tt := []struct {
		Word        string
		Language    string
		Translation string
	}{
		{
			Word:        "hello ",
			Language:    "english",
			Translation: "hello",
		}, // Notice the comma here after the closing brace
		{
			Word:        "bye",
			Language:    "German",
			Translation: "",
		},
		{
			Word:        "Hello",
			Language:    "german",
			Translation: "hallo",
		}, // And here
		{
			Word:        "hello",
			Language:    "finnish",
			Translation: "hei",
		}, // And also here
		{
			Word:        "hello",
			Language:    "dutch",
			Translation: "",
		}, // Even the last struct before closing brace should have a comma
	}
	// Arrange

	// Act
	for _, test := range tt {
		res := translation.Translate(test.Word, test.Language)
		// Assert
		if res != test.Translation {
			t.Errorf(`%v should be %v in %v, but got %v`, test.Word, test.Translation, test.Language, res)
		}
	}
}
