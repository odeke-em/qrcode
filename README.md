# qrcode

Create .png QR codes from urls by arguments either from stdin or by commandline

arguments.

## Installation

```shell
$ go get github.com/odeke-em/qrcode

$ qrcode [urls...]
```

## Example

* By arguments on the CLI

```shell
$ qrcode https://github.com/odeke-em/qrcode https://twitter.com/odeke_et
```

* By arguments from stdin

```shell
$ cat urls | qrcode
$ echo https://github.com/odeke-em/qrcode https://twitter.com/odeke_et | qrcode
```
