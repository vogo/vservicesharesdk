/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package common

import (
	"log"
	"time"

	"github.com/vogo/vogo/vos"
	"github.com/vogo/vservicesharesdk/cores"
)

const (
	// DefaultAPIURL is the default test environment URL
	DefaultAPIURL = "http://testgateway.serviceshare.com/testapi/clientapi/clientBusiness/common"
	// DefaultTimeout is the default HTTP timeout
	DefaultTimeout = 30 * time.Second
)

// CreateClient creates a new ServiceShare API client from environment variables.
// It reads the following environment variables:
//   - SS_API_URL: API endpoint URL (defaults to test environment)
//   - SS_MERCHANT_ID: Merchant ID
//   - SS_AES_KEY: AES-256 encryption key (32 bytes)
//   - SS_PRIVATE_KEY: Merchant RSA private key
//   - SS_PLATFORM_PUBLIC_KEY: Platform RSA public key
func CreateClient() *cores.Client {
	apiURL := vos.EnvString("SS_API_URL")
	if apiURL == "" {
		apiURL = DefaultAPIURL
	}

	// Create configuration
	config := cores.NewConfig(
		apiURL,                                  // BaseURL
		vos.EnvString("SS_MERCHANT_ID"),         // MerchantID
		vos.EnvString("SS_AES_KEY"),             // AesKey
		vos.EnvString("SS_PRIVATE_KEY"),         // PrivateKey
		vos.EnvString("SS_PLATFORM_PUBLIC_KEY"), // PlatformPublicKey
	)

	// Set timeout
	config.Timeout = DefaultTimeout

	// Create client
	client, err := cores.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return client
}
