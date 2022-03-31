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
	"embedded_pulsar_go/gpulsar/internal"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogo/protobuf/proto"
	pb "github.com/paashzj/pulsar_proto_go"
	"github.com/panjf2000/gnet"
	"github.com/sirupsen/logrus"
)

func Run(config Config) error {
	logrus.Info("begin to start embedded pulsar")
	runHttpServer(config)
	runTcpServer(config)
	return nil
}

func runHttpServer(config Config) {
	go func() {
		router := gin.Default()
		adminV2Router(router)
		err := router.Run(fmt.Sprintf("%s:%d", config.ListenHost, config.ListenHttpPort))
		logrus.Error("pulsar http server started error ", err)
	}()
}

func adminV2Router(rg *gin.Engine) {
	adminV2Group := rg.Group("/admin/v2")
	tenantRouter(adminV2Group)
	namespaceRouter(adminV2Group)
}

func tenantRouter(rg *gin.RouterGroup) *gin.RouterGroup {
	tenants := rg.Group("/tenants")
	tenants.PUT(":tenant", internal.TenantPutHandler)
	tenants.DELETE(":tenant", internal.TenantDeleteHandler)
	tenants.GET("", internal.TenantsGetHandler)
	tenants.GET(":tenant", internal.TenantGetHandler)
	return tenants
}

func namespaceRouter(rg *gin.RouterGroup) *gin.RouterGroup {
	tenants := rg.Group("/namespaces")
	tenants.PUT(":tenant/:namespace", internal.NamespacePutHandler)
	tenants.DELETE(":tenant/:namespace", internal.NamespaceDeleteHandler)
	tenants.GET(":tenant", internal.NamespacesGetHandler)
	tenants.GET(":tenant/:namespace", internal.NamespaceGetHandler)
	return tenants
}

func runTcpServer(config Config) {
	server := &Server{
		EventServer: nil,
		pulsarImpl:  internal.NewPulsarServer(),
	}
	go func() {
		err := gnet.Serve(server, fmt.Sprintf("tcp://%s:%d", config.ListenHost, config.ListenTcpPort), gnet.WithCodec(internal.Codec))
		logrus.Error("pulsar tcp server started error ", err)
	}()
}

type Server struct {
	*gnet.EventServer
	pulsarImpl internal.PulsarServer
}

func (s *Server) OnInitComplete(server gnet.Server) (action gnet.Action) {
	logrus.Info("Pulsar Server Started")
	return
}

func (s *Server) React(frame []byte, c gnet.Conn) ([]byte, gnet.Action) {
	cmd := &pb.BaseCommand{}
	err := proto.Unmarshal(frame[4:], cmd)
	if err != nil {
		logrus.Error("marshal request error ", err)
		return nil, gnet.Close
	}
	switch *cmd.Type {
	case pb.BaseCommand_CONNECT:
		connected, err := s.pulsarImpl.Connect(cmd.Connect)
		if err != nil {
			logrus.Error("execute error ", err)
			return nil, gnet.Close
		}
		marshal, err := connected.Marshal()
		if err != nil {
			logrus.Error("marshal error ", cmd.Type)
			return nil, gnet.Close
		}
		return marshal, gnet.None
	default:
		break
	}
	logrus.Error("unsupported protocol ", cmd.Type)
	return nil, gnet.Close
}

func (s *Server) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	logrus.Info("new connection connected ", " from ", c.RemoteAddr())
	return
}

func (s *Server) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	logrus.Info("connection closed from ", c.RemoteAddr())
	return
}
