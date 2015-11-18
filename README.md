# protogen
Simple Go (golang) tool to generate gRPC-Go interface

## Install
`go get -u github.com/jesselucas/protogen`

## Usage
Must have [protocol buffer compiler](https://github.com/google/protobuf) installed. For OS X, I recommend using [Homebrew](http://brew.sh/):

`brew install protobuf`

Then run

`protogen /path/to/myService.proto`
