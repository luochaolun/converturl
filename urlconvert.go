package urlconvert

import (
    "net/url"
)

func UrlConvert(baseUrl string, curUrl string) string{
    u, err := url.Parse(curUrl)
    if err != nil {
        return ""
    }
    base, err := url.Parse(baseUrl)
    if err != nil {
        return ""
    }
    return base.ResolveReference(u).String()
}

func Add(a,b int) int{
	return a+b
}

func Jian(a,b int) int{
	return a-b
}