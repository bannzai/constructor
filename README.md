# constructor
`constructor` is generated [Constructor](https://golang.org/doc/effective_go.html#composite_literals) for struct.

## Install
Getting `constructor` via `go get`.

```shell
$ go get -u github.com/bannzai/constructor
```

## Usage
First, You can prepare configuration files when exec `constructor setup`. After created `constructor.yaml` and `constructor.tpl`.
`constructor` necessary these configuration files.

```shell
$ constructor setup
```

Next, You edit `constructor.yaml` for your project.  (e.g

```yaml
definitions:
- package: "structure"
  sourcePath: "structure/*.go"
  ignoredPaths: ["structure/argument.go"]
  templatePaths: ["constructor.tpl"]
  destinationPath: "structure/constructors.go"
- package: "generator"
  sourcePath: "generator/*.go"
  ignoredPaths: []
  templatePaths: ["constructor.tpl"]
  destinationPath: "generator/constructors.go"
```

Last, Execute `constructor generate`, after you edited configuration files.   
And you confirm diff between before and after executing it. You found  about constructor functions `.go` file.

```shell
$ constructor generate
```

<details>

<summary> $ cat structure/constructors.go </summary>

```go
// DO NOT EDIT THIS FILE.
// File generated by constructor.
// https://github.com/bannzai/constructor
package structure

// NewCodeStructure insitanciate Code
func NewCodeStructure(
        filePath Path,
        structs []Struct,
) Code {
        return Code{
                FilePath: filePath,
                Structs:  structs,
        }
}

// NewFieldStructure insitanciate Field
func NewFieldStructure(
        name string,
        _type string,
) Field {
        return Field{
                Name: name,
                Type: _type,
        }
}

// NewStructStructure insitanciate Struct
func NewStructStructure(
        fields []Field,
        name string,
) Struct {
        return Struct{
                Fields: fields,
                Name:   name,
        }
}

// NewDefinitionStructure insitanciate Definition
func NewDefinitionStructure(
        destinationPath Path,
        ignoredPaths []Path,
        _package string,
        sourcePath Path,
        templateFilePaths []Path,
) Definition {
        return Definition{
                DestinationPath:   destinationPath,
                IgnoredPaths:      ignoredPaths,
                Package:           _package,
                SourcePath:        sourcePath,
                TemplateFilePaths: templateFilePaths,
        }
}

// NewYamlStructure insitanciate Yaml
func NewYamlStructure(
        definitions []Definition,
) Yaml {
        return Yaml{
                Definitions: definitions,
        }
}
```

</details>

## LICENSE
**constructor** is available under the MIT license. See the LICENSE file for more info.
