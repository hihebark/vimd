<p align="center">
  <a href="https://godoc.org/github.com/hihebark/pickle">
    <img src="https://godoc.org/github.com/hihebark/pickle?status.svg" alt="GoDoc">
  </a>
  <a href="https://goreportcard.com/report/github.com/hihebark/pickle">
    <img src="https://goreportcard.com/badge/github.com/hihebark/pickle" alt="GoReportCard">
  </a>
  <a href="https://github.com/hihebark/pickle/blob/master/LICENSE">
    <img src="https://img.shields.io/aur/license/yaourt.svg" alt="license">
  </a>
  <a href="https://travis-ci.org/hihebark/pickle">
    <img src="https://travis-ci.org/hihebark/pickle.svg?branch=master">
  </a>
</p>

<p align="center"><img src="iampickle.jpg"></p>

# pickle

Pickle is a tool to preview markdown syntax alike github, It use github api to generate an html file.

## Usage

I set an enviroment variable `PICKLETOKEN="abcd1234********************************"`
for a specifique file:

`$ ./pickle -file test.md -token $PICKLETOKEN`

for working directory:

`$ ./pickle -token $PICKLETOKEN`

And then open your browser on `localhost:7069` OR `[::1]:7069`

## Install

you will need `> go 1.9`

get the project `go get -u github.com/hihebark/pickle`

and then install `go install` or `go build`

## Screenshot

![screenshot](picklescreenshot.png "screenshot")
