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
	"encoding/json"
	"fmt"
)

// RequestMessage represents the API request envelope.
type RequestMessage struct {
	ReqId   string `json:"reqId"`   // Unique request identifier
	FunCode string `json:"funCode"` // Function/operation code
	MerId   string `json:"merId"`   // Merchant ID
	Version string `json:"version"` // API version
	ReqData string `json:"reqData"` // DES-encrypted business data (Hex)
	Sign    string `json:"sign"`    // RSA signature (Base64)
}

// ResponseMessage represents the API response envelope.
type ResponseMessage struct {
	ReqId   string `json:"reqId"`   // Echo of request identifier
	FunCode string `json:"funCode"` // Function/operation code
	MerId   string `json:"merId"`   // Merchant ID
	Version string `json:"version"` // API version
	ResData string `json:"resData"` // DES-encrypted response data (Hex)
	ResCode string `json:"resCode"` // Response code ("0000" = success)
	ResMsg  string `json:"resMsg"`  // Human-readable response message
	Sign    string `json:"sign"`    // RSA signature (Base64)
}

// ToJSON marshals the RequestMessage to JSON.
func (r *RequestMessage) ToJSON() ([]byte, error) {
	data, err := json.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}
	return data, nil
}

// ParseResponseMessage parses JSON into ResponseMessage.
func ParseResponseMessage(data []byte) (*ResponseMessage, error) {
	var resp ResponseMessage
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &resp, nil
}

// IsSuccess checks if the response indicates success.
func (r *ResponseMessage) IsSuccess() bool {
	return r.ResCode == "0000"
}

// GetError returns an APIError if the response is not successful.
func (r *ResponseMessage) GetError() error {
	if r.IsSuccess() {
		return nil
	}
	return NewAPIError(r.ResCode, r.ResMsg)
}
