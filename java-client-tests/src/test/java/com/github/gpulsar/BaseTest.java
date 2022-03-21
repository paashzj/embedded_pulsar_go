/**
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package com.github.gpulsar;

import org.apache.pulsar.client.admin.PulsarAdmin;
import org.apache.pulsar.client.api.PulsarClient;
import org.apache.pulsar.client.api.PulsarClientException;
import org.junit.jupiter.api.TestInstance;

@TestInstance(TestInstance.Lifecycle.PER_CLASS)
public class BaseTest {

    protected PulsarAdmin pulsarAdmin;

    protected PulsarClient pulsarClient;

    public void init() throws PulsarClientException {
        pulsarAdmin = PulsarAdmin.builder().serviceHttpUrl(PulsarConst.HTTP_URL).build();
        pulsarClient = PulsarClient.builder().serviceUrl(PulsarConst.TCP_URL).build();
    }

    public void close() throws PulsarClientException {
        pulsarAdmin.close();
        pulsarClient.close();
    }

}
