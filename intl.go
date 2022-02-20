package intl

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Intl struct {
	languages map[string]map[string]string
}

type LangSource struct {
	Lang string
	File string
}

func New(sources ...LangSource) (*Intl, error) {
	languages := make(map[string]map[string]string)
	for _, src := range sources {
		language, err := loadLanguage(src.File)
		if err != nil {
			return nil, errors.Wrapf(err, "could not load language (%s)", src.Lang)
		}
		languages[src.Lang] = language
	}

	return &Intl{
		languages: languages,
	}, nil
}

func (i *Intl) GetLocale(lang string) (*Locale, error) {
	messages, ok := i.languages[lang]
	if !ok {
		return nil, errors.New("locale not found in the bundle")
	}

	return &Locale{
		messages: messages,
	}, nil
}

func loadLanguage(filepath string) (map[string]string, error) {
	src, err := os.Open(filepath)
	if err != nil {
		return nil, errors.Wrap(err, "could not open language file")
	}

	messages := make(map[string]string)
	err = yaml.NewDecoder(src).Decode(messages)
	if err != nil {
		return nil, errors.Wrap(err, "could not load messages from file")
	}
	return messages, nil
}
