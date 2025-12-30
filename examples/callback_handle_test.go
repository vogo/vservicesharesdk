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
	"io"
	"log"
	"net/http"
	"testing"

	"github.com/vogo/vservicesharesdk/freelancers"
	"github.com/vogo/vservicesharesdk/payments"
)

func TestCallback(t *testing.T) {
	// Initialize client using environment variables
	// Make sure to set SS_MERCHANT_ID, SS_DES_KEY, SS_PRIVATE_KEY, SS_PLATFORM_PUBLIC_KEY
	client := CreateClient(t)
	freelancerService := freelancers.NewService(client)
	paymentService := payments.NewService(client)

	// Define sign callback handler
	http.HandleFunc("/callback/sign", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Read request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Failed to read body: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Parse callback
		callback, err := freelancerService.ParseSignContractCallback(body)
		if err != nil {
			log.Printf("Failed to parse sign callback: %v", err)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"resCode":"9999", "resMsg":"Failed"}`))
			return
		}

		fmt.Printf("Received Sign Callback:\n")
		fmt.Printf("Name: %s\n", callback.Name)
		fmt.Printf("CardNo: %s\n", callback.CardNo)
		fmt.Printf("State: %d\n", callback.State)
		if callback.RetMsg != "" {
			fmt.Printf("RetMsg: %s\n", callback.RetMsg)
		}

		// Return success response
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"resCode":"0000", "resMsg":"Success"}`))
	})

	// Define batch payment callback handler
	http.HandleFunc("/callback/payment", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Read request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("failed to read body | err: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Parse callback
		callback, err := paymentService.ParsePaymentCallback(body)
		if err != nil {
			log.Printf("failed to parse payment callback | err: %v", err)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"resCode":"9999", "resMsg":"Failed"}`))
			return
		}

		fmt.Printf("Received Payment Callback:\n")
		fmt.Printf("  OrderNo=%d Amt=%d State=%d\n", callback.OrderNo, callback.Amt, callback.State)

		// Return success response
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"resCode":"0000", "resMsg":"Success"}`))
	})

	// Start server
	fmt.Println("Server listening on :8080")
	fmt.Println("Use POST /callback/contract for contract signing notifications")
	fmt.Println("Use POST /callback/payment for batch payment notifications")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
