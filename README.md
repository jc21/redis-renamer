# Redis Renamer

Redis Renamer will simply rename the matching keys with a given prefix.

## Usage

```
Usage: redis-renamer --host HOST [--port PORT] [--db DB] [--user USER]
  [--pass PASS] [--filter FILTER] --prefix PREFIX [--unsafe] [--verbose]

Options:
  --host HOST            redis server hostname
  --port PORT            redis server port [default: 6379]
  --db DB                redis server db index [default: 0]
  --user USER            redis server auth username [default: jamiec]
  --pass PASS            redis server auth password
  --filter FILTER        keys filter string [default: *]
  --prefix PREFIX        key prefix to prepend
  --unsafe, -u           Don't bother checking if key already has prefix first
  --verbose, -v          Print a lot more info
  --help, -h             display this help and exit
  --version              display version and exit
```

## Install

#### Centos

RPM hosted on [yum.jc21.com](https://yum.jc21.com)

#### Go Get

```bash
go get github.com/jc21/redis-renamer
```


#### Building

```bash
git clone https://github.com/jc21/redis-renamer && cd redis-renamer
go build -ldflags="-X main.version=1.0.0" -o bin/redis-renamer cmd/redis-renamer/main.go
./bin/redis-renamer -h
```
