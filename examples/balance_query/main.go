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

package main

import (
	"fmt"
	"log"

	"github.com/vogo/vogo/vos"
	"github.com/vogo/vservicesharesdk/accounts"
	"github.com/vogo/vservicesharesdk/cores"
	"github.com/vogo/vservicesharesdk/examples/common"
)

func main() {
	// Create client from environment variables
	client := common.CreateClient()

	// Create accounts service
	accountService := accounts.NewService(client)

	// Query balance
	resp, err := accountService.QueryBalance(&accounts.BalanceQueryRequest{
		ProviderID:  vos.EnvString("SS_PROVIDER_ID"),
		PaymentType: cores.PaymentTypeBankCard, // Optional: 0=Bank, 1=Alipay, 2=WeChat
	})
	if err != nil {
		log.Fatalf("Failed to query balance: %v", err)
	}

	// Parse balance from fen to yuan
	var balanceFen int64
	fmt.Sscanf(resp.Balance, "%d", &balanceFen)
	balanceYuan := float64(balanceFen) / 100.0

	fmt.Printf("Account Balance Query Result:\n")
	fmt.Printf("  Provider ID: %s\n", resp.ProviderID)
	fmt.Printf("  Balance: %s fen (%.2f CNY)\n", resp.Balance, balanceYuan)
}
