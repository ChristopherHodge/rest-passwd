# getpw*()/getgr*() REST API

A service to provide a getpw*()/getgr*() style`REST interface to obtain user
and group information from POSIX compliant passwd/group files.

## Prerequisites

This is a Go application.  This service is based purely on the native Go stdlib.

Visit [golang.com](http://golang.com) for more information.

### Installation

Place this project in your `$GOPATH/src` and build with:

```
go build rest-passwd
```

### Configuration

By default, the application will look for a `config.json` in its own directory.  There is a
sample configuration file in the project repo.  The configuration file must contain all of
the settings to be consumable.  If the config file cannot be loaded, defaults are used.  Those
defaults are the values in the sample configuration JSON distributed with this project.

### Logging

By default, the only logger configured is for WARN level and displays to stderr.

## Development

### Testing

A full suite of tests are available.  These can be run with

```
go test -v rest-passwd/...
```

## Contact

Christopher Hodge - <[chris.l.hodge@gmail.com](mailto:chris.l.hodge@gmail.com)>

