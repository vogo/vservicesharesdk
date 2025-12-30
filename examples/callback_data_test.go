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
	"testing"

	"github.com/vogo/vservicesharesdk/payments"
)

func TestCallbackData(t *testing.T) {
	data := `{"funCode":"6001","merId":"1763878666094686","resCode":"0000","resData":"2kpWZJuJDWop20JQa8qIscfTlG3HqxYeK9K7ZW7GHvgSOrQy8L4bbCRujLo0dmYxlHscnBDioMCvwAYL31tdQjI13NG3y9FV59g/rMqT/f+CqnBaUrdn8IJvPZxcrz1xNeFoaV0sIn529CR4b3ETPPlFf7a7OJh+2tSIBtLhuQ11jA2N9sJnJ43vVWXKgTC01bTA2tCOwavb3lwdjnl1CIEXXlFxXkLKCIQ2qWT48t1RGXXCtLehkoRx3TeHwFvbbawGqaQbsUw=","resMsg":"成功","sign":"AvqJSNNgfnMkSzIloN3I199XQX2sUt35/DxuEArubUkP4E5qWkBq/s4/rOf98lfFgTlQ/bH7NSLj25iGhfAHbrumOY4waP7nf2UA/zzThwyRHsA60uNio/i8WH+Cg476MM9g5+2eWU6CsnX/uKXPV5A7mKYljBmDxmYlN+f7GiM=","version":"V1.0"}`

	// "{\"amt\":1002,\"createTime\":\"2025-12-29 16:59:20\",\"fee\":60,\"merOrderId\":\"SSOW202512291659171121\",\"orderNo\":2005564279308279809,\"resMsg\":\"成功\",\"state\":4,\"userFeeRatio\":0.000000,\"vaAddTax\":0,\"vaTax\":0}"
	client := CreateClient(t)
	paymentService := payments.NewService(client)
	callback, err := paymentService.ParsePaymentCallback([]byte(data))
	if err != nil {
		t.Fatalf("failed to parse callback data: %v", err)
	}

	t.Logf("callback: %v", callback)
}
