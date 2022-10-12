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
	"github.com/protocol-laboratory/pulsar-codec-go/pnet"
	"github.com/sirupsen/logrus"
	"sync"
)

type PulsarServer struct {
	config    *Config
	tcpServer *pnet.PulsarNetServer

	tenantsMutex sync.RWMutex
	tenantMap    map[string]*tenant
}

func NewPulsarServer(config *Config) (*PulsarServer, error) {
	pulsarServer := &PulsarServer{}
	pulsarServer.config = config
	tenantMap := make(map[string]*tenant)
	publicTenant := newTenant("public")
	publicTenant.AddNamespace(publicTenant.newNamespace("default"))
	publicTenant.AddNamespace(publicTenant.newNamespace("functions"))
	tenantMap["public"] = publicTenant
	tenantMap["pulsar"] = newTenant("pulsar")
	tenantMap["sample"] = newTenant("sample")
	pulsarServer.tenantMap = tenantMap
	return pulsarServer, nil
}

func (p *PulsarServer) Run(config Config) error {
	logrus.Info("begin to start embedded pulsar")
	p.runHttpServer(config)
	err := p.runTcpServer(config)
	if err != nil {
		return err
	}
	return nil
}

func (p *PulsarServer) runTcpServer(config Config) error {
	tcpServer, err := pnet.NewPulsarNetServer(pnet.PulsarNetServerConfig{
		Host:      config.ListenHost,
		Port:      config.ListenTcpPort,
		BufferMax: 0,
	}, p)
	if err != nil {
		return err
	}
	p.tcpServer = tcpServer
	return nil
}
