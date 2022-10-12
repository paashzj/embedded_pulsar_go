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

package gpulsar

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (p *PulsarServer) runHttpServer(config Config) {
	go func() {
		router := gin.Default()
		p.adminV2Router(router)
		err := router.Run(fmt.Sprintf("%s:%d", config.ListenHost, config.ListenHttpPort))
		logrus.Error("pulsar http server started error ", err)
	}()
}

func (p *PulsarServer) adminV2Router(rg *gin.Engine) {
	adminV2Group := rg.Group("/admin/v2")
	p.tenantRouter(adminV2Group)
	p.namespaceRouter(adminV2Group)
}

func (p *PulsarServer) tenantRouter(rg *gin.RouterGroup) *gin.RouterGroup {
	tenants := rg.Group("/tenants")
	tenants.PUT(":tenant", p.TenantPutHandler)
	tenants.DELETE(":tenant", p.TenantDeleteHandler)
	tenants.GET("", p.TenantsGetHandler)
	tenants.GET(":tenant", p.TenantGetHandler)
	return tenants
}

func (p *PulsarServer) namespaceRouter(rg *gin.RouterGroup) *gin.RouterGroup {
	tenants := rg.Group("/namespaces")
	tenants.PUT(":tenant/:namespace", p.NamespacePutHandler)
	tenants.DELETE(":tenant/:namespace", p.NamespaceDeleteHandler)
	tenants.GET(":tenant", p.NamespacesGetHandler)
	tenants.GET(":tenant/:namespace", p.NamespaceGetHandler)
	return tenants
}
