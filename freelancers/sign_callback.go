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

package freelancers

import (
	"encoding/json"
	"fmt"

	"github.com/vogo/vogo/vlog"
)

// ParseSignCallback parses and validates the contract signing callback request.
// It takes the raw JSON body of the callback request.
func (s *Service) ParseSignCallback(body []byte) (*SignContractResult, error) {
	// Verify and decrypt the notification
	decryptedData, err := s.client.VerifyAndDecryptNotification(body)
	if err != nil {
		return nil, err
	}

	vlog.Infof("service share contract sign callback | data: %s", decryptedData)

	if decryptedData == "" {
		return nil, fmt.Errorf("empty callback data")
	}

	// Unmarshal decrypted data
	var result SignContractResult
	if err := json.Unmarshal([]byte(decryptedData), &result); err != nil {
		return nil, fmt.Errorf("failed to parse callback data: %w", err)
	}

	return &result, nil
}
