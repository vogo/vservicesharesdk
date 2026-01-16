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
	"time"

	"github.com/vogo/vogo/vlog"
	"github.com/vogo/vservicesharesdk/cores"
)

// SignContractResult represents the result of sign.
type SignContractResult struct {
	Name            string              `json:"name"`                      // the freelancer's name
	CardNo          string              `json:"cardNo"`                    // the bank card number or payment account
	IdCard          string              `json:"idCard"`                    // the ID card number
	Mobile          string              `json:"mobile"`                    // the phone number registered with bank
	State           SignState           `json:"state"`                     // the sign status
	OtherParam      string              `json:"otherParam"`                // other parameters
	ProviderId      int64               `json:"providerId"`                // the service provider ID
	RetMsg          string              `json:"retMsg,omitempty"`          // the failure reason if applicable
	FaceAuthState   cores.FaceAuthState `json:"faceAuthState,omitempty"`   // the face auth status (UN_AUTH/PROCESS/SUCCESS/FAILED/EXPIRED)
	FaceAuthEndTime string              `json:"faceAuthEndTime,omitempty"` // the face auth expiration date (YYYY-MM-DD)
}

func (s SignContractResult) GetFaceAuthEndTime() *time.Time {
	return convertFaceAuthEndTime(s.FaceAuthEndTime)
}

func convertFaceAuthEndTime(s string) *time.Time {
	if s == "" {
		return nil
	}

	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		vlog.Errorf("failed to parse face auth end time | faceAuthEndTime: %s | err: %v", s, err)
		return nil
	}

	return &t
}
