/*
 *  syncmap: Map class that has proper threading support
 *  Copyright (C) 2013  Joshua Chase <jcjoshuachase@gmail.com>
 *
 *  This program is free software; you can redistribute it and/or modify
 *  it under the terms of the GNU General Public License as published by
 *  the Free Software Foundation; either version 2 of the License, or
 *  (at your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  GNU General Public License for more details.
 *
 *  You should have received a copy of the GNU General Public License along
 *  with this program; if not, write to the Free Software Foundation, Inc.,
 *  51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
*/

package syncmap

import (
	"sync"
)

type Map struct {
	internal map[interface{}]interface{}
	mut * sync.RWMutex
}

func (sm Map) Lock() {
	sm.mut.Lock()
}

func (sm Map) Unlock() {
	sm.mut.Unlock()
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

// This is the only function that doesn't auto-lock the map.
// DO NOT USE IT WITHOUT Lock()'ing IT FIRST!
func (sm Map) Map() map[interface{}]interface{} {
	m := make(map[interface{}]interface{})
	for i,v := range sm.internal {
		m[i] = v
	}
	return m
}

func (sm Map) LockMap() map[interface{}]interface{} {
	sm.Lock()
	ret := sm.Map()
	sm.Unlock()
	return ret
}


func New() Map {
	var mut sync.RWMutex
	m := make(map[interface{}]interface{})

	return Map{m,&mut}
}
