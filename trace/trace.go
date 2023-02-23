package trace

import (
	"fmt"
	"github.com/juxuny/env"
	"github.com/petermattis/goid"
	"runtime/debug"
	"sync"
)

var (
	reqIdMap  = &sync.Map{}
	uidMap    = &sync.Map{}       // 按goId 记录uid
	waitGroup = &sync.WaitGroup{} // 系统退出的时候等所有协程结束之后才退出

	IsDebug = env.GetString("LOG_LEVEL") == "debug"
)

func init() {
	InitReqId()
}

func GetReqId() string {
	id := goid.Get()
	if v, ok := reqIdMap.Load(id); ok {
		return v.(string)
	}
	return ""
}

func GenChildReqId() string {
	reqId := GetReqId() + "." + genReqId(6)
	return reqId
}

func InitReqId(reqId ...string) {
	id := goid.Get()
	if len(reqId) == 0 {
		var newReqId string
		if IsDebug {
			newReqId = "dev" + genReqId(20)
		} else {
			newReqId = "prod" + genReqId(20)
		}
		reqIdMap.LoadOrStore(id, newReqId)
	} else {
		reqIdMap.Store(id, reqId[0])
	}
	return
}

func CleanReqId() {
	id := goid.Get()
	reqIdMap.Delete(id)
	uidMap.Delete(id)
}

func SetUid(uid int64) {
	id := goid.Get()
	uidMap.Store(id, uid)
}

func GetUid() (uid int64, found bool) {
	id := goid.Get()
	value, ok := uidMap.Load(id)
	if ok {
		v, ok := value.(int64)
		return v, ok
	}
	return 0, false
}

func GoRun(f func()) {
	id := goid.Get()
	reqId := GetReqId()
	uid, foundUid := GetUid()
	waitGroup.Add(1)
	go func(parentId int64, parentReqId string) {
		defer func() {
			waitGroup.Done()
		}()
		reqId := genReqId(6)
		newReqId := fmt.Sprintf("%s.%s", parentReqId, reqId)
		id := goid.Get()
		reqIdMap.Store(id, newReqId)
		if foundUid {
			uidMap.Store(id, uid)
		}
		defer func() {
			reqIdMap.Delete(id)
			uidMap.Delete(id)
		}()
		defer func() {
			if err := recover(); err != nil {
				debug.PrintStack()
			}
		}()
		f()
	}(id, reqId)
}

func Wait() {
	waitGroup.Wait()
}
