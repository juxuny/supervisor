package trace

import (
	"math/rand"
	"sync"
	"time"
)

var (
	randPool = &sync.Pool{New: func() interface{} {
		return rand.New(rand.NewSource(time.Now().UnixNano()))
	}}
)

func genReqId(lens ...interface{}) string {
	r := randPool.Get().(*rand.Rand)
	defer randPool.Put(r)
	r.Seed(time.Now().UnixNano())
	l := 8
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	if len(lens) > 0 {
		l = lens[0].(int)
	}
	if len(lens) > 1 {
		str = lens[1].(string)
	}

	bytes := []byte(str)
	var result []byte
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
