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

// PaymentState represents the payment transaction state.
type PaymentState int

const (
	PaymentStateProcessing     PaymentState = 1 // indicates the payment is being processed
	PaymentStateSuccess        PaymentState = 3 // indicates the payment succeeded
	PaymentStateFailed         PaymentState = 4 // indicates the payment failed
	PaymentStatePendingConfirm PaymentState = 6 // indicates awaiting user confirmation
	PaymentStateCancelled      PaymentState = 7 // indicates the payment was cancelled
)

// PaymentResult represents the detailed result of a payment query.
type PaymentBaseResult struct {
	MerOrderId   string       `json:"merOrderId"`   // the merchant order ID
	State        PaymentState `json:"state"`        // the payment state
	Amt          int64        `json:"amt"`          // the payment amount in fen
	Fee          int64        `json:"fee"`          // the service fee in fen
	UserFee      int64        `json:"userFee"`      // the user's fee in fen
	Tax          int64        `json:"tax"`          // the tax amount in fen
	UserDueAmt   int64        `json:"userDueAmt"`   // the amount due to user in fen
	UserFeeRatio float64      `json:"userFeeRatio"` // the user's fee ratio
	VaTax        int64        `json:"vaTax"`        // the VAT tax amount in fen
	VaAddTax     int64        `json:"vaAddTax"`     // the VAT additional tax amount in fen
	CreateTime   string       `json:"createTime"`   // the order creation time (format: yyyy-MM-dd HH:mm:ss)
	EndTime      string       `json:"endTime"`      // the transaction completion time (format: yyyy-MM-dd HH:mm:ss)

	ResCode     string `json:"resCode"`     // the result code
	ResMsg      string `json:"resMsg"`      // the result message
	PackageInfo string `json:"packageInfo"` // the package information for wechat pay 零钱
	MchId       string `json:"mchId"`       // the merchant ID for wechat pay 零钱
}

// PaymentResult represents the detailed result of a payment query.
type PaymentResult struct {
	PaymentBaseResult
	OrderNo int64 `json:"orderNo"` // the platform order number (use as primary transaction identifier)
}

// PaymentExecuteResponse represents the detailed result of a payment query.
type PaymentExecuteResponse struct {
	PaymentBaseResult
	OrderNo string `json:"orderNo"` // the platform order number (v1.0版本下单返回是字符串)
}

// BatchPaymentResult represents the response for batch payment query.
type BatchPaymentResult struct {
	MerId      string          `json:"merId"`      // merchant ID
	MerBatchId string          `json:"merBatchId"` // merchant batch number
	QueryItems []PaymentResult `json:"queryItems"` // query items
}
