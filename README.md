# ServiceShare SDK for Go

Go SDK for ServiceShare (佣金保) API integration.

## Features

- Secure DES-ECB encryption for request/response data
- RSA-SHA1 signing for request authentication
- Complete API coverage:
  - Account balance query (6003)
  - Freelancer silent contract signing (6010)
  - Freelancer contract status query (6011)
  - Merchant batch payment (6001)
  - Batch payment status query (6002)
- Flexible key formats: PEM or raw base64
- Clean architecture following Go best practices
- Type-safe API with comprehensive error handling

## Installation

```bash
go get github.com/vogo/vservicesharesdk
```

## Quick Start

```go
// Create configuration
config := cores.NewConfig(
    "http://testgateway.serviceshare.com/testapi/clientapi/clientBusiness/common",
    "YOUR_MERCHANT_ID",
    "12345678901234567890123456789012", // DES key (uses first 8 bytes)
    "YOUR_RSA_PRIVATE_KEY",              // PEM format or raw base64
    "YOUR_PLATFORM_PUBLIC_KEY",          // PEM format or raw base64
)

// Create client
client, err := cores.NewClient(config)
if err != nil {
    log.Fatal(err)
}

// Create service and query balance
accountService := accounts.NewService(client)
resp, err := accountService.QueryBalance(&accounts.BalanceQueryRequest{
    ProviderID: "YOUR_PROVIDER_ID",
})
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Balance: %s fen\n", resp.Balance)
```

## Configuration

### Config Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `BaseURL` | string | Yes | API endpoint URL |
| `MerchantID` | string | Yes | Merchant ID from platform |
| `DesKey` | string | Yes | DES encryption key (uses first 8 bytes) |
| `PrivateKey` | string | Yes | Merchant RSA private key (PEM or raw base64) |
| `PlatformPublicKey` | string | Yes | Platform RSA public key (PEM or raw base64) |
| `Version` | string | No | API version (default: "V1.0") |
| `Timeout` | time.Duration | No | HTTP timeout (default: 60s) |

### Key Formats

**RSA Keys** support two formats:
- **PEM format:** Standard format with `-----BEGIN/END-----` headers
- **Raw base64:** Base64-encoded DER format without headers

### Environment URLs

**Test Environment:**
```
http://testgateway.serviceshare.com/testapi/clientapi/clientBusiness/common
```

**Production Environment:**
Contact ServiceShare operations team for production URL.

## API Reference

### Accounts Service

**Balance Query (FunCode: 6003)**
```go
accountService := accounts.NewService(client)
resp, err := accountService.QueryBalance(&accounts.BalanceQueryRequest{
    ProviderID:  "YOUR_PROVIDER_ID",
    PaymentType: cores.PaymentTypeBankCard, // Optional
})
// resp.Balance in fen (1 yuan = 100 fen)
```

### Freelancers Service

**Silent Contract Signing (FunCode: 6010)**
```go
freelancerService := freelancers.NewService(client)
resp, err := freelancerService.SilentSign(&freelancers.SilentSignRequest{
    Name:        "张三",
    CardNo:      "6222021234567890123",
    IdCard:      "110101199001011234",
    Mobile:      "13800138000",
    PaymentType: cores.PaymentTypeBankCard,
    ProviderId:  "YOUR_PROVIDER_ID",
    IdCardPic1:  "HEX_ENCODED_FRONT_PHOTO",
    IdCardPic2:  "HEX_ENCODED_BACK_PHOTO",
})
// Asynchronous operation - check callback or use contract query
```

**Contract Status Query (FunCode: 6011)**
```go
resp, err := freelancerService.QuerySign(&freelancers.SignQueryRequest{
    Name:       "张三",
    IdCard:     "110101199001011234",
    Mobile:     "13800138000",
    ProviderId: "YOUR_PROVIDER_ID",
})
// resp.State: 0=unsigned, 1=signed, 2=not found, 3=pending, 4=failed, 5=cancelled
```

### Payments Service

