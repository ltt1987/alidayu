package alidayu

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"sort"
	"strings"
)

func getRequestBody(m map[string]string) (reader io.Reader, size int64) {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	v := url.Values{}

	signString := AppSecret
	for _, k := range keys {
		v.Set(k, m[k])
		signString += k + m[k]
	}
	signString += AppSecret

	signByte := md5.Sum([]byte(signString))
	sign := strings.ToUpper(fmt.Sprintf("%x", signByte))
	v.Set("sign", sign)

	return ioutil.NopCloser(strings.NewReader(v.Encode())), int64(len(v.Encode()))
}
