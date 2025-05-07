# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/v1/transactions.proto](#api_v1_transactions-proto)
    - [BetRequest](#api-v1-BetRequest)
    - [BetResponse](#api-v1-BetResponse)
    - [WinRequest](#api-v1-WinRequest)
    - [WinResponse](#api-v1-WinResponse)
  
    - [TransactionService](#api-v1-TransactionService)
  
- [api/v1/wallet.proto](#api_v1_wallet-proto)
    - [CreateWalletRequest](#api-v1-CreateWalletRequest)
    - [CreateWalletResponse](#api-v1-CreateWalletResponse)
    - [DeleteWalletRequest](#api-v1-DeleteWalletRequest)
    - [DeleteWalletResponse](#api-v1-DeleteWalletResponse)
    - [GetBalanceRequest](#api-v1-GetBalanceRequest)
    - [GetBalanceResponse](#api-v1-GetBalanceResponse)
  
    - [WalletService](#api-v1-WalletService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api_v1_transactions-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/transactions.proto



<a name="api-v1-BetRequest"></a>

### BetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| operation_id | [string](#string) |  |  |
| user_id | [string](#string) |  |  |
| amount | [uint64](#uint64) |  |  |






<a name="api-v1-BetResponse"></a>

### BetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| operation_id | [string](#string) |  |  |
| balance | [int64](#int64) |  |  |






<a name="api-v1-WinRequest"></a>

### WinRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| operation_id | [string](#string) |  |  |
| bet_operation_id | [string](#string) |  |  |
| user_id | [string](#string) |  |  |
| amount | [uint64](#uint64) |  |  |






<a name="api-v1-WinResponse"></a>

### WinResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| operation_id | [string](#string) |  |  |
| balance | [int64](#int64) |  |  |





 

 

 


<a name="api-v1-TransactionService"></a>

### TransactionService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Bet | [BetRequest](#api-v1-BetRequest) | [BetResponse](#api-v1-BetResponse) |  |
| Win | [WinRequest](#api-v1-WinRequest) | [WinResponse](#api-v1-WinResponse) |  |

 



<a name="api_v1_wallet-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/wallet.proto



<a name="api-v1-CreateWalletRequest"></a>

### CreateWalletRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  |  |
| balance | [int64](#int64) |  |  |






<a name="api-v1-CreateWalletResponse"></a>

### CreateWalletResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| balance | [int64](#int64) |  |  |






<a name="api-v1-DeleteWalletRequest"></a>

### DeleteWalletRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  |  |






<a name="api-v1-DeleteWalletResponse"></a>

### DeleteWalletResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| wallet_id | [string](#string) |  |  |






<a name="api-v1-GetBalanceRequest"></a>

### GetBalanceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  |  |






<a name="api-v1-GetBalanceResponse"></a>

### GetBalanceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| wallet_id | [string](#string) |  |  |
| balance | [int64](#int64) |  |  |





 

 

 


<a name="api-v1-WalletService"></a>

### WalletService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateWallet | [CreateWalletRequest](#api-v1-CreateWalletRequest) | [CreateWalletResponse](#api-v1-CreateWalletResponse) |  |
| GetBalance | [GetBalanceRequest](#api-v1-GetBalanceRequest) | [GetBalanceResponse](#api-v1-GetBalanceResponse) |  |
| DeleteWallet | [DeleteWalletRequest](#api-v1-DeleteWalletRequest) | [DeleteWalletResponse](#api-v1-DeleteWalletResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

