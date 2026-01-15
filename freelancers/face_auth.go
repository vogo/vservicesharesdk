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

// FaceAuthRequest represents the request for face recognition authentication.
type FaceAuthRequest struct {
	Name            string `json:"name"`                      // the freelancer's full name (max 25 chars)
	IdCard          string `json:"idCard"`                    // the ID card number (18 chars)
	Mobile          string `json:"mobile"`                    // the phone number registered with bank (11 chars)
	RedirectUrl     string `json:"redirectUrl,omitempty"`     // the redirect URL after authentication (optional)
	RedirectBtnName string `json:"redirectBtnName,omitempty"` // the redirect button name on result page (optional)
}

// FaceAuthResponse represents the response for face recognition authentication.
type FaceAuthResponse struct {
	Url string `json:"url"` // the face recognition H5 page URL (valid for one day)
}

// FaceAuth initiates face recognition authentication for a freelancer.
// This API requires the user to have a successfully signed contract.
// Returns an H5 page URL for the user to complete face recognition (valid for one day).
func (s *Service) FaceAuth(req *FaceAuthRequest) (*FaceAuthResponse, error) {
	// Validate request
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.Name == "" {
		return nil, fmt.Errorf("name is required")
	}
	if req.IdCard == "" {
		return nil, fmt.Errorf("idCard is required")
	}
	if req.Mobile == "" {
		return nil, fmt.Errorf("mobile is required")
	}

	// Call API with function code 6009
	respData, err := s.client.Do(cores.FunCodeFaceAuth, req)
	if err != nil {
		return nil, err
	}

	// Unmarshal decrypted response
	var resp FaceAuthResponse
	if err := json.Unmarshal([]byte(respData), &resp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &resp, nil
}
