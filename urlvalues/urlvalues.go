package urlvalues

import (
	"net/url"
	"sort"
	"strconv"
	"strings"
)

// URLValues is own implementation of url.Values (with proper sorting, required
// for Phabricator)
type URLValues url.Values

// Encode encodes parameters with proper indexing
func (v URLValues) Encode() string {
	if v == nil {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		keyEscaped := url.QueryEscape(k)
		for i, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			if len(vs) > 1 {
				keyEscaped = url.QueryEscape(k + "[" + strconv.Itoa(i) + "]")
			}
			buf.WriteString(keyEscaped)
			buf.WriteByte('=')
			buf.WriteString(url.QueryEscape(v))
		}
	}
	return buf.String()
}

// Set is just a wrapper for url.Values.Set
func (v URLValues) Set(key string, value string) {
	url.Values(v).Set(key, value)
}

// Add is just a wrapper for url.Values.Add
func (v URLValues) Add(key string, value string) {
	url.Values(v).Add(key, value)
}
