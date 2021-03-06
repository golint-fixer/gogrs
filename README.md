gogrs [![Release](https://img.shields.io/github/release/toomore/gogrs.svg)](https://github.com/toomore/gogrs/releases)
======

[![GoDoc](https://godoc.org/github.com/toomore/gogrs?status.svg)](https://godoc.org/github.com/toomore/gogrs) [![Coverage Status](https://coveralls.io/repos/toomore/gogrs/badge.svg?branch=master)](https://coveralls.io/r/toomore/gogrs?branch=master) [![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/toomore/gogrs/master/LICENSE) [![Build Status](https://travis-ci.org/toomore/gogrs.svg?branch=master)](https://travis-ci.org/toomore/gogrs) [![Go Report Card](https://goreportcard.com/badge/github.com/toomore/gogrs)](https://goreportcard.com/report/github.com/toomore/gogrs)

gogrs is a tool for fetching data from Taiwan Stock Exchange(TWSE) and dockerizing. gogrs now is still in development. I will try my best to speed up to completed the same function with [grs](https://github.com/toomore/grs) (Python). gogrs 是擷取台灣上市股票股價資訊工具，目前還在大量的開發中。原始工具是用 [grs](https://github.com/toomore/grs)（Python 套件），目標是將基本功能用 go 來實作。

Packages
---------

1. realtime - [擷取盤中個股、指數即時股價資訊](https://godoc.org/github.com/toomore/gogrs/realtime) **（使用太過頻繁會有被擋掉的風險，目前無完善的解決辦法）**
2. twse - [擷取台灣股市上市/上櫃股票資訊、上市/上櫃類股清單、外資及陸資持股比率前二十名彙總表、 三大法人買賣金額統計表、三大法人買賣超日報、自營商、投信、外資及陸資買賣超彙總表](https://godoc.org/github.com/toomore/gogrs/twse)
3. tradingdays - [股市開休市判斷（支援非國定假日：颱風假）與當日區間判斷（盤中、盤後、盤後盤）](https://godoc.org/github.com/toomore/gogrs/tradingdays)
4. utils - [套件所需的公用工具（總和、平均、序列差、持續天數、民國日期解析、簡單亂數、標準差、簡單 net/http 快取）](https://godoc.org/github.com/toomore/gogrs/utils)

Cmd
----

![gogrs cmd demo](https://s3-ap-northeast-1.amazonaws.com/toomore/gogrs/gogrs_cmd_demo_20150615.png "gogrs cmd demo")

1. example - [簡單範例測試](cmd/docs/gogrs_example.md)
2. realtime - [擷取盤中即時資訊與大盤、上櫃、寶島指數](cmd/docs/gogrs_realtime.md)
3. report - [每日收盤後產生符合選股條件的報告](cmd/docs/gogrs_report.md)
4. cache - [清除 twse cache](cmd/docs/gogrs_cache.md)
5. server - [提供簡單的日期查詢 API Server](cmd/docs/gogrs_server.md)

相關的操作請參考 `-h` 的說明，或 [cmd/docs](cmd/docs/gogrs.md)

Docker
-------

Download image.

    docker pull toomore/gogrs

`tag:latest` bind to `branch:master`, more docker [info](https://registry.hub.docker.com/u/toomore/gogrs/).

Or ...

Build minify gogrs docker(`toomore/gogrs:mini`)

    cd ./build-mini.sh

Run `tradingdays server`.

    docker run -d -p 80:59123 toomore/gogrs:mini gogrs server

Or login run other cmd

    docker run -it toomore/gogrs:mini

Create a ramdisk volume

    docker create -v /run/shm/:/run/shm --name ramdisk toomore/gogrs:mini

Run with ramdisk volume

    docker run -it --volumes-from ramdisk toomore/gogrs:mini

TODO
-----

1. In English comment.
2. 盤中預估量能。
3. 個股對應分類股資訊。
4. 個股融資融券資訊。
5. 顯示三大法人。
6. 各類股盤後買賣超表。
7. 新聞內容擷取。

Known Issue
------------

1. None.

## Changelog

### v1.4 (2017/07/08)
* cmd using spf13/cobra

### v1.3 (2017/07/07)
* Fixed twse change source link. #4

### v1.2

#### v1.2.0 (2017/04/25)
* Fixed `realtime` #1 #3

#### v1.2.1 (2017/04/26)
* Pass realtime testing
* Put gogrs:mini into test

### v1.1

#### v1.1.0 (2017/04/07)
* Fixed `GetOSRamdiskPath`
* Add `Weight`, `WeightVolume`
* Add more test case

#### v1.1.1 (2017/04/08)
* Fixed Dockerfile

#### v1.1.2 (2017/04/08)
* Docker image using `alpine`
* Fixed image:`gogrs:mini` (`sh ./docker-mini/make.sh`)

#### v1.1.3 (2017/04/10)
* Fixed Dockerfile
* Update open time 2011 ~ 2017

### v1.0.0 (2016/10/08)
* Stable release for dependency tools.

### v0.1.0 (2016/10/08)
* gogrs first package version.

License
--------

The MIT License (MIT)

Copyright © 2015, 2016, 2017 Toomore Chiang, https://toomore.net/ <toomore0929@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

http://toomore.mit-license.org/
