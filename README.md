# intl

Minimalistic Go internationalization (i18n) package.

## Usage

Setting up an INTL instance

```go
langSrcs = []LangSource{
    {
        Lang:     "EN",
        Filepath: "path/to/file.yml",
    },
}

i, err := intl.New(langSrcs...)
```

Getting the specific locale before use

```go
locale, err := i.GetLocale("EN")
```

Using locale to get a basic message

```go
msg, err := locale.Msg("msg-name")
```

Using locale to get a template message

```go
data := MsgTmpl{"name": "John Doe"}
msg, err := locale.Msg("template-msg-name", data)
```

## Language files

Files are just a collection of messages.
Messages can optionally be templates, following the default go `text/template` logic.

```yaml
# example.yml

greet: "Welcome!"
dialog-start: "Hello, {{ .name }}!"
```
