# go-enum
an enum generator for go. It generates a set of accompanying methods to a slice/array of given struct. 

*Bare in mind this is the first draft version 0.1, the API will change*

My experience with enums is from java, this implementation bear some resemblance to that. 
As go does not support the concept of `const` for structs or slice/array this implementation rely on wrapping the structs in "virtually" constant functions.
This means that the compiled "enum" will have the slightly different syntax `enumPkg.ENUMLITERAL()` In stead of `enumPkg.ENUMLITERAL` that we know from other languages and from regular go `const`.

## Installation 
```
go install github.com/MikkelHJuul/go-enum
```

## Usage 
Generate the functions using
`//go:generate go-enum` comment in the package where your enum is situated.
Your enum should be expressed as the package level Exported function `Values() T` where `T` is either slice or array of the enum's type (or wrapping interface)

