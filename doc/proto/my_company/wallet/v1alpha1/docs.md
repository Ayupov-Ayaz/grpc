# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [aayupov/wallet/v1alpha1/resources.proto](#aayupov_wallet_v1alpha1_resources-proto)
    - [ResourceType](#aayupov-wallet-v1alpha1-ResourceType)
  
- [aayupov/wallet/v1alpha1/transactions.proto](#aayupov_wallet_v1alpha1_transactions-proto)
    - [BetRequest](#aayupov-wallet-v1alpha1-BetRequest)
    - [BetResponse](#aayupov-wallet-v1alpha1-BetResponse)
    - [WinRequest](#aayupov-wallet-v1alpha1-WinRequest)
    - [WinResponse](#aayupov-wallet-v1alpha1-WinResponse)
  
    - [TransactionService](#aayupov-wallet-v1alpha1-TransactionService)
  
- [aayupov/wallet/v1alpha1/wallet.proto](#aayupov_wallet_v1alpha1_wallet-proto)
    - [CreateWalletRequest](#aayupov-wallet-v1alpha1-CreateWalletRequest)
    - [CreateWalletResponse](#aayupov-wallet-v1alpha1-CreateWalletResponse)
    - [DeleteWalletRequest](#aayupov-wallet-v1alpha1-DeleteWalletRequest)
    - [DeleteWalletResponse](#aayupov-wallet-v1alpha1-DeleteWalletResponse)
    - [GetBalanceRequest](#aayupov-wallet-v1alpha1-GetBalanceRequest)
    - [GetBalanceResponse](#aayupov-wallet-v1alpha1-GetBalanceResponse)
  
    - [WalletService](#aayupov-wallet-v1alpha1-WalletService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="aayupov_wallet_v1alpha1_resources-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## aayupov/wallet/v1alpha1/resources.proto


 


<a name="aayupov-wallet-v1alpha1-ResourceType"></a>

### ResourceType


| Name | Number | Description |
| ---- | ------ | ----------- |
| RESOURCE_TYPE_UNSPECIFIED | 0 |  |
| RESOURCE_TYPE_WALLET | 1 |  |
| RESOURCE_TYPE_TRANSACTION | 2 |  |


 

 

 



<a name="aayupov_wallet_v1alpha1_transactions-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## aayupov/wallet/v1alpha1/transactions.proto



<a name="aayupov-wallet-v1alpha1-BetRequest"></a>

### BetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| user_id | [string](#string) |  |  |
| amount | [uint64](#uint64) |  |  |






<a name="aayupov-wallet-v1alpha1-BetResponse"></a>

### BetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transaction_id | [string](#string) |  |  |
| balance | [int64](#int64) |  |  |






<a name="aayupov-wallet-v1alpha1-WinRequest"></a>

### WinRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| bet_transaction_id | [string](#string) |  |  |
| user_id | [string](#string) |  |  |
| amount | [uint64](#uint64) |  |  |






<a name="aayupov-wallet-v1alpha1-WinResponse"></a>

### WinResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transaction_id | [string](#string) |  |  |
| balance | [int64](#int64) |  |  |





 

 

 


<a name="aayupov-wallet-v1alpha1-TransactionService"></a>

### TransactionService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Bet | [BetRequest](#aayupov-wallet-v1alpha1-BetRequest) | [BetResponse](#aayupov-wallet-v1alpha1-BetResponse) |  |
| Win | [WinRequest](#aayupov-wallet-v1alpha1-WinRequest) | [WinResponse](#aayupov-wallet-v1alpha1-WinResponse) |  |

 



<a name="aayupov_wallet_v1alpha1_wallet-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## aayupov/wallet/v1alpha1/wallet.proto



<a name="aayupov-wallet-v1alpha1-CreateWalletRequest"></a>

### CreateWalletRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  |  |
| balance | [int64](#int64) |  |  |






<a name="aayupov-wallet-v1alpha1-CreateWalletResponse"></a>

### CreateWalletResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| balance | [int64](#int64) |  |  |






<a name="aayupov-wallet-v1alpha1-DeleteWalletRequest"></a>

### DeleteWalletRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  |  |






<a name="aayupov-wallet-v1alpha1-DeleteWalletResponse"></a>

### DeleteWalletResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| wallet_id | [string](#string) |  |  |






<a name="aayupov-wallet-v1alpha1-GetBalanceRequest"></a>

### GetBalanceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  |  |






<a name="aayupov-wallet-v1alpha1-GetBalanceResponse"></a>

### GetBalanceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| wallet_id | [string](#string) |  |  |
| balance | [int64](#int64) |  |  |





 

 

 


<a name="aayupov-wallet-v1alpha1-WalletService"></a>

### WalletService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateWallet | [CreateWalletRequest](#aayupov-wallet-v1alpha1-CreateWalletRequest) | [CreateWalletResponse](#aayupov-wallet-v1alpha1-CreateWalletResponse) |  |
| GetBalance | [GetBalanceRequest](#aayupov-wallet-v1alpha1-GetBalanceRequest) | [GetBalanceResponse](#aayupov-wallet-v1alpha1-GetBalanceResponse) |  |
| DeleteWallet | [DeleteWalletRequest](#aayupov-wallet-v1alpha1-DeleteWalletRequest) | [DeleteWalletResponse](#aayupov-wallet-v1alpha1-DeleteWalletResponse) |  |

 



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

