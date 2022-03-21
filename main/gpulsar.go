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

package main

import (
	"embedded_pulsar_go/gpulsar"
	"embedded_pulsar_go/main/util"
	"os"
	"os/signal"
)

func main() {
	config := gpulsar.Config{}
	config.ListenHost = util.GetEnvStr("GPULSAR_LISTEN_HOST", "0.0.0.0")
	config.ListenTcpPort = util.GetEnvInt("GPULSAR_LISTEN_TCP_PORT", 6650)
	config.ListenHttpPort = util.GetEnvInt("GPULSAR_LISTEN_HTTP_PORT", 8080)
	err := gpulsar.Run(config)
	if err != nil {
		panic(err)
	}
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	for {
		<-interrupt
		return
	}
}
