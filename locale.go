package intl

import (
	"bytes"
	"text/template"

	"github.com/pkg/errors"
)

type Locale struct {
	messages map[string]string
}

type MsgTmpl map[string]interface{}

func (l *Locale) Msg(id string, tmpl ...MsgTmpl) (string, error) {
	msg, ok := l.messages[id]
	if !ok {
		return "", errors.New("message not found in the bundle")
	}

	if len(tmpl) == 0 {
		return msg, nil
	}

	t, err := template.New("msg").Parse(msg)
	if err != nil {
		return "", errors.Wrap(err, "could not parse message template")
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, tmpl[0])
	if err != nil {
		return "", errors.Wrap(err, "could not fill message template with values")
	}
	return buf.String(), nil
}
