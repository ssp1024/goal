package goal

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
)

func StringIn(s string, vals []string) bool {
	for _, v := range vals {
		if v == s {
			return true
		}
	}

	return false
}

func Prefixed(s, prefix string) string {
	if !strings.HasPrefix(s, prefix) {
		return prefix + s
	}
	return s
}

func Suffixed(s, suffix string) string {
	if !strings.HasSuffix(s, suffix) {
		return s + suffix
	}

	return s
}

func AppendNewline(s string) string {
	return Suffixed(s, "\n")
}

func Split2(s, sep string) (s1, s2 string) {
	index := strings.Index(s, sep)
	if index < 0 {
		return s, ""
	} else {
		return s[:index], s[index+len(sep):]
	}
}

func RSplit2(s, sep string) (s1, s2 string) {
	index := strings.LastIndex(s, sep)
	if index < 0 {
		return s, ""
	} else {
		return s[:index], s[index+len(sep):]
	}
}

func Atoi(s string) int {
	n, _ := strconv.ParseInt(s, 10, 64)
	return int(n)
}

func Itoa(n int) string {
	return strconv.FormatInt(int64(n), 10)
}

func JSONEncode(v interface{}) []byte {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.Encode(v)
	return buffer.Bytes()
}

func JSONEncodeString(v interface{}) string {
	return string(JSONEncode(v))
}
