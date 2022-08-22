/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package utils

import "sync"

/**
simple queue struct
 */
type SimpleQueue struct {
	Qs []string
	Size int
	Lock sync.Mutex
}

/*
* add data
 */
func (self *SimpleQueue) QueueAdd(data string) {
	self.Lock.Lock()
	defer self.Lock.Unlock()
	self.Qs = append(self.Qs, data)
	self.Size += 1
}

/*
pop data
 */
func (self *SimpleQueue) QueuePop() string {
	self.Lock.Lock()
	defer self.Lock.Unlock()
	if self.Size == 0 {
		return ""
	}
	v := self.Qs[0]
	self.Qs = self.Qs[1:]
	self.Size -= 1
	return v
}

/*
* get all queue data
 */
func (self *SimpleQueue) GetAllQueueData() []string{
	self.Lock.Lock()
	defer self.Lock.Unlock()
	v := self.Qs
	return v
}

