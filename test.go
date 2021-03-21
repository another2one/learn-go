// 暂时测试
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

type CookieSameSite string
type CookiePriority string

// CookieSourceScheme represents the source scheme of the origin that
// originally set the cookie. A value of "Unset" allows protocol clients to
// emulate legacy cookie scope for the scheme. This is a temporary ability and
// it will be removed in the future.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#type-CookieSourceScheme
type CookieSourceScheme string

// String returns the CookieSourceScheme as string value.
func (t CookieSourceScheme) String() string {
	return string(t)
}

// CookieSourceScheme values.
const (
	CookieSourceSchemeUnset     CookieSourceScheme = "Unset"
	CookieSourceSchemeNonSecure CookieSourceScheme = "NonSecure"
	CookieSourceSchemeSecure    CookieSourceScheme = "Secure"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t CookieSourceScheme) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t CookieSourceScheme) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *CookieSourceScheme) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch CookieSourceScheme(in.String()) {
	case CookieSourceSchemeUnset:
		*t = CookieSourceSchemeUnset
	case CookieSourceSchemeNonSecure:
		*t = CookieSourceSchemeNonSecure
	case CookieSourceSchemeSecure:
		*t = CookieSourceSchemeSecure

	default:
		in.AddError(errors.New("unknown CookieSourceScheme value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
// func (t *CookieSourceScheme) UnmarshalJSON(buf []byte) error {
// 	return easyjson.Unmarshal(buf, t)
// }

// SetCookiesParams sets given cookies.
type SetCookiesParams struct {
	Cookies []CookieParam `json:"cookies"` // Cookies to be set.
}

type CookieParam struct {
	Name         string             `json:"name"`                   // Cookie name.
	Value        string             `json:"value"`                  // Cookie value.
	URL          string             `json:"url,omitempty"`          // The request-URI to associate with the setting of the cookie. This value can affect the default domain, path, source port, and source scheme values of the created cookie.
	Domain       string             `json:"domain,omitempty"`       // Cookie domain.
	Path         string             `json:"path,omitempty"`         // Cookie path.
	Secure       bool               `json:"secure,omitempty"`       // True if cookie is secure.
	HTTPOnly     bool               `json:"httpOnly,omitempty"`     // True if cookie is http-only.
	SameSite     CookieSameSite     `json:"sameSite,omitempty"`     // Cookie SameSite type.
	Priority     CookiePriority     `json:"priority,omitempty"`     // Cookie Priority.
	SameParty    bool               `json:"sameParty,omitempty"`    // True if cookie is SameParty.
	SourceScheme CookieSourceScheme `json:"sourceScheme,omitempty"` // Cookie source scheme type.
	SourcePort   int64              `json:"sourcePort,omitempty"`   // Cookie source port. Valid values are {-1, [1, 65535]}, -1 indicates an unspecified port. An unspecified port value allows protocol clients to emulate legacy cookie scope for the port. This is a temporary ability and it will be removed in the future.
}

func main() {

	// json 测试
	jsonString, err := ioutil.ReadFile("D:/go/learn/app/spider/chromedp/cookies.tmp")
	if err != nil {
		log.Println("cookie file empty: ", err)
		return
	}
	var test SetCookiesParams
	if err := json.Unmarshal([]byte(jsonString), &test); err != nil {
		fmt.Println("json unmarshal error:", err)
	}
	fmt.Printf("%v", test)
	return

	ch := 'b'
	fmt.Println(ch / 2.0)

	//s1 := []int{1, 2, 4}
	s1 := make(map[int]int, 6)
	s1[1] = 1
	printAddress(s1)
	for i, v := range s1 {
		printAddress(s1)
		s1[i+1] = v + 1
	}
	fmt.Println(s1)

	for _, v := range []int{1, 2, 4} {
		fmt.Printf("v address is %v \n", &v)
		go func(v int) {
			fmt.Println(v)
		}(v)
	}
	time.Sleep(time.Microsecond * 100)
}

func printAddress(v interface{}) {
	fmt.Printf("v address is %p \n", &v)
}
