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

	"github.com/vogo/vservicesharesdk/cores"
)

// SilentSignRequest represents the request for freelancer silent contract signing.
type SilentSignRequest struct {
	// Name is the freelancer's full name (required, max 25 chars)
	Name string `json:"name"`

	// CardNo is the bank card number, Alipay account (phone/email), or WeChat OpenID (required, max 25 chars)
	CardNo string `json:"cardNo"`

	// IdCard is the ID card number, age typically 18-65 (required, max 18 chars)
	IdCard string `json:"idCard"`

	// Mobile is the phone number registered with bank (required, 11 chars)
	Mobile string `json:"mobile"`

	// PaymentType is the payment method (required)
	// 0 = Bank Card, 1 = Alipay, 2 = WeChat
	PaymentType cores.PaymentType `json:"paymentType"`

	// ProviderId is the service provider ID (required, max 20 chars)
	ProviderId int64 `json:"providerId"`

	// IdCardPic1 is the ID card front photo in hex format (required)
	IdCardPic1 string `json:"idCardPic1"`

	// IdCardPic2 is the ID card back photo in hex format (required)
	IdCardPic2 string `json:"idCardPic2"`

	// OtherParam is a pass-through parameter (optional)
	OtherParam string `json:"otherParam,omitempty"`

	// NotifyUrl is the callback URL for signing results (optional)
	NotifyUrl string `json:"notifyUrl,omitempty"`

	// TagList is an array of freelancer skill tags (optional)
	// Examples: design, marketing, etc.
	TagList []string `json:"tagList,omitempty"`
}

// SilentSignResponse represents the response for silent contract signing.
type SilentSignResponse struct {
	// OtherParam is the pass-through parameter returned (optional)
	OtherParam string `json:"otherParam,omitempty"`
}

// SilentSign initiates silent contract signing for a freelancer.
// This is an asynchronous operation; the synchronous response only confirms receipt.
// Results should be obtained through async notifications or signature query interface.
//
// Note: Contracts are validated by merchant ID + name + ID card + phone + provider ID.
func (s *Service) SilentSign(req *SilentSignRequest) (*SilentSignResponse, error) {
	// Validate request
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.Name == "" {
		return nil, fmt.Errorf("name is required")
	}
	if req.CardNo == "" {
		return nil, fmt.Errorf("cardNo is required")
	}
	if req.IdCard == "" {
		return nil, fmt.Errorf("idCard is required")
	}
	if req.Mobile == "" {
		return nil, fmt.Errorf("mobile is required")
	}
	if req.ProviderId == 0 {
		return nil, fmt.Errorf("providerId is required")
	}
	if req.IdCardPic1 == "" {
		return nil, fmt.Errorf("idCardPic1 is required")
	}
	if req.IdCardPic2 == "" {
		return nil, fmt.Errorf("idCardPic2 is required")
	}

	// Call API with function code 6010
	respData, err := s.client.Do(cores.FunCodeSilentSign, req)
	if err != nil {
		return nil, fmt.Errorf("silent sign failed: %w", err)
	}

	// Handle empty response (valid when otherParam is not provided)
	if respData == "" {
		return &SilentSignResponse{}, nil
	}

	// Unmarshal decrypted response
	var resp SilentSignResponse
	if err := json.Unmarshal([]byte(respData), &resp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &resp, nil
}
