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

	"github.com/vogo/vservicesharesdk/cores"
)

// PaymentItem represents a single payment item in a batch.
type PaymentItem struct {
	MerOrderId  string            `json:"merOrderId"`          // the merchant order ID
	Amt         int64             `json:"amt"`                 // the payment amount in fen
	PayeeName   string            `json:"payeeName"`           // the payee's name
	PayeeAcc    string            `json:"payeeAcc"`            // the payee's account(bank card/Alipay/WeChat)
	IdCard      string            `json:"idCard"`              // the payee's ID card number
	Mobile      string            `json:"mobile"`              // the payee's phone number
	Memo        string            `json:"memo,omitempty"`      // the payment note
	PaymentType cores.PaymentType `json:"paymentType"`         // the payment method
	NotifyUrl   string            `json:"notifyUrl,omitempty"` // the callback URL for this order
}

// PaymentRequest represents the request for batch payment.
type PaymentRequest struct {
	MerBatchId string        `json:"merBatchId"` // the merchant batch number
	PayItems   []PaymentItem `json:"payItems"`   // the list of payment items
	TaskId     int64         `json:"taskId"`     // the task code for payment reason
	ProviderId int64         `json:"providerId"` // the service provider ID
}

// PaymentResponse represents the response for batch payment.
type PaymentResponse struct {
	SuccessNum    int                    `json:"successNum"`    // the count of accepted orders
	FailureNum    int                    `json:"failureNum"`    // the count of rejected orders
	MerBatchId    string                 `json:"merBatchId"`    // the merchant batch number
	PayResultList []PaymentExecuteResult `json:"payResultList"` // the list of payment results
}

// Payment processes batch payment transactions for multiple freelancers.
//
// IMPORTANT: The synchronous response only indicates that the system has received the request.
// It does NOT represent the final transaction status. Always verify the final status via
// async notifications or the query interface.
//
// Single transaction limits: ¥0.1 to ¥98,000 (10 to 9,800,000 fen).
func (s *Service) Payment(req *PaymentRequest) (*PaymentResponse, error) {
	// Validate request
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.MerBatchId == "" {
		return nil, fmt.Errorf("merBatchId is required")
	}
	if len(req.PayItems) == 0 {
		return nil, fmt.Errorf("payItems cannot be empty")
	}
	if req.ProviderId == 0 {
		return nil, fmt.Errorf("providerId is required")
	}

	// Validate each payment item
	for i, item := range req.PayItems {
		if item.MerOrderId == "" {
			return nil, fmt.Errorf("payItems[%d].merOrderId is required", i)
		}
		if item.Amt < 10 || item.Amt > 9800000 {
			return nil, fmt.Errorf("payItems[%d].amt must be between 10 and 9800000 fen", i)
		}
		if item.PayeeName == "" {
			return nil, fmt.Errorf("payItems[%d].payeeName is required", i)
		}
		if item.PayeeAcc == "" {
			return nil, fmt.Errorf("payItems[%d].payeeAcc is required", i)
		}
		if item.IdCard == "" {
			return nil, fmt.Errorf("payItems[%d].idCard is required", i)
		}
		if item.Mobile == "" {
			return nil, fmt.Errorf("payItems[%d].mobile is required", i)
		}
	}

	// Call API with function code 6001
	respData, err := s.client.Do(cores.FunCodePayment, req)
	if err != nil {
		return nil, err
	}

	// Handle empty response
	if respData == "" {
		return nil, fmt.Errorf("empty response data")
	}

	// Unmarshal decrypted response
	var resp PaymentResponse
	if err := json.Unmarshal([]byte(respData), &resp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &resp, nil
}
