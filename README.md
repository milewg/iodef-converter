# IODEF Converter

IODEF Converter is a tool for converting IODEF (Incident Object Description Exchange Format) data object to different formats. Currently, it supports converting XML <-> JSON format. This implmentation is based on the Internet drafts below.

* RFC7970
* draft-ietf-mile-jsoniodef-06(developing version)
developing version: https://github.com/milewg/draft-ietf-mile-jsoniodef/blob/master/draft-ietf-mile-jsoniodef-06.xml

## Setup

### Requirements
* [git](https://git-scm.com/)
* [go](https://golang.org/)
  * go 1.10.3 or later

### Recommandation Environment
* Windows 10
* Ubuntu 16.04

### Install Go

[Download and install go](https://golang.org/doc/install)

### Set GOPATH
[Setting GOPATH in your environment](https://github.com/golang/go/wiki/SettingGOPATH)
 
### Related libraries and patch files
    $ go get github.com/metaleap/go-util
    $ go get github.com/metaleap/go-buildrun
    $ go get github.com/sirupsen/logrus
    
Download customized go-xsd.tar.gz and go-xsd-pkg.tar.gz from https://github.com/TakeshiTakahashi/2018iodef-cbor/tree/master/iodef-converter and decompress into `$GOPATH/src/github.com/metaleap/`

    $ tar zxvf go-xsd.tar.gz -C $GOPATH/src/github.com/metaleap/
    $ tar zxvf go-xsd-pkg.tar.gz -C $GOPATH/src/github.com/metaleap/

Download patch files for go-xsd-pkg from https://github.com/TakeshiTakahashi/2018iodef-cbor/tree/master/iodef-converter/go-xsd-pkg
and overwrite onto `$GOPATH/src/github.com/metaleap/go-xsd-pkg`


### Converter
    $ mkdir $GOPATH/src/iodef-converter/iodef_converter
Download source code of converter from https://github.com/TakeshiTakahashi/2018iodef-cbor/tree/master/iodef-converter/iodef_converter to this folder

### Data
    $ mkdir $GOPATH/src/iodef-converter/infiles
    $ mkdir $GOPATH/src/iodef-converter/outfiles
Download sample input files from https://github.com/TakeshiTakahashi/2018iodef-cbor/tree/master/iodef-converter/infiles `$GOPATH/src/iodef-converter/infiles`

### Usage
Execution command

    $ cd $GOPATH/src/iodef-converter
    $ go run iodef_converter/main.go -file=xxx -option=yyy``

Parameters
- *-file*: the input file for conversion. The file is supposed to be in `$GOPATH/src/iodef-converter/infiles/` directory
- *-option*:
  - *xml*: convert input file into XML format
  - *json*: convert input file into JSON format

Examples

    $ go run iodef_converter/main.go -file=iodef-2.0.json -option=xml
    $ go run iodef_converter/main.go -file=iodef-2.0.xml -option=json

Output files are supposed to be in `$GOPATH/src/iodef-converter/outfiles/` directory
