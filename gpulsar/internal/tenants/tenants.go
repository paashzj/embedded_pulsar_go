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

package tenants

import "sync"

var (
	mutex     sync.RWMutex
	tenantMap map[string]string
)

func init() {
	tenantMap = make(map[string]string)
	tenantMap["public"] = "public"
	tenantMap["pulsar"] = "pulsar"
	tenantMap["sample"] = "sample"
}

func AddTenant(tenant string) {
	mutex.Lock()
	tenantMap[tenant] = tenant
	mutex.Unlock()
}

func DelTenant(tenant string) {
	mutex.Lock()
	delete(tenantMap, tenant)
	mutex.Unlock()
}

func GetTenant(tenant string) string {
	mutex.RLock()
	defer mutex.RUnlock()
	return tenantMap[tenant]
}

func GetTenants() []string {
	mutex.RLock()
	defer mutex.RUnlock()
	res := make([]string, 0)
	for _, value := range tenantMap {
		res = append(res, value)
	}
	return res
}
