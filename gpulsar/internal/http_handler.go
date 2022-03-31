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
	"github.com/gin-gonic/gin"
	"net/http"
)

func TenantPutHandler(c *gin.Context) {
	tenant := c.Param("tenant")
	AddTenant(tenant)
	c.Status(http.StatusNoContent)
}

func TenantDeleteHandler(c *gin.Context) {
	tenant := c.Param("tenant")
	DelTenant(tenant)
	c.Status(http.StatusNoContent)
}

func TenantsGetHandler(c *gin.Context) {
	c.JSON(http.StatusOK, GetTenantNameList())
}

func TenantGetHandler(c *gin.Context) {
	tenant := c.Param("tenant")
	c.JSON(http.StatusOK, GetTenant(tenant))
}

func NamespacePutHandler(c *gin.Context) {
	tenantParam := c.Param("tenant")
	namespace := c.Param("namespace")
	tenant := GetTenant(tenantParam)
	tenant.AddNamespace(tenant.newNamespace(namespace))
	c.Status(http.StatusNoContent)
}

func NamespaceDeleteHandler(c *gin.Context) {
	tenant := c.Param("tenant")
	namespace := c.Param("namespace")
	GetTenant(tenant).DelNamespace(namespace)
	c.Status(http.StatusNoContent)
}

func NamespacesGetHandler(c *gin.Context) {
	tenant := c.Param("tenant")
	namespaces := GetTenant(tenant).GetNamespaces()
	res := make([]string, 0)
	for _, val := range namespaces {
		res = append(res, val.fullName)
	}
	c.JSON(http.StatusOK, res)
}

func NamespaceGetHandler(c *gin.Context) {
	tenant := c.Param("tenant")
	namespace := c.Param("namespace")
	c.JSON(http.StatusOK, GetTenant(tenant).GetNamespace(namespace))
}
