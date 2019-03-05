vlto
====

vlto shows velocity of your projects of [Toggl](https://toggl.com)

## Description

If you use Toggl to track working hours of some projects, vlto can show you velocity of them.

This tool calls Toggl API and gets total and iterative achieved hours.
Then it indicates when each project will be finished if you keep
the pace of this iteration.

## Requirements

* Go

## Usage

```
$ vlto --help
vlto shows velocity of your projects of Toggl

Usage:
  vlto [flags]

Flags:
      --config string   config file (default is $HOME/.config/vlto.toml)
      --format string   the output format 'table' (default) or 'json'
  -h, --help            help for vlto
      --version         version for vlto
```

## Install

```
$ go get -u github.com/it-akumi/vlto
```

## Author

[Takumi Ishii](https://github.com/it-akumi)

## License

[MIT](https://github.com/it-akumi/vlto/blob/master/LICENSE)
