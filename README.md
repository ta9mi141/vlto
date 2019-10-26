vlto
====

vlto shows velocity of your projects of [Toggl](https://toggl.com)

## Description

If you use Toggl to track working hours of some projects, vlto can show you velocity of them.

This tool calls Toggl API and gets total and iterative achieved hours.
Then it indicates when each project will be finished if you keep
the pace of the iteration.

## Demo

![vlto.gif](vlto.gif)

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
      --proxy string    the URL of an environment variable HTTPS_PROXY
      --version         version for vlto
```

## Install

```
$ go get -u github.com/it-akumi/vlto
```

## Configuration

vlto requires a configuration written in [TOML](https://github.com/toml-lang/toml).

This tool uses `$HOME/.config/vlto.toml` in defalut but there is `--config` option
so that you can put your config anywhere you want and specify the path.

Configuration is composed of following properties.

| Property Name | Type            | Description                                                   |
| ------------- | --------------- | ------------------------------------------------------------- |
| ApiToken      | String          | Your api token can be found in https://toggl.com/app/profile. |
| WorkSpaceId   | String          | A workspace which has time entries of your projects.          |
| Projects      | Array of Tables | Settings of each project.                                     |

And following properties consist settings of each project.

| Property Name | Type            | Description                                     |
| ------------- | --------------- | ----------------------------------------------- |
| Name          | String          | Project name defined in Toggl.                  |
| TargetHour    | Integer         | Working hours you try to achieve.               |
| StartDate     | Local Date-Time | A date when aggregation of time entries starts. |
| IterationDays | Integer         | Number of days in 1 iteration.                  |

Example:

```
ApiToken = "0123456789abcdefghijklmnopqrstuv"
WorkSpaceId = "1234567"

[[Projects]]
Name = "Sample Project 1"
TargetHour = 1000
StartDate = 2016-10-11T00:00:00+00:00
IterationDays = 7

[[Projects]]
Name = "Sample Project 2"
TargetHour = 2000
StartDate = 2019-01-01T00:00:00+00:00
IterationDays = 14
```

## Author

[Takumi Ishii](https://github.com/it-akumi)

## License

[MIT](https://github.com/it-akumi/vlto/blob/master/LICENSE)
