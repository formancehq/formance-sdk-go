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

- API version: v1.0.20230228
- Package version: v1.0.20230228
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
*AccountsApi* | [**AddMetadataToAccount**](docs/AccountsApi.md#addmetadatatoaccount) | **Post** /api/ledger/{ledger}/accounts/{address}/metadata | Add metadata to an account
*AccountsApi* | [**CountAccounts**](docs/AccountsApi.md#countaccounts) | **Head** /api/ledger/{ledger}/accounts | Count the accounts from a ledger
*AccountsApi* | [**GetAccount**](docs/AccountsApi.md#getaccount) | **Get** /api/ledger/{ledger}/accounts/{address} | Get account by its address
*AccountsApi* | [**ListAccounts**](docs/AccountsApi.md#listaccounts) | **Get** /api/ledger/{ledger}/accounts | List accounts from a ledger
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
*DefaultApi* | [**GetServerInfo**](docs/DefaultApi.md#getserverinfo) | **Get** /api/auth/_info | Get server info
*DefaultApi* | [**PaymentsgetServerInfo**](docs/DefaultApi.md#paymentsgetserverinfo) | **Get** /api/payments/_info | Get server info
*DefaultApi* | [**SearchgetServerInfo**](docs/DefaultApi.md#searchgetserverinfo) | **Get** /api/search/_info | Get server info
*LedgerApi* | [**GetLedgerInfo**](docs/LedgerApi.md#getledgerinfo) | **Get** /api/ledger/{ledger}/_info | Get information about a ledger
*LogsApi* | [**ListLogs**](docs/LogsApi.md#listlogs) | **Get** /api/ledger/{ledger}/logs | List the logs from a ledger
*MappingApi* | [**GetMapping**](docs/MappingApi.md#getmapping) | **Get** /api/ledger/{ledger}/mapping | Get the mapping of a ledger
*MappingApi* | [**UpdateMapping**](docs/MappingApi.md#updatemapping) | **Put** /api/ledger/{ledger}/mapping | Update the mapping of a ledger
*OrchestrationApi* | [**CancelEvent**](docs/OrchestrationApi.md#cancelevent) | **Put** /api/orchestration/instances/{instanceID}/abort | Cancel a running workflow
*OrchestrationApi* | [**CreateWorkflow**](docs/OrchestrationApi.md#createworkflow) | **Post** /api/orchestration/workflows | Create workflow
*OrchestrationApi* | [**GetInstance**](docs/OrchestrationApi.md#getinstance) | **Get** /api/orchestration/instances/{instanceID} | Get a workflow instance by id
*OrchestrationApi* | [**GetInstanceHistory**](docs/OrchestrationApi.md#getinstancehistory) | **Get** /api/orchestration/instances/{instanceID}/history | Get a workflow instance history by id
*OrchestrationApi* | [**GetInstanceStageHistory**](docs/OrchestrationApi.md#getinstancestagehistory) | **Get** /api/orchestration/instances/{instanceID}/stages/{number}/history | Get a workflow instance stage history
*OrchestrationApi* | [**GetWorkflow**](docs/OrchestrationApi.md#getworkflow) | **Get** /api/orchestration/workflows/{flowId} | Get a flow by id
*OrchestrationApi* | [**ListInstances**](docs/OrchestrationApi.md#listinstances) | **Get** /api/orchestration/instances | List instances of a workflow
*OrchestrationApi* | [**ListWorkflows**](docs/OrchestrationApi.md#listworkflows) | **Get** /api/orchestration/workflows | List registered workflows
*OrchestrationApi* | [**OrchestrationgetServerInfo**](docs/OrchestrationApi.md#orchestrationgetserverinfo) | **Get** /api/orchestration/_info | Get server info
*OrchestrationApi* | [**RunWorkflow**](docs/OrchestrationApi.md#runworkflow) | **Post** /api/orchestration/workflows/{workflowID}/instances | Run workflow
*OrchestrationApi* | [**SendEvent**](docs/OrchestrationApi.md#sendevent) | **Post** /api/orchestration/instances/{instanceID}/events | Send an event to a running workflow
*PaymentsApi* | [**ConnectorsStripeTransfer**](docs/PaymentsApi.md#connectorsstripetransfer) | **Post** /api/payments/connectors/stripe/transfers | Transfer funds between Stripe accounts
*PaymentsApi* | [**ConnectorsTransfer**](docs/PaymentsApi.md#connectorstransfer) | **Post** /api/payments/connectors/{connector}/transfers | Transfer funds between Connector accounts
*PaymentsApi* | [**GetConnectorTask**](docs/PaymentsApi.md#getconnectortask) | **Get** /api/payments/connectors/{connector}/tasks/{taskId} | Read a specific task of the connector
*PaymentsApi* | [**GetPayment**](docs/PaymentsApi.md#getpayment) | **Get** /api/payments/payments/{paymentId} | Get a payment
*PaymentsApi* | [**InstallConnector**](docs/PaymentsApi.md#installconnector) | **Post** /api/payments/connectors/{connector} | Install a connector
*PaymentsApi* | [**ListAllConnectors**](docs/PaymentsApi.md#listallconnectors) | **Get** /api/payments/connectors | List all installed connectors
*PaymentsApi* | [**ListConfigsAvailableConnectors**](docs/PaymentsApi.md#listconfigsavailableconnectors) | **Get** /api/payments/connectors/configs | List the configs of each available connector
*PaymentsApi* | [**ListConnectorTasks**](docs/PaymentsApi.md#listconnectortasks) | **Get** /api/payments/connectors/{connector}/tasks | List tasks from a connector
*PaymentsApi* | [**ListConnectorsTransfers**](docs/PaymentsApi.md#listconnectorstransfers) | **Get** /api/payments/connectors/{connector}/transfers | List transfers and their statuses
*PaymentsApi* | [**ListPayments**](docs/PaymentsApi.md#listpayments) | **Get** /api/payments/payments | List payments
*PaymentsApi* | [**PaymentslistAccounts**](docs/PaymentsApi.md#paymentslistaccounts) | **Get** /api/payments/accounts | List accounts
*PaymentsApi* | [**ReadConnectorConfig**](docs/PaymentsApi.md#readconnectorconfig) | **Get** /api/payments/connectors/{connector}/config | Read the config of a connector
*PaymentsApi* | [**ResetConnector**](docs/PaymentsApi.md#resetconnector) | **Post** /api/payments/connectors/{connector}/reset | Reset a connector
*PaymentsApi* | [**UninstallConnector**](docs/PaymentsApi.md#uninstallconnector) | **Delete** /api/payments/connectors/{connector} | Uninstall a connector
*PaymentsApi* | [**UpdateMetadata**](docs/PaymentsApi.md#updatemetadata) | **Patch** /api/payments/payments/{paymentId}/metadata | Update metadata
*ScopesApi* | [**AddTransientScope**](docs/ScopesApi.md#addtransientscope) | **Put** /api/auth/scopes/{scopeId}/transient/{transientScopeId} | Add a transient scope to a scope
*ScopesApi* | [**CreateScope**](docs/ScopesApi.md#createscope) | **Post** /api/auth/scopes | Create scope
*ScopesApi* | [**DeleteScope**](docs/ScopesApi.md#deletescope) | **Delete** /api/auth/scopes/{scopeId} | Delete scope
*ScopesApi* | [**DeleteTransientScope**](docs/ScopesApi.md#deletetransientscope) | **Delete** /api/auth/scopes/{scopeId}/transient/{transientScopeId} | Delete a transient scope from a scope
*ScopesApi* | [**ListScopes**](docs/ScopesApi.md#listscopes) | **Get** /api/auth/scopes | List scopes
*ScopesApi* | [**ReadScope**](docs/ScopesApi.md#readscope) | **Get** /api/auth/scopes/{scopeId} | Read scope
*ScopesApi* | [**UpdateScope**](docs/ScopesApi.md#updatescope) | **Put** /api/auth/scopes/{scopeId} | Update scope
*ScriptApi* | [**RunScript**](docs/ScriptApi.md#runscript) | **Post** /api/ledger/{ledger}/script | Execute a Numscript
*SearchApi* | [**Search**](docs/SearchApi.md#search) | **Post** /api/search/ | Search
*ServerApi* | [**GetInfo**](docs/ServerApi.md#getinfo) | **Get** /api/ledger/_info | Show server information
*StatsApi* | [**ReadStats**](docs/StatsApi.md#readstats) | **Get** /api/ledger/{ledger}/stats | Get statistics from a ledger
*TransactionsApi* | [**AddMetadataOnTransaction**](docs/TransactionsApi.md#addmetadataontransaction) | **Post** /api/ledger/{ledger}/transactions/{txid}/metadata | Set the metadata of a transaction by its ID
*TransactionsApi* | [**CountTransactions**](docs/TransactionsApi.md#counttransactions) | **Head** /api/ledger/{ledger}/transactions | Count the transactions from a ledger
*TransactionsApi* | [**CreateTransaction**](docs/TransactionsApi.md#createtransaction) | **Post** /api/ledger/{ledger}/transactions | Create a new transaction to a ledger
*TransactionsApi* | [**CreateTransactions**](docs/TransactionsApi.md#createtransactions) | **Post** /api/ledger/{ledger}/transactions/batch | Create a new batch of transactions to a ledger
*TransactionsApi* | [**GetTransaction**](docs/TransactionsApi.md#gettransaction) | **Get** /api/ledger/{ledger}/transactions/{txid} | Get transaction from a ledger by its ID
*TransactionsApi* | [**ListTransactions**](docs/TransactionsApi.md#listtransactions) | **Get** /api/ledger/{ledger}/transactions | List transactions from a ledger
*TransactionsApi* | [**RevertTransaction**](docs/TransactionsApi.md#reverttransaction) | **Post** /api/ledger/{ledger}/transactions/{txid}/revert | Revert a ledger transaction by its ID
*UsersApi* | [**ListUsers**](docs/UsersApi.md#listusers) | **Get** /api/auth/users | List users
*UsersApi* | [**ReadUser**](docs/UsersApi.md#readuser) | **Get** /api/auth/users/{userId} | Read user
*WalletsApi* | [**ConfirmHold**](docs/WalletsApi.md#confirmhold) | **Post** /api/wallets/holds/{hold_id}/confirm | Confirm a hold
*WalletsApi* | [**CreateBalance**](docs/WalletsApi.md#createbalance) | **Post** /api/wallets/wallets/{id}/balances | Create a balance
*WalletsApi* | [**CreateWallet**](docs/WalletsApi.md#createwallet) | **Post** /api/wallets/wallets | Create a new wallet
*WalletsApi* | [**CreditWallet**](docs/WalletsApi.md#creditwallet) | **Post** /api/wallets/wallets/{id}/credit | Credit a wallet
*WalletsApi* | [**DebitWallet**](docs/WalletsApi.md#debitwallet) | **Post** /api/wallets/wallets/{id}/debit | Debit a wallet
*WalletsApi* | [**GetBalance**](docs/WalletsApi.md#getbalance) | **Get** /api/wallets/wallets/{id}/balances/{balanceName} | Get detailed balance
*WalletsApi* | [**GetHold**](docs/WalletsApi.md#gethold) | **Get** /api/wallets/holds/{holdID} | Get a hold
*WalletsApi* | [**GetHolds**](docs/WalletsApi.md#getholds) | **Get** /api/wallets/holds | Get all holds for a wallet
*WalletsApi* | [**GetTransactions**](docs/WalletsApi.md#gettransactions) | **Get** /api/wallets/transactions | 
*WalletsApi* | [**GetWallet**](docs/WalletsApi.md#getwallet) | **Get** /api/wallets/wallets/{id} | Get a wallet
*WalletsApi* | [**ListBalances**](docs/WalletsApi.md#listbalances) | **Get** /api/wallets/wallets/{id}/balances | List balances of a wallet
*WalletsApi* | [**ListWallets**](docs/WalletsApi.md#listwallets) | **Get** /api/wallets/wallets | List all wallets
*WalletsApi* | [**UpdateWallet**](docs/WalletsApi.md#updatewallet) | **Patch** /api/wallets/wallets/{id} | Update a wallet
*WalletsApi* | [**VoidHold**](docs/WalletsApi.md#voidhold) | **Post** /api/wallets/holds/{hold_id}/void | Cancel a hold
*WalletsApi* | [**WalletsgetServerInfo**](docs/WalletsApi.md#walletsgetserverinfo) | **Get** /api/wallets/_info | Get server info
*WebhooksApi* | [**ActivateConfig**](docs/WebhooksApi.md#activateconfig) | **Put** /api/webhooks/configs/{id}/activate | Activate one config
*WebhooksApi* | [**ChangeConfigSecret**](docs/WebhooksApi.md#changeconfigsecret) | **Put** /api/webhooks/configs/{id}/secret/change | Change the signing secret of a config
*WebhooksApi* | [**DeactivateConfig**](docs/WebhooksApi.md#deactivateconfig) | **Put** /api/webhooks/configs/{id}/deactivate | Deactivate one config
*WebhooksApi* | [**DeleteConfig**](docs/WebhooksApi.md#deleteconfig) | **Delete** /api/webhooks/configs/{id} | Delete one config
*WebhooksApi* | [**GetManyConfigs**](docs/WebhooksApi.md#getmanyconfigs) | **Get** /api/webhooks/configs | Get many configs
*WebhooksApi* | [**InsertConfig**](docs/WebhooksApi.md#insertconfig) | **Post** /api/webhooks/configs | Insert a new config
*WebhooksApi* | [**TestConfig**](docs/WebhooksApi.md#testconfig) | **Get** /api/webhooks/configs/{id}/test | Test one config


## Documentation For Models

 - [Account](docs/Account.md)
 - [AccountResponse](docs/AccountResponse.md)
 - [AccountWithVolumesAndBalances](docs/AccountWithVolumesAndBalances.md)
 - [AccountsCursor](docs/AccountsCursor.md)
 - [AccountsCursorCursor](docs/AccountsCursorCursor.md)
 - [AccountsCursorCursorAllOf](docs/AccountsCursorCursorAllOf.md)
 - [AccountsCursorResponse](docs/AccountsCursorResponse.md)
 - [AccountsCursorResponseCursor](docs/AccountsCursorResponseCursor.md)
 - [ActivityConfirmHold](docs/ActivityConfirmHold.md)
 - [ActivityCreateTransaction](docs/ActivityCreateTransaction.md)
 - [ActivityCreditWallet](docs/ActivityCreditWallet.md)
 - [ActivityDebitWallet](docs/ActivityDebitWallet.md)
 - [ActivityGetAccount](docs/ActivityGetAccount.md)
 - [ActivityGetPayment](docs/ActivityGetPayment.md)
 - [ActivityGetWallet](docs/ActivityGetWallet.md)
 - [ActivityRevertTransaction](docs/ActivityRevertTransaction.md)
 - [ActivityVoidHold](docs/ActivityVoidHold.md)
 - [AggregateBalancesResponse](docs/AggregateBalancesResponse.md)
 - [AssetHolder](docs/AssetHolder.md)
 - [Attempt](docs/Attempt.md)
 - [AttemptResponse](docs/AttemptResponse.md)
 - [Balance](docs/Balance.md)
 - [BalanceWithAssets](docs/BalanceWithAssets.md)
 - [BalancesCursorResponse](docs/BalancesCursorResponse.md)
 - [BalancesCursorResponseCursor](docs/BalancesCursorResponseCursor.md)
 - [BankingCircleConfig](docs/BankingCircleConfig.md)
 - [Client](docs/Client.md)
 - [ClientAllOf](docs/ClientAllOf.md)
 - [ClientOptions](docs/ClientOptions.md)
 - [ClientSecret](docs/ClientSecret.md)
 - [Config](docs/Config.md)
 - [ConfigChangeSecret](docs/ConfigChangeSecret.md)
 - [ConfigInfo](docs/ConfigInfo.md)
 - [ConfigInfoResponse](docs/ConfigInfoResponse.md)
 - [ConfigResponse](docs/ConfigResponse.md)
 - [ConfigUser](docs/ConfigUser.md)
 - [ConfigsResponse](docs/ConfigsResponse.md)
 - [ConfigsResponseCursor](docs/ConfigsResponseCursor.md)
 - [ConfigsResponseCursorAllOf](docs/ConfigsResponseCursorAllOf.md)
 - [ConfirmHoldRequest](docs/ConfirmHoldRequest.md)
 - [Connector](docs/Connector.md)
 - [ConnectorConfig](docs/ConnectorConfig.md)
 - [ConnectorConfigResponse](docs/ConnectorConfigResponse.md)
 - [ConnectorsConfigsResponse](docs/ConnectorsConfigsResponse.md)
 - [ConnectorsConfigsResponseData](docs/ConnectorsConfigsResponseData.md)
 - [ConnectorsConfigsResponseDataConnector](docs/ConnectorsConfigsResponseDataConnector.md)
 - [ConnectorsConfigsResponseDataConnectorKey](docs/ConnectorsConfigsResponseDataConnectorKey.md)
 - [ConnectorsResponse](docs/ConnectorsResponse.md)
 - [ConnectorsResponseDataInner](docs/ConnectorsResponseDataInner.md)
 - [Contract](docs/Contract.md)
 - [CreateBalanceResponse](docs/CreateBalanceResponse.md)
 - [CreateClientResponse](docs/CreateClientResponse.md)
 - [CreateScopeResponse](docs/CreateScopeResponse.md)
 - [CreateSecretResponse](docs/CreateSecretResponse.md)
 - [CreateWalletRequest](docs/CreateWalletRequest.md)
 - [CreateWalletResponse](docs/CreateWalletResponse.md)
 - [CreateWorkflowResponse](docs/CreateWorkflowResponse.md)
 - [CreditWalletRequest](docs/CreditWalletRequest.md)
 - [CurrencyCloudConfig](docs/CurrencyCloudConfig.md)
 - [Cursor](docs/Cursor.md)
 - [CursorBase](docs/CursorBase.md)
 - [DebitWalletRequest](docs/DebitWalletRequest.md)
 - [DebitWalletResponse](docs/DebitWalletResponse.md)
 - [DummyPayConfig](docs/DummyPayConfig.md)
 - [Error](docs/Error.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [ErrorsEnum](docs/ErrorsEnum.md)
 - [ExpandedDebitHold](docs/ExpandedDebitHold.md)
 - [ExpandedDebitHoldAllOf](docs/ExpandedDebitHoldAllOf.md)
 - [GetBalanceResponse](docs/GetBalanceResponse.md)
 - [GetHoldResponse](docs/GetHoldResponse.md)
 - [GetHoldsResponse](docs/GetHoldsResponse.md)
 - [GetHoldsResponseCursor](docs/GetHoldsResponseCursor.md)
 - [GetHoldsResponseCursorAllOf](docs/GetHoldsResponseCursorAllOf.md)
 - [GetTransactionsResponse](docs/GetTransactionsResponse.md)
 - [GetTransactionsResponseCursor](docs/GetTransactionsResponseCursor.md)
 - [GetTransactionsResponseCursorAllOf](docs/GetTransactionsResponseCursorAllOf.md)
 - [GetWalletResponse](docs/GetWalletResponse.md)
 - [GetWorkflowInstanceHistoryResponse](docs/GetWorkflowInstanceHistoryResponse.md)
 - [GetWorkflowInstanceHistoryStageResponse](docs/GetWorkflowInstanceHistoryStageResponse.md)
 - [GetWorkflowInstanceResponse](docs/GetWorkflowInstanceResponse.md)
 - [GetWorkflowResponse](docs/GetWorkflowResponse.md)
 - [Hold](docs/Hold.md)
 - [LedgerAccountSubject](docs/LedgerAccountSubject.md)
 - [LedgerInfo](docs/LedgerInfo.md)
 - [LedgerInfoResponse](docs/LedgerInfoResponse.md)
 - [LedgerInfoStorage](docs/LedgerInfoStorage.md)
 - [LedgerStorage](docs/LedgerStorage.md)
 - [ListBalancesResponse](docs/ListBalancesResponse.md)
 - [ListBalancesResponseCursor](docs/ListBalancesResponseCursor.md)
 - [ListBalancesResponseCursorAllOf](docs/ListBalancesResponseCursorAllOf.md)
 - [ListClientsResponse](docs/ListClientsResponse.md)
 - [ListRunsResponse](docs/ListRunsResponse.md)
 - [ListScopesResponse](docs/ListScopesResponse.md)
 - [ListUsersResponse](docs/ListUsersResponse.md)
 - [ListWalletsResponse](docs/ListWalletsResponse.md)
 - [ListWalletsResponseCursor](docs/ListWalletsResponseCursor.md)
 - [ListWalletsResponseCursorAllOf](docs/ListWalletsResponseCursorAllOf.md)
 - [ListWorkflowsResponse](docs/ListWorkflowsResponse.md)
 - [Log](docs/Log.md)
 - [LogsCursorResponse](docs/LogsCursorResponse.md)
 - [LogsCursorResponseCursor](docs/LogsCursorResponseCursor.md)
 - [Mapping](docs/Mapping.md)
 - [MappingResponse](docs/MappingResponse.md)
 - [MigrationInfo](docs/MigrationInfo.md)
 - [ModulrConfig](docs/ModulrConfig.md)
 - [Monetary](docs/Monetary.md)
 - [Payment](docs/Payment.md)
 - [PaymentAdjustment](docs/PaymentAdjustment.md)
 - [PaymentMetadata](docs/PaymentMetadata.md)
 - [PaymentResponse](docs/PaymentResponse.md)
 - [PaymentStatus](docs/PaymentStatus.md)
 - [PaymentsAccount](docs/PaymentsAccount.md)
 - [PaymentsCursor](docs/PaymentsCursor.md)
 - [PaymentsCursorCursor](docs/PaymentsCursorCursor.md)
 - [PaymentsCursorCursorAllOf](docs/PaymentsCursorCursorAllOf.md)
 - [PostTransaction](docs/PostTransaction.md)
 - [PostTransactionScript](docs/PostTransactionScript.md)
 - [Posting](docs/Posting.md)
 - [Query](docs/Query.md)
 - [ReadClientResponse](docs/ReadClientResponse.md)
 - [ReadUserResponse](docs/ReadUserResponse.md)
 - [ReadWorkflowResponse](docs/ReadWorkflowResponse.md)
 - [Response](docs/Response.md)
 - [RunWorkflowResponse](docs/RunWorkflowResponse.md)
 - [Scope](docs/Scope.md)
 - [ScopeAllOf](docs/ScopeAllOf.md)
 - [ScopeOptions](docs/ScopeOptions.md)
 - [Script](docs/Script.md)
 - [ScriptResponse](docs/ScriptResponse.md)
 - [Secret](docs/Secret.md)
 - [SecretAllOf](docs/SecretAllOf.md)
 - [SecretOptions](docs/SecretOptions.md)
 - [SendEventRequest](docs/SendEventRequest.md)
 - [ServerInfo](docs/ServerInfo.md)
 - [Stage](docs/Stage.md)
 - [StageDelay](docs/StageDelay.md)
 - [StageSend](docs/StageSend.md)
 - [StageSendDestination](docs/StageSendDestination.md)
 - [StageSendDestinationPayment](docs/StageSendDestinationPayment.md)
 - [StageSendSource](docs/StageSendSource.md)
 - [StageSendSourceAccount](docs/StageSendSourceAccount.md)
 - [StageSendSourcePayment](docs/StageSendSourcePayment.md)
 - [StageSendSourceWallet](docs/StageSendSourceWallet.md)
 - [StageStatus](docs/StageStatus.md)
 - [StageWaitEvent](docs/StageWaitEvent.md)
 - [Stats](docs/Stats.md)
 - [StatsResponse](docs/StatsResponse.md)
 - [StripeConfig](docs/StripeConfig.md)
 - [StripeTransferRequest](docs/StripeTransferRequest.md)
 - [Subject](docs/Subject.md)
 - [TaskBankingCircle](docs/TaskBankingCircle.md)
 - [TaskBankingCircleAllOf](docs/TaskBankingCircleAllOf.md)
 - [TaskBankingCircleAllOfDescriptor](docs/TaskBankingCircleAllOfDescriptor.md)
 - [TaskBase](docs/TaskBase.md)
 - [TaskCurrencyCloud](docs/TaskCurrencyCloud.md)
 - [TaskCurrencyCloudAllOf](docs/TaskCurrencyCloudAllOf.md)
 - [TaskCurrencyCloudAllOfDescriptor](docs/TaskCurrencyCloudAllOfDescriptor.md)
 - [TaskDummyPay](docs/TaskDummyPay.md)
 - [TaskDummyPayAllOf](docs/TaskDummyPayAllOf.md)
 - [TaskDummyPayAllOfDescriptor](docs/TaskDummyPayAllOfDescriptor.md)
 - [TaskModulr](docs/TaskModulr.md)
 - [TaskModulrAllOf](docs/TaskModulrAllOf.md)
 - [TaskModulrAllOfDescriptor](docs/TaskModulrAllOfDescriptor.md)
 - [TaskResponse](docs/TaskResponse.md)
 - [TaskStripe](docs/TaskStripe.md)
 - [TaskStripeAllOf](docs/TaskStripeAllOf.md)
 - [TaskStripeAllOfDescriptor](docs/TaskStripeAllOfDescriptor.md)
 - [TaskWise](docs/TaskWise.md)
 - [TaskWiseAllOf](docs/TaskWiseAllOf.md)
 - [TaskWiseAllOfDescriptor](docs/TaskWiseAllOfDescriptor.md)
 - [TasksCursor](docs/TasksCursor.md)
 - [TasksCursorCursor](docs/TasksCursorCursor.md)
 - [TasksCursorCursorAllOf](docs/TasksCursorCursorAllOf.md)
 - [TasksCursorCursorAllOfDataInner](docs/TasksCursorCursorAllOfDataInner.md)
 - [Total](docs/Total.md)
 - [Transaction](docs/Transaction.md)
 - [TransactionData](docs/TransactionData.md)
 - [TransactionResponse](docs/TransactionResponse.md)
 - [Transactions](docs/Transactions.md)
 - [TransactionsCursorResponse](docs/TransactionsCursorResponse.md)
 - [TransactionsCursorResponseCursor](docs/TransactionsCursorResponseCursor.md)
 - [TransactionsResponse](docs/TransactionsResponse.md)
 - [TransferRequest](docs/TransferRequest.md)
 - [TransferResponse](docs/TransferResponse.md)
 - [TransfersResponse](docs/TransfersResponse.md)
 - [TransfersResponseDataInner](docs/TransfersResponseDataInner.md)
 - [UpdateWalletRequest](docs/UpdateWalletRequest.md)
 - [User](docs/User.md)
 - [Volume](docs/Volume.md)
 - [Wallet](docs/Wallet.md)
 - [WalletSubject](docs/WalletSubject.md)
 - [WalletWithBalances](docs/WalletWithBalances.md)
 - [WalletWithBalancesBalances](docs/WalletWithBalancesBalances.md)
 - [WalletsCursor](docs/WalletsCursor.md)
 - [WalletsErrorResponse](docs/WalletsErrorResponse.md)
 - [WalletsTransaction](docs/WalletsTransaction.md)
 - [WalletsVolume](docs/WalletsVolume.md)
 - [WebhooksConfig](docs/WebhooksConfig.md)
 - [WiseConfig](docs/WiseConfig.md)
 - [Workflow](docs/Workflow.md)
 - [WorkflowConfig](docs/WorkflowConfig.md)
 - [WorkflowInstance](docs/WorkflowInstance.md)
 - [WorkflowInstanceHistory](docs/WorkflowInstanceHistory.md)
 - [WorkflowInstanceHistoryStage](docs/WorkflowInstanceHistoryStage.md)
 - [WorkflowInstanceHistoryStageInput](docs/WorkflowInstanceHistoryStageInput.md)
 - [WorkflowInstanceHistoryStageOutput](docs/WorkflowInstanceHistoryStageOutput.md)


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

