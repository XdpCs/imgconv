# imgconv

imgconv is a tool for converting images into images of other formats.

![GitHub watchers](https://img.shields.io/github/watchers/XdpCs/imgconv?style=social)
![GitHub stars](https://img.shields.io/github/stars/XdpCs/imgconv?style=social)
![GitHub forks](https://img.shields.io/github/forks/XdpCs/imgconv?style=social)
![GitHub last commit](https://img.shields.io/github/last-commit/XdpCs/imgconv?style=flat-square)
![GitHub repo size](https://img.shields.io/github/repo-size/XdpCs/imgconv?style=flat-square)
![GitHub license](https://img.shields.io/github/license/XdpCs/imgconv?style=flat-square)

## Install

`go get`

```shell
go get -u github.com/XdpCs/imgconv
```

`go mod`

```shell
require github.com/XdpCs/imgconv
```

## cli

```shell
go run cli.go -url https://avatars.githubusercontent.com/u/72180730 -format webp
```

If you want to convert an image from a URL, you can use the `-url` flag.

More information can be found by running `go run cli.go -h`.

## License

imgconv is under the [MIT](LICENSE). Please take a look at LICENSE for more information.