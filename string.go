package goal

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
)

//StringIn check if array contains a string value.
func StringIn(s string, vals []string) bool {
	for _, v := range vals {
		if v == s {
			return true
		}
	}

	return false
}

//Prefixed ensure string has prefix, and not prepend if it already has.
func Prefixed(s, prefix string) string {
	if !strings.HasPrefix(s, prefix) {
		return prefix + s
	}
	return s
}

//Suffixed ensure string has suffer, and not append if it already has.
func Suffixed(s, suffix string) string {
	if !strings.HasSuffix(s, suffix) {
		return s + suffix
	}

	return s
}

//SuffixedNewline append '\n' to string if it does't have one
func SuffixedNewline(s string) string {
	return Suffixed(s, "\n")
}

//Split2 split string by `sep` to 2 sub string.
//if no split happen, second return value will be empty string.
func Split2(s, sep string) (s1, s2 string) {
	index := strings.Index(s, sep)
	if index < 0 {
		return s, ""
	} else {
		return s[:index], s[index+len(sep):]
	}
}

//RSplit2 reverse split string by `sep` to 2 sub string, see `Split2`.
func RSplit2(s, sep string) (s1, s2 string) {
	index := strings.LastIndex(s, sep)
	if index < 0 {
		return s, ""
	} else {
		return s[:index], s[index+len(sep):]
	}
}

//Atoi name after c libary function `atoi`, convert string to integer, return 0 if convert failed.
func Atoi(s string) int {
	n, _ := strconv.ParseInt(s, 10, 64)
	return int(n)
}

//Itoa name after c library function `itoa`, convert integer to string.
func Itoa(n int) string {
	return strconv.FormatInt(int64(n), 10)
}

//JSONEncode encode object to []byte, ignore error happens.
//
//Encode will disable html escape, similar with python default behavior.
func JSONEncode(v interface{}) []byte {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.Encode(v)
	return buffer.Bytes()
}

//JSONEncodeString encode object to string, ignore error happens.
//
//Encode will disable html escape, similar with python default behavior.
func JSONEncodeString(v interface{}) string {
	return string(JSONEncode(v))
}
