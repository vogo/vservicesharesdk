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

package payments

import (
	"encoding/json"
	"fmt"
)

// BatchPaymentCallback represents the data in the batch payment callback.
type BatchPaymentCallback struct {
	// MerId is the merchant ID
	MerId string `json:"merId"`

	// MerBatchId is the merchant batch number
	MerBatchId string `json:"merBatchId"`

	// QueryItems contains the transaction details
	// Reusing PaymentQueryResult from batch_payment_query.go as the structure is typically identical
	QueryItems []PaymentQueryResult `json:"queryItems"`
}

// ParseBatchPaymentCallback parses and validates the batch payment callback request.
// It takes the raw JSON body of the callback request.
func (s *Service) ParseBatchPaymentCallback(body []byte) (*BatchPaymentCallback, error) {
	// Verify and decrypt the notification
	decryptedData, err := s.client.VerifyAndDecryptNotification(body)
	if err != nil {
		return nil, fmt.Errorf("failed to verify and decrypt callback: %w", err)
	}

	if decryptedData == "" {
		return nil, fmt.Errorf("empty callback data")
	}

	// Unmarshal decrypted data
	var callback BatchPaymentCallback
	if err := json.Unmarshal([]byte(decryptedData), &callback); err != nil {
		return nil, fmt.Errorf("failed to parse callback data: %w", err)
	}

	return &callback, nil
}
