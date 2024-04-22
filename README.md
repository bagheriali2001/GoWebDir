# GoWebDir

GoWebDir is a simple web server that serves the contents of a directory. It is written in Go without any external web server dependencies.

## Usage

``` bash
go-web-dir /path/to/directory

# Example

go-web-dir -p 8080 /path/to/directory
```

## Installation

Note: Currently, the only way to install GoWebDir is to build it from source (compiled sources will be added to release in future versions) Download the source code and run the following command:

``` bash
go build -o go-web-dir
```

and then move the binary to a directory in your PATH.

## License

GoWebDir is licensed under the MIT License. See the LICENSE file for more information.