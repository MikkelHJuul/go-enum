# go-enum
an enum generator for go. It generates a set of accompanying methods to a slice/array of given struct. 

*Bare in mind this is the first draft version 0.1, the API will change*

My experience with enums is from java, this implementation bear some resemblance to that. 
As go does not support the concept of `const` for structs or slice/array this implementation rely on wrapping the structs in "virtually" constant `functions`.
This means that the compiled "enum" will have the slightly different syntax `enumPkg.ENUMLITERAL()` In stead of `enumPkg.ENUMLITERAL` that we know from other languages and from regular go `const`.

## Installation 
```
go install github.com/MikkelHJuul/go-enum
```

## Usage 
Generate the functions using
`//go:generate go-enum` comment in the package where your enum is situated.
Your enum should be expressed as the package level Exported function `Values() T` where `T` is either slice or array of the enum's type (or interface). `T` must implement [`enum.Named`](enum/enum.go)
- Each enum will have a method compiled (`$NAME() E` - where E is th instance type) that output exactly that enum (currently via array lookup with its index). 
- a method `EnumMap map[string]T` is built for name lookup
- a name-lookup method is compiled: `ValueOf(string) (T, bool)`
- a `const` `int` is compiled for each ordinal(index number) following the syntax `${NAME}Ordinal`

In order to assert full runtime certainty if the enum constants the `Values` method should fully contain the enumset as such:
```go
func Values() [...]MyEnum {
    return [size]MyEnum{
          {...},
          ...,
    }
}
```
Also you may wish to use unexported "guarded" types of you wish to remove the ability for downstream code to mess with the enums.

## future changes
- Remove the interface `enum.Named`

  The package should rely on a tag to figure which struct item should describe the enum's name. Build a small name-enum for modifiers: (camelCase,PascalCase,SCREAMING_SNAKE_CASE, etc.) - although allowing full control via the `Name` method is nice.  It requires runtime to work (whence the test-injection)
- build using ast

  And inline "enums" in each method
