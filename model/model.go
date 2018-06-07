package model

import (
	"fmt"
	"github.com/StevenZack/tools"
	"net/http"
	"net/url"
)

func ZHtoEN(s string) (string, error) {
	fd := url.Values{}
	fd.Set("q", s)
	fd.Set("from", "zh-CHS")
	fd.Set("to", "EN")
	// fd.Set("appKey", "ff889495-4b45-46d9-8f48-946554334f2a")
	appKey := "49f2b570d59a1d43"
	fd.Set("appKey", appKey)
	salt := tools.NewNumToken()
	fd.Set("salt", salt)
	fd.Set("sign", tools.MD5from(appKey+s+salt+"qTM0zXntDUWqFPtkWEnYrRNCWNeooA4U"))
	fmt.Println(tools.JsonObj(fd))
	rp, e := http.PostForm("http://openapi.youdao.com/api", fd)
	if e != nil {
		fmt.Println(e)
		return "", e
	}
	return tools.ReadAll(rp.Body), nil
}
func ENtoZH(s string) (string, error) {
	fd := url.Values{}
	fd.Set("q", s)
	fd.Set("to", "zh-CHS")
	fd.Set("from", "EN")
	// fd.Set("appKey", "ff889495-4b45-46d9-8f48-946554334f2a")
	appKey := "49f2b570d59a1d43"
	fd.Set("appKey", appKey)
	salt := tools.NewNumToken()
	fd.Set("salt", salt)
	fd.Set("sign", tools.MD5from(appKey+s+salt+"qTM0zXntDUWqFPtkWEnYrRNCWNeooA4U"))
	fmt.Println(tools.JsonObj(fd))
	rp, e := http.PostForm("http://openapi.youdao.com/api", fd)
	if e != nil {
		fmt.Println(e)
		return "", e
	}
	return tools.ReadAll(rp.Body), nil
}
func Auto2Auto(s string) (string, error) {
	fd := url.Values{}
	fd.Set("q", s)
	fd.Set("to", "auto")
	fd.Set("from", "auto")
	// fd.Set("appKey", "ff889495-4b45-46d9-8f48-946554334f2a")
	appKey := "49f2b570d59a1d43"
	fd.Set("appKey", appKey)
	salt := tools.NewNumToken()
	fd.Set("salt", salt)
	fd.Set("sign", tools.MD5from(appKey+s+salt+"qTM0zXntDUWqFPtkWEnYrRNCWNeooA4U"))
	fmt.Println(tools.JsonObj(fd))
	rp, e := http.PostForm("http://openapi.youdao.com/api", fd)
	if e != nil {
		fmt.Println(e)
		return "", e
	}
	return tools.ReadAll(rp.Body), nil
}
