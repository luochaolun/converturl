# converturl

描述

第一个例子 使用converturl将相对URL转化成绝对URL

安装方法
```sh
go get github.com/luochaolun/converturl
```

使用例子：
```go
package main

import (
    "fmt"
    obj "github.com/luochaolun/converturl"
)  

func main() {
    s := obj.UrlConvert("http://www.baidu.com", "a/b.html")
    fmt.Println(s)
}
```