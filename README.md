# Go API client for formance

Open, modular foundation for unique payments flows

# Introduction
This API is documented in **OpenAPI format**.

# Authentication
Formance Stack offers one forms of authentication:
  - OAuth2
OAuth2 - an open protocol to allow secure authorization in a simple
and standard method from web, mobile and desktop applications.
<SecurityDefinitions />


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: develop
- Package version: v1.0.0-beta.3
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://www.formance.com](https://www.formance.com)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import formance "github.com/formancehq/formance-sdk-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), formance.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), formance.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), formance.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), formance.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountsApi* | [**AddMetadataToAccount**](docs/AccountsApi.md#addmetadatatoaccount) | **Post** /api/ledger/{ledger}/accounts/{address}/metadata | Add metadata to an account.
*AccountsApi* | [**CountAccounts**](docs/AccountsApi.md#countaccounts) | **Head** /api/ledger/{ledger}/accounts | Count the accounts from a ledger.
*AccountsApi* | [**GetAccount**](docs/AccountsApi.md#getaccount) | **Get** /api/ledger/{ledger}/accounts/{address} | Get account by its address.
*AccountsApi* | [**ListAccounts**](docs/AccountsApi.md#listaccounts) | **Get** /api/ledger/{ledger}/accounts | List accounts from a ledger.
*BalancesApi* | [**GetBalances**](docs/BalancesApi.md#getbalances) | **Get** /api/ledger/{ledger}/balances | Get the balances from a ledger&#39;s account
*BalancesApi* | [**GetBalancesAggregated**](docs/BalancesApi.md#getbalancesaggregated) | **Get** /api/ledger/{ledger}/aggregate/balances | Get the aggregated balances from selected accounts
*ClientsApi* | [**AddScopeToClient**](docs/ClientsApi.md#addscopetoclient) | **Put** /api/auth/clients/{clientId}/scopes/{scopeId} | Add scope to client
*ClientsApi* | [**CreateClient**](docs/ClientsApi.md#createclient) | **Post** /api/auth/clients | Create client
*ClientsApi* | [**CreateSecret**](docs/ClientsApi.md#createsecret) | **Post** /api/auth/clients/{clientId}/secrets | Add a secret to a client
*ClientsApi* | [**DeleteClient**](docs/ClientsApi.md#deleteclient) | **Delete** /api/auth/clients/{clientId} | Delete client
*ClientsApi* | [**DeleteScopeFromClient**](docs/ClientsApi.md#deletescopefromclient) | **Delete** /api/auth/clients/{clientId}/scopes/{scopeId} | Delete scope from client
*ClientsApi* | [**DeleteSecret**](docs/ClientsApi.md#deletesecret) | **Delete** /api/auth/clients/{clientId}/secrets/{secretId} | Delete a secret from a client
*ClientsApi* | [**ListClients**](docs/ClientsApi.md#listclients) | **Get** /api/auth/clients | List clients
*ClientsApi* | [**ReadClient**](docs/ClientsApi.md#readclient) | **Get** /api/auth/clients/{clientId} | Read client
*ClientsApi* | [**UpdateClient**](docs/ClientsApi.md#updateclient) | **Put** /api/auth/clients/{clientId} | Update client
*MappingApi* | [**GetMapping**](docs/MappingApi.md#getmapping) | **Get** /api/ledger/{ledger}/mapping | Get the mapping of a ledger.
*MappingApi* | [**UpdateMapping**](docs/MappingApi.md#updatemapping) | **Put** /api/ledger/{ledger}/mapping | Update the mapping of a ledger.
*PaymentsApi* | [**ConnectorsStripeTransfer**](docs/PaymentsApi.md#connectorsstripetransfer) | **Post** /api/payments/connectors/stripe/transfer | Transfer funds between Stripe accounts
*PaymentsApi* | [**GetAllConnectors**](docs/PaymentsApi.md#getallconnectors) | **Get** /api/payments/connectors | Get all installed connectors
*PaymentsApi* | [**GetAllConnectorsConfigs**](docs/PaymentsApi.md#getallconnectorsconfigs) | **Get** /api/payments/connectors/configs | Get all available connectors configs
*PaymentsApi* | [**GetConnectorTask**](docs/PaymentsApi.md#getconnectortask) | **Get** /api/payments/connectors/{connector}/tasks/{taskId} | Read a specific task of the connector
*PaymentsApi* | [**GetPayment**](docs/PaymentsApi.md#getpayment) | **Get** /api/payments/payments/{paymentId} | Returns a payment.
*PaymentsApi* | [**InstallConnector**](docs/PaymentsApi.md#installconnector) | **Post** /api/payments/connectors/{connector} | Install connector
*PaymentsApi* | [**ListConnectorTasks**](docs/PaymentsApi.md#listconnectortasks) | **Get** /api/payments/connectors/{connector}/tasks | List connector tasks
*PaymentsApi* | [**ListPayments**](docs/PaymentsApi.md#listpayments) | **Get** /api/payments/payments | Returns a list of payments.
*PaymentsApi* | [**ReadConnectorConfig**](docs/PaymentsApi.md#readconnectorconfig) | **Get** /api/payments/connectors/{connector}/config | Read connector config
*PaymentsApi* | [**ResetConnector**](docs/PaymentsApi.md#resetconnector) | **Post** /api/payments/connectors/{connector}/reset | Reset connector
*PaymentsApi* | [**UninstallConnector**](docs/PaymentsApi.md#uninstallconnector) | **Delete** /api/payments/connectors/{connector} | Uninstall connector
*ScopesApi* | [**AddTransientScope**](docs/ScopesApi.md#addtransientscope) | **Put** /api/auth/scopes/{scopeId}/transient/{transientScopeId} | Add a transient scope to a scope
*ScopesApi* | [**CreateScope**](docs/ScopesApi.md#createscope) | **Post** /api/auth/scopes | Create scope
*ScopesApi* | [**DeleteScope**](docs/ScopesApi.md#deletescope) | **Delete** /api/auth/scopes/{scopeId} | Delete scope
*ScopesApi* | [**DeleteTransientScope**](docs/ScopesApi.md#deletetransientscope) | **Delete** /api/auth/scopes/{scopeId}/transient/{transientScopeId} | Delete a transient scope from a scope
*ScopesApi* | [**ListScopes**](docs/ScopesApi.md#listscopes) | **Get** /api/auth/scopes | List scopes
*ScopesApi* | [**ReadScope**](docs/ScopesApi.md#readscope) | **Get** /api/auth/scopes/{scopeId} | Read scope
*ScopesApi* | [**UpdateScope**](docs/ScopesApi.md#updatescope) | **Put** /api/auth/scopes/{scopeId} | Update scope
*ScriptApi* | [**RunScript**](docs/ScriptApi.md#runscript) | **Post** /api/ledger/{ledger}/script | Execute a Numscript.
*SearchApi* | [**Search**](docs/SearchApi.md#search) | **Post** /api/search/ | Search
*ServerApi* | [**GetInfo**](docs/ServerApi.md#getinfo) | **Get** /api/ledger/_info | Show server information.
*StatsApi* | [**ReadStats**](docs/StatsApi.md#readstats) | **Get** /api/ledger/{ledger}/stats | Get Stats
*TransactionsApi* | [**AddMetadataOnTransaction**](docs/TransactionsApi.md#addmetadataontransaction) | **Post** /api/ledger/{ledger}/transactions/{txid}/metadata | Set the metadata of a transaction by its ID.
*TransactionsApi* | [**CountTransactions**](docs/TransactionsApi.md#counttransactions) | **Head** /api/ledger/{ledger}/transactions | Count the transactions from a ledger.
*TransactionsApi* | [**CreateTransaction**](docs/TransactionsApi.md#createtransaction) | **Post** /api/ledger/{ledger}/transactions | Create a new transaction to a ledger.
*TransactionsApi* | [**CreateTransactions**](docs/TransactionsApi.md#createtransactions) | **Post** /api/ledger/{ledger}/transactions/batch | Create a new batch of transactions to a ledger.
*TransactionsApi* | [**GetTransaction**](docs/TransactionsApi.md#gettransaction) | **Get** /api/ledger/{ledger}/transactions/{txid} | Get transaction from a ledger by its ID.
*TransactionsApi* | [**ListTransactions**](docs/TransactionsApi.md#listtransactions) | **Get** /api/ledger/{ledger}/transactions | List transactions from a ledger.
*TransactionsApi* | [**RevertTransaction**](docs/TransactionsApi.md#reverttransaction) | **Post** /api/ledger/{ledger}/transactions/{txid}/revert | Revert a ledger transaction by its ID.
*UsersApi* | [**ListUsers**](docs/UsersApi.md#listusers) | **Get** /api/auth/users | List users
*UsersApi* | [**ReadUser**](docs/UsersApi.md#readuser) | **Get** /api/auth/users/{userId} | Read user
*WebhooksApi* | [**ActivateOneConfig**](docs/WebhooksApi.md#activateoneconfig) | **Put** /api/webhooks/configs/{id}/activate | Activate one config
*WebhooksApi* | [**ChangeOneConfigSecret**](docs/WebhooksApi.md#changeoneconfigsecret) | **Put** /api/webhooks/configs/{id}/secret/change | Change the signing secret of a config
*WebhooksApi* | [**DeactivateOneConfig**](docs/WebhooksApi.md#deactivateoneconfig) | **Put** /api/webhooks/configs/{id}/deactivate | Deactivate one config
*WebhooksApi* | [**DeleteOneConfig**](docs/WebhooksApi.md#deleteoneconfig) | **Delete** /api/webhooks/configs/{id} | Delete one config
*WebhooksApi* | [**GetManyConfigs**](docs/WebhooksApi.md#getmanyconfigs) | **Get** /api/webhooks/configs | Get many configs
*WebhooksApi* | [**InsertOneConfig**](docs/WebhooksApi.md#insertoneconfig) | **Post** /api/webhooks/configs | Insert a new config 
*WebhooksApi* | [**TestOneConfig**](docs/WebhooksApi.md#testoneconfig) | **Get** /api/webhooks/configs/{id}/test | Test one config


## Documentation For Models

 - [Account](docs/Account.md)
 - [AccountWithVolumesAndBalances](docs/AccountWithVolumesAndBalances.md)
 - [AddMetadataToAccount409Response](docs/AddMetadataToAccount409Response.md)
 - [Attempt](docs/Attempt.md)
 - [AttemptResponse](docs/AttemptResponse.md)
 - [BankingCircleConfig](docs/BankingCircleConfig.md)
 - [ChangeOneConfigSecretRequest](docs/ChangeOneConfigSecretRequest.md)
 - [Client](docs/Client.md)
 - [ClientAllOf](docs/ClientAllOf.md)
 - [ClientOptions](docs/ClientOptions.md)
 - [ClientSecret](docs/ClientSecret.md)
 - [Config](docs/Config.md)
 - [ConfigInfo](docs/ConfigInfo.md)
 - [ConfigInfoResponse](docs/ConfigInfoResponse.md)
 - [ConfigResponse](docs/ConfigResponse.md)
 - [ConfigUser](docs/ConfigUser.md)
 - [ConnectorBaseInfo](docs/ConnectorBaseInfo.md)
 - [ConnectorConfig](docs/ConnectorConfig.md)
 - [ConnectorTask](docs/ConnectorTask.md)
 - [Contract](docs/Contract.md)
 - [CreateClientResponse](docs/CreateClientResponse.md)
 - [CreateScopeResponse](docs/CreateScopeResponse.md)
 - [CreateSecretResponse](docs/CreateSecretResponse.md)
 - [CreateTransaction400Response](docs/CreateTransaction400Response.md)
 - [CreateTransaction409Response](docs/CreateTransaction409Response.md)
 - [CreateTransactions400Response](docs/CreateTransactions400Response.md)
 - [CurrencyCloudConfig](docs/CurrencyCloudConfig.md)
 - [Cursor](docs/Cursor.md)
 - [DummyPayConfig](docs/DummyPayConfig.md)
 - [ErrorCode](docs/ErrorCode.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [GetAccount200Response](docs/GetAccount200Response.md)
 - [GetAccount400Response](docs/GetAccount400Response.md)
 - [GetBalances200Response](docs/GetBalances200Response.md)
 - [GetBalances200ResponseCursor](docs/GetBalances200ResponseCursor.md)
 - [GetBalances200ResponseCursorAllOf](docs/GetBalances200ResponseCursorAllOf.md)
 - [GetBalancesAggregated200Response](docs/GetBalancesAggregated200Response.md)
 - [GetBalancesAggregated400Response](docs/GetBalancesAggregated400Response.md)
 - [GetManyConfigs200Response](docs/GetManyConfigs200Response.md)
 - [GetManyConfigs200ResponseCursor](docs/GetManyConfigs200ResponseCursor.md)
 - [GetManyConfigs200ResponseCursorAllOf](docs/GetManyConfigs200ResponseCursorAllOf.md)
 - [GetPaymentResponse](docs/GetPaymentResponse.md)
 - [GetTransaction400Response](docs/GetTransaction400Response.md)
 - [GetTransaction404Response](docs/GetTransaction404Response.md)
 - [LedgerStorage](docs/LedgerStorage.md)
 - [ListAccounts200Response](docs/ListAccounts200Response.md)
 - [ListAccounts200ResponseCursor](docs/ListAccounts200ResponseCursor.md)
 - [ListAccounts200ResponseCursorAllOf](docs/ListAccounts200ResponseCursorAllOf.md)
 - [ListAccounts400Response](docs/ListAccounts400Response.md)
 - [ListClientsResponse](docs/ListClientsResponse.md)
 - [ListConnectorsConfigsResponse](docs/ListConnectorsConfigsResponse.md)
 - [ListConnectorsConfigsResponseConnector](docs/ListConnectorsConfigsResponseConnector.md)
 - [ListConnectorsConfigsResponseConnectorKey](docs/ListConnectorsConfigsResponseConnectorKey.md)
 - [ListConnectorsResponse](docs/ListConnectorsResponse.md)
 - [ListPaymentsResponse](docs/ListPaymentsResponse.md)
 - [ListScopesResponse](docs/ListScopesResponse.md)
 - [ListTransactions200Response](docs/ListTransactions200Response.md)
 - [ListTransactions200ResponseCursor](docs/ListTransactions200ResponseCursor.md)
 - [ListTransactions200ResponseCursorAllOf](docs/ListTransactions200ResponseCursorAllOf.md)
 - [ListUsersResponse](docs/ListUsersResponse.md)
 - [Mapping](docs/Mapping.md)
 - [MappingResponse](docs/MappingResponse.md)
 - [ModulrConfig](docs/ModulrConfig.md)
 - [Payment](docs/Payment.md)
 - [Posting](docs/Posting.md)
 - [Query](docs/Query.md)
 - [ReadClientResponse](docs/ReadClientResponse.md)
 - [ReadUserResponse](docs/ReadUserResponse.md)
 - [Response](docs/Response.md)
 - [ResponseCursor](docs/ResponseCursor.md)
 - [ResponseCursorTotal](docs/ResponseCursorTotal.md)
 - [RunScript400Response](docs/RunScript400Response.md)
 - [Scope](docs/Scope.md)
 - [ScopeAllOf](docs/ScopeAllOf.md)
 - [ScopeOptions](docs/ScopeOptions.md)
 - [Script](docs/Script.md)
 - [ScriptResult](docs/ScriptResult.md)
 - [Secret](docs/Secret.md)
 - [SecretAllOf](docs/SecretAllOf.md)
 - [SecretOptions](docs/SecretOptions.md)
 - [Stats](docs/Stats.md)
 - [StatsResponse](docs/StatsResponse.md)
 - [StripeConfig](docs/StripeConfig.md)
 - [StripeTask](docs/StripeTask.md)
 - [StripeTransferRequest](docs/StripeTransferRequest.md)
 - [Transaction](docs/Transaction.md)
 - [TransactionData](docs/TransactionData.md)
 - [TransactionResponse](docs/TransactionResponse.md)
 - [Transactions](docs/Transactions.md)
 - [TransactionsResponse](docs/TransactionsResponse.md)
 - [User](docs/User.md)
 - [Volume](docs/Volume.md)
 - [WebhooksConfig](docs/WebhooksConfig.md)
 - [WebhooksCursor](docs/WebhooksCursor.md)
 - [WiseConfig](docs/WiseConfig.md)


## Documentation For Authorization



### Authorization


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: N/A

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@formance.com

