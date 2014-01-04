package syncmap

import (
	"sync"
)

type Map struct {
	internal map[interface{}]interface{}
	mut * sync.RWMutex
}

func (sm Map) Get(key interface{}) (interface{},bool) {
	sm.mut.RLock()

	val,ok := sm.internal[key]

	sm.mut.RUnlock()

	return val,ok
}

func (sm Map) Set(key, val interface{}) {
	sm.mut.Lock()

	sm.internal[key] = val

	sm.mut.Unlock()

}

func (sm Map) Delete(key interface{}) {
	sm.mut.Lock()
	delete(sm.internal,key)
	sm.mut.Unlock()
}

func (sm Map) Map() map[interface{}]interface{} {
	m := make(map[interface{}]interface{})
	sm.mut.RLock()
	for i,v := range sm.internal {
		m[i] = v
	}
	sm.mut.RUnlock()
	return m
}


func New() Map {
	var mut sync.RWMutex
	m := make(map[interface{}]interface{})

	return Map{m,&mut}
}
