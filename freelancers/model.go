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

// SignResult represents the result of sign.
type SignResult struct {
	Name       string    `json:"name"`             // the freelancer's name
	CardNo     string    `json:"cardNo"`           // the bank card number or payment account
	IdCard     string    `json:"idCard"`           // the ID card number
	Mobile     string    `json:"mobile"`           // the phone number registered with bank
	State      SignState `json:"state"`            // the sign status
	OtherParam string    `json:"otherParam"`       // other parameters
	ProviderId int64     `json:"providerId"`       // the service provider ID
	RetMsg     string    `json:"retMsg,omitempty"` // the failure reason if applicable
}
