package intl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	englishLang     = "en"
	nonexistentLang = "xx"

	testFilePath = "./test/test_lang.yml"

	expectedTestMsgName = "greet"
	expectedTestMsgText = "Hello!"
)

var (
	langSrcs = []LangSource{
		{
			Lang:     englishLang,
			Filepath: testFilePath,
		},
	}
)

func TestGetLocale(t *testing.T) {
	i, err := New(langSrcs...)
	assert.NoError(t, err)

	t.Run("locale exists", func(t *testing.T) {
		l, err := i.GetLocale(englishLang)

		assert.NoError(t, err)
		assert.Equal(t, expectedTestMsgText, l.messages[expectedTestMsgName])
	})

	t.Run("locale does not exist", func(t *testing.T) {
		_, err := i.GetLocale(nonexistentLang)
		assert.Error(t, err)
	})
}
