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

package cores

import (
	"fmt"
	"time"
)

// Config holds the configuration for the ServiceShare API client.
type Config struct {
	BaseURL           string        //  the API endpoint URL
	MerchantID        string        // the merchant identifier assigned by the platform
	Version           string        // the API version (default: "V1.0")
	DesKey            string        // the DES encryption key (uses first 8 bytes)
	PrivateKey        string        // the merchant's RSA private key in PEM format
	PlatformPublicKey string        // the platform's RSA public key in PEM format
	Timeout           time.Duration // the HTTP request timeout (default: 60 seconds)
	TaskID            int64         // the task identifier for the request
}

// NewConfig creates a new Config with default values.
func NewConfig(baseURL, merchantID, desKey, privateKey, platformPublicKey string, taskID int64) *Config {
	return &Config{
		BaseURL:           baseURL,
		MerchantID:        merchantID,
		Version:           "V1.0",
		DesKey:            desKey,
		PrivateKey:        privateKey,
		PlatformPublicKey: platformPublicKey,
		Timeout:           60 * time.Second,
		TaskID:            taskID,
	}
}

// Validate checks if the configuration is valid.
func (c *Config) Validate() error {
	if c.BaseURL == "" {
		return fmt.Errorf("%w: BaseURL is required", ErrInvalidConfig)
	}
	if c.MerchantID == "" {
		return fmt.Errorf("%w: MerchantID is required", ErrInvalidConfig)
	}
	if c.DesKey == "" {
		return fmt.Errorf("%w: DesKey is required", ErrInvalidConfig)
	}
	if len(c.DesKey) < 8 {
		return fmt.Errorf("%w: DesKey must be at least 8 bytes", ErrInvalidConfig)
	}
	if c.PrivateKey == "" {
		return fmt.Errorf("%w: PrivateKey is required", ErrInvalidConfig)
	}
	if c.PlatformPublicKey == "" {
		return fmt.Errorf("%w: PlatformPublicKey is required", ErrInvalidConfig)
	}
	if c.Version == "" {
		c.Version = "V1.0"
	}
	if c.Timeout == 0 {
		c.Timeout = 60 * time.Second
	}
	return nil
}
