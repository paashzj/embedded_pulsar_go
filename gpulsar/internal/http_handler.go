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
	"embedded_pulsar_go/gpulsar/internal/tenants"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TenantPutHandler(c *gin.Context) {
	tenant := c.Param("tenants")
	tenants.AddTenant(tenant)
	c.Status(http.StatusNoContent)
}

func TenantDeleteHandler(c *gin.Context) {
	tenant := c.Param("tenants")
	tenants.DelTenant(tenant)
	c.Status(http.StatusNoContent)
}

func TenantsGetHandler(c *gin.Context) {
	c.JSON(http.StatusOK, tenants.GetTenants())
}

func TenantGetHandler(c *gin.Context) {
	tenant := c.Param("tenants")
	c.JSON(http.StatusOK, tenants.GetTenant(tenant))
}
