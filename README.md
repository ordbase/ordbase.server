# Ordinal Base Server

ordbase - zero-config web server / services; (auto-)downloads & serves pre-configured pixel art collections "out-of-the-box"; incl. 2x/4x/8x zoom for bigger image sizes and more; binaries for easy "xcopy" installation for windows, linux & friends




## Download Binaries For Easy "Xcopy" Installation

Find the archives to download  - about 3 Megabytes (MB) - for Windows, Linux and Friends at the [**Releases Page »**](https://github.com/pixelartexchange/ordbase.server/releases).

Unpack the archive (e.g. `ordbase-*.tar.gz` or `ordbase-*.zip`) and than start / run the binary:

```
$ ordbase
```

This will start-up a (web) server (listening on port 8080). To test open up `http://localhost:8080` in your browser (to get the index web page listing all collections).


## Build & Run From Source


Use / issue / type  (in the `/ordbase.server` directory):

```
$ go build ordbase.go
```

to get a zero-config x-copy binary for your operation system / architecture.
To run use:

```
$ ordbase
```

This will start-up a (web) server (listening on port 8080). To test open up `http://localhost:8080` in your browser (to get the index web page listing all collections).











## Questions? Comments?

Post them over at the [Help & Support](https://github.com/geraldb/help) page. Thanks.