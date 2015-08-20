# goisvdo

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/shamsher31/goisvdo)

Check if filepath is a video file

### How to install
```go
go get github.com/shamsher31/goisvdo
```

### How to use
```go
package main

import (
	"fmt"
	"github.com/shamsher31/goisvdo"
)

func main() {
   fmt.Println(video.Is("./golang.mp4"))
   // true
}
```
###Aliasing Imports
If you already have package name ```video``` you can use following.
```go
package main

import (
	"fmt"
	videoPath "github.com/shamsher31/goisvdo"
)

func main() {
   fmt.Println(videoPath.Is("./golang.mp4"))
   // true
}
```

### Related
[govdoext](https://github.com/shamsher31/govdoext)<br>
[goistext](https://github.com/ferhatelmas/goistext)<br>
[goisimg](https://github.com/ferhatelmas/goisimg)<br>

### Why
This package is inspired by [is-video](https://www.npmjs.com/package/is-video) npm module.

### License
MIT © [Shamsher Ansari](https://github.com/shamsher31)
