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

import (
	"sync"
)

var (
	tenantsMutex sync.RWMutex
	tenantMap    map[string]*tenant
)

func init() {
	tenantMap = make(map[string]*tenant)
	publicTenant := newTenant("public")
	publicTenant.AddNamespace(publicTenant.newNamespace("default"))
	publicTenant.AddNamespace(publicTenant.newNamespace("functions"))
	tenantMap["public"] = publicTenant
	tenantMap["pulsar"] = newTenant("pulsar")
	tenantMap["sample"] = newTenant("sample")
}

type tenant struct {
	name         string
	namespaceMap map[string]*namespace
	mutex        sync.RWMutex
}

func newTenant(name string) *tenant {
	t := &tenant{}
	t.name = name
	t.namespaceMap = make(map[string]*namespace)
	return t
}

func AddTenant(name string) {
	tenantsMutex.Lock()
	tenantMap[name] = newTenant(name)
	tenantsMutex.Unlock()
}

func DelTenant(tenant string) {
	tenantsMutex.Lock()
	delete(tenantMap, tenant)
	tenantsMutex.Unlock()
}

func GetTenant(name string) *tenant {
	tenantsMutex.RLock()
	defer tenantsMutex.RUnlock()
	return tenantMap[name]
}

func GetTenants() []*tenant {
	tenantsMutex.RLock()
	defer tenantsMutex.RUnlock()
	res := make([]*tenant, 0)
	for _, value := range tenantMap {
		res = append(res, value)
	}
	return res
}

func GetTenantNameList() []string {
	tenants := GetTenants()
	res := make([]string, 0)
	for _, val := range tenants {
		res = append(res, val.name)
	}
	return res
}

func (t *tenant) newNamespace(name string) *namespace {
	n := &namespace{}
	n.name = name
	n.fullName = t.name + "/" + name
	return n
}

func (t *tenant) AddNamespace(namespace *namespace) {
	t.mutex.Lock()
	t.namespaceMap[namespace.name] = namespace
	t.mutex.Unlock()
}

func (t *tenant) DelNamespace(name string) {
	t.mutex.Lock()
	delete(t.namespaceMap, name)
	t.mutex.Unlock()
}

func (t *tenant) GetNamespace(name string) *namespace {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return t.namespaceMap[name]
}

func (t *tenant) GetNamespaces() []*namespace {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	res := make([]*namespace, 0)
	for _, value := range t.namespaceMap {
		res = append(res, value)
	}
	return res
}
