// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package internal

import "sync"

type namespace struct {
	name                             string
	fullName                         string
	mutex                            sync.Mutex
	persistentPartitionedTopicMap    map[string]*persistentPartitionedTopic
	persistentTopicMap               map[string]*persistentTopic
	nonPersistentPartitionedTopicMap map[string]*nonPersistentPartitionedTopic
	nonPersistentTopicMap            map[string]*nonPersistentTopic
}

func (n *namespace) newPersistentPartitionedTopic(name string, partition int) *persistentPartitionedTopic {
	p := &persistentPartitionedTopic{}
	p.name = name
	p.partition = partition
	return p
}

func (n *namespace) newPersistentTopic(name string) *persistentTopic {
	p := &persistentTopic{}
	p.name = name
	return p
}

func (n *namespace) newNonPersistentPartitionedTopic(name string, partition int) *nonPersistentPartitionedTopic {
	p := &nonPersistentPartitionedTopic{}
	p.name = name
	p.partition = partition
	return p
}

func (n *namespace) newNonPersistentTopic(name string) *nonPersistentTopic {
	p := &nonPersistentTopic{}
	p.name = name
	return p
}

func (n *namespace) AddPersistentPartitionedTopic(name string, partition int) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	partitionedTopic := n.newPersistentPartitionedTopic(name, partition)
	n.persistentPartitionedTopicMap[name] = partitionedTopic
}

func (n *namespace) DelPersistentPartitionedTopic(name string) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	delete(n.persistentPartitionedTopicMap, name)
}

func (n *namespace) GetPersistentPartitionedTopic(name string) *persistentPartitionedTopic {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	return n.persistentPartitionedTopicMap[name]
}

func (n *namespace) GetPersistentPartitionedTopics() []*persistentPartitionedTopic {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	res := make([]*persistentPartitionedTopic, 0)
	for _, val := range n.persistentPartitionedTopicMap {
		res = append(res, val)
	}
	return res
}

func (n *namespace) AddPersistentTopic(name string) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	topic := n.newPersistentTopic(name)
	n.persistentTopicMap[name] = topic
}

func (n *namespace) DelPersistentTopic(name string) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	delete(n.persistentTopicMap, name)
}

func (n *namespace) GetPersistentTopic(name string) *persistentTopic {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	return n.persistentTopicMap[name]
}

func (n *namespace) GetPersistentTopics() []*persistentTopic {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	res := make([]*persistentTopic, 0)
	for _, val := range n.persistentTopicMap {
		res = append(res, val)
	}
	return res
}

func (n *namespace) AddNonPersistentPartitionedTopic(name string, partition int) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	partitionedTopic := n.newNonPersistentPartitionedTopic(name, partition)
	n.nonPersistentPartitionedTopicMap[name] = partitionedTopic
}

func (n *namespace) DelNonPersistentPartitionedTopic(name string) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	delete(n.nonPersistentPartitionedTopicMap, name)
}

func (n *namespace) GetNonPersistentPartitionedTopic(name string) *nonPersistentPartitionedTopic {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	return n.nonPersistentPartitionedTopicMap[name]
}

func (n *namespace) GetNonPersistentPartitionedTopics() []*nonPersistentPartitionedTopic {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	res := make([]*nonPersistentPartitionedTopic, 0)
	for _, val := range n.nonPersistentPartitionedTopicMap {
		res = append(res, val)
	}
	return res
}

func (n *namespace) AddNonPersistentTopic(name string) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	topic := n.newNonPersistentTopic(name)
	n.nonPersistentTopicMap[name] = topic
}

func (n *namespace) DelNonPersistentTopic(name string) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	delete(n.nonPersistentTopicMap, name)
}

func (n *namespace) GetNonPersistentTopic(name string) *nonPersistentTopic {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	return n.nonPersistentTopicMap[name]
}

func (n *namespace) GetNonPersistentTopics() []*nonPersistentTopic {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	res := make([]*nonPersistentTopic, 0)
	for _, val := range n.nonPersistentTopicMap {
		res = append(res, val)
	}
	return res
}
