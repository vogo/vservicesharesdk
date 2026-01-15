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

// SyncFaceAuthRequest represents the request for syncing face recognition record.
type SyncFaceAuthRequest struct {
	Name        string            `json:"name"`        // the freelancer's full name (max 25 chars)
	IdCard      string            `json:"idCard"`      // the ID card number (18 chars)
	Mobile      string            `json:"mobile"`      // the phone number registered with bank (11 chars)
	ThirdId     string            `json:"thirdId"`     // the unique traceable code for face authentication (max 50 chars)
	AuthTime    string            `json:"authTime"`    // the face authentication completion time (format: yyyy-MM-dd HH:mm:ss)
	Urls        []string          `json:"urls"`        // the face recognition photo/video URLs (max 2MB each, supports jpg/png/jpeg for images, mp4 for videos)
	AuthChannel cores.AuthChannel `json:"authChannel"` // the authentication channel (see cores.AuthChannel constants)
}

// SyncFaceAuthResponse represents the response for syncing face recognition record.
type SyncFaceAuthResponse struct {
	FaceAuthEndTime string `json:"faceAuthEndTime,omitempty"` // the face authentication expiration date (format: YYYY-MM-DD)
}

// SyncFaceAuth syncs a face recognition record for a freelancer.
// This API requires the user to have a successfully signed contract.
//
// Note: If error code 6323 is returned, it means a face authentication record already exists
// and no sync is needed.
func (s *Service) SyncFaceAuth(req *SyncFaceAuthRequest) (*SyncFaceAuthResponse, error) {
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
	if req.ThirdId == "" {
		return nil, fmt.Errorf("thirdId is required")
	}
	if req.AuthTime == "" {
		return nil, fmt.Errorf("authTime is required")
	}
	if len(req.Urls) == 0 {
		return nil, fmt.Errorf("urls is required")
	}
	if req.AuthChannel == "" {
		return nil, fmt.Errorf("authChannel is required")
	}

	// Call API with function code 6008
	respData, err := s.client.Do(cores.FunCodeSyncFaceAuth, req)
	if err != nil {
		return nil, err
	}

	// Handle empty response
	if respData == "" {
		return &SyncFaceAuthResponse{}, nil
	}

	// Unmarshal decrypted response
	var resp SyncFaceAuthResponse
	if err := json.Unmarshal([]byte(respData), &resp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &resp, nil
}
