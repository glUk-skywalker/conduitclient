package urlvalues

import (
	"net/url"
	"strings"
)

type URLValues url.Values

func (v URLValues) NoSortEncode() string {
	if v == nil {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	// sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		keyEscaped := url.QueryEscape(k)
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(keyEscaped)
			buf.WriteByte('=')
			buf.WriteString(url.QueryEscape(v))
		}
	}
	return buf.String()
}

func (v URLValues) Set(key string, value string) {
	url.Values(v).Set(key, value)
}

func (v URLValues) Add(key string, value string) {
	url.Values(v).Add(key, value)
}
