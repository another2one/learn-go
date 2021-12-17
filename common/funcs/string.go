package funcs

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/url"
	"sort"
	"strings"
	"time"
)

// md5V 同 php md5
func Md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Sign(arr map[string]string) string {
	// 排序
	keys := make([]string, len(arr))
	i := 0
	for key := range arr {
		keys[i] = key
		i++
	}
	sort.Strings(sort.StringSlice(keys))

	// 开始构建签名串
	signStr := ""
	for _, key := range keys {
		if key == "sign" {
			continue
		}
		signStr += key + "=" + arr[key] + "&"
	}
	signStr = strings.TrimRight(signStr, "&")

	// debug
	fmt.Printf("\n sign string: %+v \n", signStr)
	return Md5V(Md5V(signStr) + "3214e23sa12wsased34ewdxw1edwxs")
}

// RandomStr 指定长度的随机的传
func RandomStr(length int) string {
	rand.Seed(time.Now().Unix())
	baseStr := "abcdefghijklmnopqrstuvwxyz123456789"
	baseLen := len(baseStr)
	randBytes := make([]byte, length)
	for i := 0; i < length; i++ {
		key := rand.Intn(baseLen)
		randBytes[i] = baseStr[key]
	}
	return string(randBytes)
}

func BuildQuery(a map[string]string, isEncode bool) string {
	str := ""
	for k, v := range a {
		if isEncode {
			str += k + "=" + url.QueryEscape(v) + "&"
		} else {
			str += k + "=" + v + "&"
		}
	}
	return strings.TrimRight(str, "&")
}

func GetFormValue(a map[string]string) url.Values {
	u := make(url.Values, len(a))
	for k, v := range a {
		u[k] = []string{v}
	}
	return u
}
