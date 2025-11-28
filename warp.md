# Warp Development Guide - libxml2

A Go interface to libxml2 with DOM support.

## Quick Start

### Prerequisites
```bash
# Ensure libxml2 development files are installed
# Arch Linux:
sudo pacman -S libxml2

# Verify pkg-config can find libxml2
pkg-config --list-all | grep libxml
```

### Build
```bash
# Standard build
go build ./...

# Static build
go build -tags static_build
```

## Testing

### Run All Tests
```bash
go test ./...
```

### Run Tests with Verbose Output
```bash
go test -v ./...
```

### Run Benchmarks
```bash
go test -v -run=none -benchmem -benchtime=5s -bench .
```

### Run Specific Package Tests
```bash
go test ./parser
go test ./xpath
go test ./xsd
```

## Linting & Code Quality

### Run Linter
```bash
golangci-lint run
```

### Format Code
```bash
go fmt ./...
```

## Project Structure

- `libxml2` - Global utility functions (ParseString, etc.)
- `types` - Common data types (types.Node)
- `parser` - Parser routines
- `dom` - DOM manipulation of XML documents/nodes
- `xpath` - XPath tools
- `xsd` - XML Schema tools
- `clib` - C libxml2 wrapper (do not modify unless certain)

## Common Tasks

### Parse XML
```go
d, err := libxml2.ParseString(xmlstring)
if err != nil {
    return err
}
defer d.Free()
```

### Parse HTML
```go
doc, err := libxml2.ParseHTMLReader(reader)
if err != nil {
    return err
}
defer doc.Free()
```

### XPath Queries
```go
text := xpath.String(node.Find("//xpath/expression"))
```

### Create Documents
```go
d := dom.CreateDocument()
e, err := d.CreateElement("foo")
if err != nil {
    return err
}
d.SetDocumentElement(e)
```

## Troubleshooting

### Build Failures
- Ensure libxml2-dev/devel package is installed
- Check pkg-config can find libxml-2.0: `pkg-config --list-all | grep libxml`
- Verify C compiler can find headers in search path
- For non-standard installations, set PKG_CONFIG_PATH

### Static Linking
Use the `static_build` tag:
```bash
go build -tags static_build
```

## Related Projects
- [go-xmlsec](https://github.com/lestrrat-go/xmlsec) - XML Security library
