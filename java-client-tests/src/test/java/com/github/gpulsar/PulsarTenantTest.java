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

import org.apache.pulsar.client.admin.PulsarAdminException;
import org.apache.pulsar.client.api.PulsarClientException;
import org.apache.pulsar.common.policies.data.TenantInfo;
import org.apache.pulsar.common.policies.data.TenantInfoImpl;
import org.junit.jupiter.api.AfterAll;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.Test;

import java.util.Collections;

public class PulsarTenantTest extends BaseTest {

    @BeforeAll
    public void beforeAll() throws PulsarClientException {
        super.init();
    }

    @Test
    public void tenantLifecycleTest() throws PulsarAdminException {
        Assertions.assertEquals(3, pulsarAdmin.tenants().getTenants().size());
        pulsarAdmin.tenants().createTenant("alice", TenantInfo.builder()
                .allowedClusters(Collections.singleton("standalone"))
                .build());
        Assertions.assertEquals(4, pulsarAdmin.tenants().getTenants().size());
        pulsarAdmin.tenants().deleteTenant("alice");
        Assertions.assertEquals(3, pulsarAdmin.tenants().getTenants().size());
    }

    @AfterAll
    public void afterAll() throws PulsarClientException {
        super.close();
    }

}
