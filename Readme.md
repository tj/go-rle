
# go-rle

Go Run-length encoding (currently just ints). RLE is great for data with low cardinality,
for example a log severity enum is an especially good use-case as most logs are INFO, many millions of points can be compressed to a few bytes.

## Badges

[![GoDoc](https://godoc.org/github.com/tj/go-rle?status.svg)](https://godoc.org/github.com/tj/go-rle)
![](https://img.shields.io/badge/license-MIT-blue.svg)
![](https://img.shields.io/badge/status-stable-green.svg)
[![](http://apex.sh/images/badge.svg)](https://apex.sh/)

---

> [tjholowaychuk.com](http://tjholowaychuk.com) &nbsp;&middot;&nbsp;
> GitHub [@tj](https://github.com/tj) &nbsp;&middot;&nbsp;
> Twitter [@tjholowaychuk](https://twitter.com/tjholowaychuk)
