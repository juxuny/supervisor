package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetString(key string, defaultValue ...string) string {
	r := os.Getenv(key)
	if r == "" && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return r
}

func GetStringList(key string, splitter string, defaultValue ...[]string) []string {
	s := GetString(key, "")
	if s == "" && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return strings.Split(s, splitter)
}

func GetInt64List(key string, splitter string, defaultValue ...[]int64) []int64 {
	sl := GetStringList(key, splitter)
	if len(sl) == 0 && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	var ret []int64
	for _, item := range sl {
		v, err := strconv.ParseInt(item, 10, 64)
		if err == nil {
			ret = append(ret, v)
		}
	}
	return ret
}

func GetIntList(key string, splitter string, defaultValue ...[]int) []int {
	sl := GetInt64List(key, splitter)
	if len(sl) == 0 && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	ret := make([]int, len(sl))
	for i := range sl {
		ret[i] = int(sl[i])
	}
	return ret
}

func GetInt64(key string, defaultValue ...int64) int64 {
	r := os.Getenv(key)
	if r == "" && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	ret, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return 0
	}
	return ret
}

func GetInt(key string, defaultValue ...int) int {
	if len(defaultValue) > 0 {
		return int(GetInt64(key, int64(defaultValue[0])))
	}
	return int(GetInt64(key, 0))
}

func GetBool(key string, defaultValue ...bool) bool {
	var s string
	if len(defaultValue) > 0 {
		s = GetString(key, fmt.Sprintf("%v", defaultValue[0]))
	}
	ret, err := strconv.ParseBool(s)
	if err != nil {
		return false
	}
	return ret
}
