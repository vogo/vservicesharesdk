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

const (
	// FunCodeBatchPayment is the function code for batch payment API
	FunCodeBatchPayment = "6001"
)

// PaymentItem represents a single payment item in a batch.
type PaymentItem struct {
	// MerOrderId is the merchant order ID (required, max 32 chars, unique)
	MerOrderId string `json:"merOrderId"`

	// Amt is the payment amount in fen (required, 10-9800000, i.e., ¥0.1-¥98000)
	Amt int64 `json:"amt"`

	// PayeeName is the payee's name (required, max 25 chars)
	PayeeName string `json:"payeeName"`

	// PayeeAcc is the payee's account (bank card/Alipay/WeChat) (required, max 25 chars)
	PayeeAcc string `json:"payeeAcc"`

	// IdCard is the payee's ID card number (required, max 18 chars)
	IdCard string `json:"idCard"`

	// Mobile is the payee's phone number (required, 11 chars)
	Mobile string `json:"mobile"`

	// Memo is the payment note (optional, max 100 chars)
	Memo string `json:"memo,omitempty"`

	// PaymentType is the payment method (required)
	// 0 = Bank Card, 1 = Alipay, 2 = WeChat
	PaymentType cores.PaymentType `json:"paymentType"`

	// NotifyUrl is the callback URL for this order (optional)
	NotifyUrl string `json:"notifyUrl,omitempty"`
}

// BatchPaymentRequest represents the request for batch payment.
type BatchPaymentRequest struct {
	// MerBatchId is the merchant batch number (required, max 32 chars, unique)
	MerBatchId string `json:"merBatchId"`

	// PayItems is the list of payment items (required)
	PayItems []PaymentItem `json:"payItems"`

	// TaskId is the task code for payment reason (required)
	TaskId int `json:"taskId"`

	// ProviderId is the service provider ID (required)
	ProviderId string `json:"providerId"`
}

// PaymentResult represents the result of a single payment order.
type PaymentResult struct {
	// MerOrderId is the merchant order ID
	MerOrderId string `json:"merOrderId"`

	// OrderNo is the platform order number
	OrderNo string `json:"orderNo"`

	// Amt is the payment amount in fen
	Amt int64 `json:"amt"`

	// Fee is the service fee in fen
	Fee int64 `json:"fee"`

	// ResCode is the result code
	ResCode string `json:"resCode"`

	// ResMsg is the result message
	ResMsg string `json:"resMsg"`
}

// BatchPaymentResponse represents the response for batch payment.
type BatchPaymentResponse struct {
	// SuccessNum is the count of accepted orders
	SuccessNum int `json:"successNum"`

	// FailureNum is the count of rejected orders
	FailureNum int `json:"failureNum"`

	// MerBatchId is the merchant batch number
	MerBatchId string `json:"merBatchId"`

	// PayResultList contains individual order results
	PayResultList []PaymentResult `json:"payResultList"`
}

// BatchPayment processes batch payment transactions for multiple freelancers.
//
// IMPORTANT: The synchronous response only indicates that the system has received the request.
// It does NOT represent the final transaction status. Always verify the final status via
// async notifications or the query interface.
//
// Single transaction limits: ¥0.1 to ¥98,000 (10 to 9,800,000 fen).
func (s *Service) BatchPayment(req *BatchPaymentRequest) (*BatchPaymentResponse, error) {
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
	if req.ProviderId == "" {
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
	respData, err := s.client.Do(FunCodeBatchPayment, req)
	if err != nil {
		return nil, fmt.Errorf("batch payment failed: %w", err)
	}

	// Handle empty response
	if respData == "" {
		return nil, fmt.Errorf("empty response data")
	}

	// Unmarshal decrypted response
	var resp BatchPaymentResponse
	if err := json.Unmarshal([]byte(respData), &resp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &resp, nil
}
