package intl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMsg(t *testing.T) {
	messages := map[string]string{
		"valid-basic-msg":      `Hello!`,
		"valid-template-msg":   `Len of {{ .word }} is {{ .word | len }}`,
		"invalid-template-msg": `Len of }} .word {{ is }} .word | len {{`,
	}

	locale := Locale{
		messages: messages,
	}

	t.Run("basic message exists", func(t *testing.T) {
		msg, err := locale.Msg("valid-basic-msg")

		assert.NoError(t, err)
		assert.Equal(t, msg, "Hello!")
	})

	t.Run("template message exists", func(t *testing.T) {
		data := MsgTmpl{"word": "exquisite"}
		msg, err := locale.Msg("valid-template-msg", data)

		assert.NoError(t, err)
		assert.Equal(t, msg, "Len of exquisite is 9")
	})

	t.Run("message does not exist", func(t *testing.T) {
		msg, err := locale.Msg("nonexistent-msg")

		assert.Error(t, err)
		assert.Equal(t, "", msg)
	})

	t.Run("message template is invalid", func(t *testing.T) {
		data := MsgTmpl{}
		msg, err := locale.Msg("invalid-template-msg", data)

		assert.Error(t, err)
		assert.Equal(t, "", msg)
	})

	t.Run("template data is invalid", func(t *testing.T) {
		data := MsgTmpl{"word": 0}
		msg, err := locale.Msg("valid-template-msg", data)

		assert.Error(t, err)
		assert.Equal(t, "", msg)
	})
}
