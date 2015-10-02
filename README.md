# graphite-phpfpm

## What is this?

PHP-FPM status to graphite gateway.

## Building it

1. Install [go](http://golang.org/doc/install)

2. Install "graphite-golang" go get -u github.com/marpaia/graphite-golang

2. Install "go-logging" go get -u github.com/op/go-logging

4. Compile graphite-phpfpm

        git clone git://github.com/akashihi/graphite-phpfpm.git
        cd graphite-phpfpm
        go build .

## Running it

Generally:

    graphite-phpfpm -url https://localhost/status -metrics-host 192.168.1.1 -metrics-port 2003 -metrics-prefix test -period 60

All parameters could be omited. Run with --help to het parameters description

## License 

See LICENSE file.

Copyright 2015 Denis V Chapligin <akashihi@gmail.com>
