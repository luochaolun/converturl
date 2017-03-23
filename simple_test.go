package urlconvert

import (
	"fmt"
	"testing"
)

func Test_Zh(t *testing.T) {
	s := UrlConvert("http://www.baidu.com/", "a/b.html")
	t.Log(s)
	fmt.Println(s)
}

func Test_Add(t *testing.T) {
	i := Add(1, 2)
	t.Log(i)
	fmt.Printf("%d\n", i)
}