**Batch Payment (FunCode: 6001)**
```go
paymentService := payments.NewService(client)
resp, err := paymentService.BatchPayment(&payments.BatchPaymentRequest{
    MerBatchId: "BATCH_001",
    PayItems: []payments.PaymentItem{
        {
            MerOrderId:  "ORDER_001",
            Amt:         10000, // 100 CNY in fen
            PayeeName:   "张三",
            PayeeAcc:    "6222021234567890123",
            IdCard:      "110101199001011234",
            Mobile:      "13800138000",
            PaymentType: cores.PaymentTypeBankCard,
        },
    },
    TaskId:     1001,
    ProviderId: "YOUR_PROVIDER_ID",
})
// resp.SuccessNum, resp.FailureNum, resp.PayResultList
// Note: Synchronous response only confirms receipt
```

**Batch Payment Query (FunCode: 6002)**
```go
resp, err := paymentService.QueryBatchPayment(&payments.BatchPaymentQueryRequest{
    MerBatchId: "BATCH_001",
    // Omit QueryItems to get all orders
})
// resp.QueryItems[].State: 1=processing, 3=success, 4=failed, 6=pending, 7=cancelled
```

## Handling Notifications

The SDK provides helpers to handle asynchronous callbacks from the platform.

### Contract Signing Notification (FunCode: 6010/5.1.4)
```go
// In your HTTP handler
body, _ := io.ReadAll(r.Body)
callback, err := freelancerService.ParseSignCallback(body)
if err != nil {
    // Handle error
    return
}
fmt.Printf("Sign Result: ContractID=%s Status=%s\n", callback.ContractId, callback.Status)
```

### Batch Payment Notification (FunCode: 6001/5.3.4)
```go
// In your HTTP handler
body, _ := io.ReadAll(r.Body)
callback, err := paymentService.ParseBatchPaymentCallback(body)
if err != nil {
    // Handle error
    return
}
fmt.Printf("Batch Payment: BatchID=%s Items=%d\n", callback.MerBatchId, len(callback.QueryItems))
```

## Error Handling

```go
resp, err := accountService.QueryBalance(req)
if err != nil {
    if apiErr, ok := err.(*cores.APIError); ok {
        // API error with Code and Message
        log.Printf("API Error [%s]: %s", apiErr.Code, apiErr.Message)
    }
    return err
}
```

**Common Error Codes:** `0000` (Success), `6001` (Parameter error), `6003` (Not found), `6006` (Signature failed), `6007` (Decryption failed), `6019` (Insufficient balance)

## Security Best Practices

- **Never hardcode keys** - Use environment variables or secret management services
- **Use HTTPS in production** - Test environment uses HTTP, production must use HTTPS
- **Separate credentials** - Keep test and production keys separate
- **Secure key storage** - Use proper file permissions (0600) for key files

```go
// Load from environment variables
config := cores.NewConfig(
    os.Getenv("SS_API_URL"),
    os.Getenv("SS_MERCHANT_ID"),
    os.Getenv("SS_DES_KEY"),
    os.Getenv("SS_PRIVATE_KEY"),
    os.Getenv("SS_PLATFORM_PUBLIC_KEY"),
)
```

## Architecture

```
vservicesharesdk/
├── cores/          # Core SDK functionality
│   ├── client.go   # HTTP client with encryption/signing
│   ├── crypto.go   # DES encryption/decryption
│   ├── sign.go     # RSA signing/verification
│   ├── consts.go   # Constants (PaymentType, etc.)
│   └── errors.go   # Error types
├── accounts/       # Account service APIs (balance query)
├── freelancers/    # Freelancer APIs (signing, contract query)
├── payments/       # Payment APIs (batch payment, query)
└── examples/       # Usage examples with common helper
```

### Request Flow

1. Marshal request data to JSON
2. Encrypt JSON with DES
3. Sign encrypted data with RSA private key
4. Send HTTP POST with encrypted + signed payload
5. Verify response signature with platform public key
6. Decrypt response data with DES
7. Return typed response

## Testing

For testing, you can use the demo credentials from:
https://gitee.com/bubibi1/bosskg-demo

## Contributing

Contributions are welcome! Please ensure:
- Code follows Go conventions and best practices
- All tests pass
- Documentation is updated
- Commits are clear and descriptive

## License

Apache License 2.0

## Support

For API documentation and support:
- API Docs: https://apidoc.serviceshare.com/bosskg.md
- Demo Code: https://gitee.com/bubibi1/bosskg-demo
