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

package examples

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/vogo/vogo/vos"
	"github.com/vogo/vservicesharesdk/cores"
	"github.com/vogo/vservicesharesdk/payments"
)

func TestBatchPayment(t *testing.T) {
	// Create client from environment variables
	client := CreateClient(t)

	// Create payments service
	paymentService := payments.NewService(client)

	// Generate unique batch ID
	batchId := fmt.Sprintf("BATCH_%d", time.Now().Unix())

	// Submit batch payment
	resp, err := paymentService.Payment(&payments.PaymentRequest{
		MerBatchId: batchId,
		PayItems: []payments.PaymentItem{
			{
				MerOrderId:  fmt.Sprintf("ORDER_%d_001", time.Now().Unix()),
				Amt:         10202, // CNY in fen
				PayeeName:   vos.EnvString("SS_FREELANCER_NAME"),
				PayeeAcc:    vos.EnvString("SS_FREELANCER_CARD_NO"),
				IdCard:      vos.EnvString("SS_FREELANCER_ID_CARD"),
				Mobile:      vos.EnvString("SS_FREELANCER_MOBILE"),
				PaymentType: cores.PaymentTypeBankCard,
				Memo:        "Freelance payment",
			},
		},
		TaskId:     vos.EnvInt("SS_TASK_ID"),
		ProviderId: vos.EnvInt64("SS_PROVIDER_ID"),
	})
	if err != nil {
		log.Fatalf("failed to submit batch payment | err: %v", err)
	}

	fmt.Printf("Batch Payment Submitted:\n")
	fmt.Printf("  Batch ID: %s\n", resp.MerBatchId)
	fmt.Printf("  Success Count: %d\n", resp.SuccessNum)
	fmt.Printf("  Failure Count: %d\n", resp.FailureNum)
	fmt.Printf("\nOrder Results:\n")
	for i, result := range resp.PayResultList {
		fmt.Printf("  [%d] Order: %s\n", i+1, result.MerOrderId)
		fmt.Printf("      Platform Order: %s\n", result.OrderNo)
		fmt.Printf("      Amount: %d fen (%.2f CNY)\n", result.Amt, float64(result.Amt)/100)
		fmt.Printf("      Fee: %d fen (%.2f CNY)\n", result.Fee, float64(result.Fee)/100)
		fmt.Printf("      Result: [%s] %s\n", result.ResCode, result.ResMsg)
	}

	fmt.Printf("\nIMPORTANT: This response only confirms receipt.\n")
	fmt.Printf("Use QueryBatchPayment to check final transaction status.\n")

	// Wait a bit then query the batch
	time.Sleep(2 * time.Second)
	fmt.Printf("\n--- Querying Batch Payment Status ---\n")

	queryResp, err := paymentService.PaymentQuery(&payments.PaymentQueryRequest{
		MerBatchId: batchId,
		// Omit QueryItems to get all orders
	})
	if err != nil {
		log.Fatalf("Failed to query batch payment: %v", err)
	}

	stateNames := map[payments.PaymentState]string{
		payments.PaymentStateProcessing:     "Processing",
		payments.PaymentStateSuccess:        "Success",
		payments.PaymentStateFailed:         "Failed",
		payments.PaymentStatePendingConfirm: "Pending Confirmation",
		payments.PaymentStateCancelled:      "Cancelled",
	}

	fmt.Printf("\nBatch Payment Query Result:\n")
	fmt.Printf("  Merchant ID: %s\n", queryResp.MerId)
	fmt.Printf("  Batch ID: %s\n", queryResp.MerBatchId)
	fmt.Printf("\nTransaction Details:\n")
	for i, item := range queryResp.QueryItems {
		fmt.Printf("  [%d] Order: %s\n", i+1, item.MerOrderId)
		fmt.Printf("      Platform Order: %d (use this as primary ID)\n", item.OrderNo)
		fmt.Printf("      State: %s (%d)\n", stateNames[item.State], item.State)
		fmt.Printf("      Amount: %d fen (%.2f CNY)\n", item.Amt, float64(item.Amt)/100)
		fmt.Printf("      Fee: %d fen\n", item.Fee)
		fmt.Printf("      User Due: %d fen (%.2f CNY)\n", item.UserDueAmt, float64(item.UserDueAmt)/100)
		fmt.Printf("      Result: [%s] %s\n", item.ResCode, item.ResMsg)
	}
}
