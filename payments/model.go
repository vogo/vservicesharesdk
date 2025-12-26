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

// PaymentQueryResult represents the detailed result of a payment query.
type PaymentQueryResult struct {
	// MerOrderId is the merchant order ID
	MerOrderId string `json:"merOrderId"`

	// OrderNo is the platform order number (use as primary transaction identifier)
	OrderNo int64 `json:"orderNo"`

	// State is the payment state
	State PaymentState `json:"state"`

	// Amt is the payment amount in fen
	Amt int64 `json:"amt"`

	// Fee is the service fee in fen
	Fee int64 `json:"fee"`

	// UserFee is the user's fee in fen
	UserFee int64 `json:"userFee"`

	// Tax is the tax amount in fen
	Tax int64 `json:"tax"`

	// UserDueAmt is the amount due to user in fen
	UserDueAmt int64 `json:"userDueAmt"`

	// ResCode is the result code
	ResCode string `json:"resCode"`

	// ResMsg is the result message
	ResMsg string `json:"resMsg"`
}
