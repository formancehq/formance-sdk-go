# github.com/formancehq/formance-sdk-go/v2

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


## 🏗 **Welcome to your new SDK!** 🏗

It has been generated successfully based on your OpenAPI spec. However, it is not yet ready for production use. Here are some next steps:
- [ ] 🛠 Make your SDK feel handcrafted by [customizing it](https://www.speakeasyapi.dev/docs/customize-sdks)
- [ ] ♻️ Refine your SDK quickly by iterating locally with the [Speakeasy CLI](https://github.com/speakeasy-api/speakeasy)
- [ ] 🎁 Publish your SDK to package managers by [configuring automatic publishing](https://www.speakeasyapi.dev/docs/productionize-sdks/publish-sdks)
- [ ] ✨ When ready to productionize, delete this section from the README

<!-- Start Summary [summary] -->
## Summary

Formance Stack API: Open, modular foundation for unique payments flows

# Introduction
This API is documented in **OpenAPI format**.

# Authentication
Formance Stack offers one forms of authentication:
  - OAuth2
OAuth2 - an open protocol to allow secure authorization in a simple
and standard method from web, mobile and desktop applications.
<SecurityDefinitions />
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [github.com/formancehq/formance-sdk-go/v2](#githubcomformancehqformance-sdk-gov2)
  * [🏗 **Welcome to your new SDK!** 🏗](#welcome-to-your-new-sdk)
* [Introduction](#introduction)
* [Authentication](#authentication)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
  * [Authentication](#authentication-1)
  * [Retries](#retries)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/formancehq/formance-sdk-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"log"
)

func main() {
	ctx := context.Background()

	s := formancesdkgo.New(
		formancesdkgo.WithSecurity(shared.Security{
			ClientID:     formancesdkgo.String("<YOUR_CLIENT_ID_HERE>"),
			ClientSecret: formancesdkgo.String("<YOUR_CLIENT_SECRET_HERE>"),
		}),
	)

	res, err := s.GetVersions(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.GetVersionsResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Auth](docs/sdks/auth/README.md)


#### [Auth.V1](docs/sdks/v1/README.md)

* [CreateClient](docs/sdks/v1/README.md#createclient) - Create client
* [CreateSecret](docs/sdks/v1/README.md#createsecret) - Add a secret to a client
* [DeleteClient](docs/sdks/v1/README.md#deleteclient) - Delete client
* [DeleteSecret](docs/sdks/v1/README.md#deletesecret) - Delete a secret from a client
* [GetOIDCWellKnowns](docs/sdks/v1/README.md#getoidcwellknowns) - Retrieve OpenID connect well-knowns.
* [GetServerInfo](docs/sdks/v1/README.md#getserverinfo) - Get server info
* [ListClients](docs/sdks/v1/README.md#listclients) - List clients
* [ListUsers](docs/sdks/v1/README.md#listusers) - List users
* [ReadClient](docs/sdks/v1/README.md#readclient) - Read client
* [ReadUser](docs/sdks/v1/README.md#readuser) - Read user
* [UpdateClient](docs/sdks/v1/README.md#updateclient) - Update client

### [Formance SDK](docs/sdks/formance/README.md)

* [GetVersions](docs/sdks/formance/README.md#getversions) - Show stack version information

### [Ledger](docs/sdks/ledger/README.md)


#### [Ledger.V1](docs/sdks/formancev1/README.md)

* [CreateTransactions](docs/sdks/formancev1/README.md#createtransactions) - Create a new batch of transactions to a ledger
* [AddMetadataOnTransaction](docs/sdks/formancev1/README.md#addmetadataontransaction) - Set the metadata of a transaction by its ID
* [AddMetadataToAccount](docs/sdks/formancev1/README.md#addmetadatatoaccount) - Add metadata to an account
* [CountAccounts](docs/sdks/formancev1/README.md#countaccounts) - Count the accounts from a ledger
* [CountTransactions](docs/sdks/formancev1/README.md#counttransactions) - Count the transactions from a ledger
* [CreateTransaction](docs/sdks/formancev1/README.md#createtransaction) - Create a new transaction to a ledger
* [GetAccount](docs/sdks/formancev1/README.md#getaccount) - Get account by its address
* [GetBalances](docs/sdks/formancev1/README.md#getbalances) - Get the balances from a ledger's account
* [GetBalancesAggregated](docs/sdks/formancev1/README.md#getbalancesaggregated) - Get the aggregated balances from selected accounts
* [GetInfo](docs/sdks/formancev1/README.md#getinfo) - Show server information
* [GetLedgerInfo](docs/sdks/formancev1/README.md#getledgerinfo) - Get information about a ledger
* [GetMapping](docs/sdks/formancev1/README.md#getmapping) - Get the mapping of a ledger
* [GetTransaction](docs/sdks/formancev1/README.md#gettransaction) - Get transaction from a ledger by its ID
* [ListAccounts](docs/sdks/formancev1/README.md#listaccounts) - List accounts from a ledger
* [ListLogs](docs/sdks/formancev1/README.md#listlogs) - List the logs from a ledger
* [ListTransactions](docs/sdks/formancev1/README.md#listtransactions) - List transactions from a ledger
* [ReadStats](docs/sdks/formancev1/README.md#readstats) - Get statistics from a ledger
* [RevertTransaction](docs/sdks/formancev1/README.md#reverttransaction) - Revert a ledger transaction by its ID
* [~~RunScript~~](docs/sdks/formancev1/README.md#runscript) - Execute a Numscript :warning: **Deprecated**
* [UpdateMapping](docs/sdks/formancev1/README.md#updatemapping) - Update the mapping of a ledger

#### [Ledger.V2](docs/sdks/v2/README.md)

* [AddMetadataOnTransaction](docs/sdks/v2/README.md#addmetadataontransaction) - Set the metadata of a transaction by its ID
* [AddMetadataToAccount](docs/sdks/v2/README.md#addmetadatatoaccount) - Add metadata to an account
* [CountAccounts](docs/sdks/v2/README.md#countaccounts) - Count the accounts from a ledger
* [CountTransactions](docs/sdks/v2/README.md#counttransactions) - Count the transactions from a ledger
* [CreateBulk](docs/sdks/v2/README.md#createbulk) - Bulk request
* [CreateLedger](docs/sdks/v2/README.md#createledger) - Create a ledger
* [CreateTransaction](docs/sdks/v2/README.md#createtransaction) - Create a new transaction to a ledger
* [DeleteAccountMetadata](docs/sdks/v2/README.md#deleteaccountmetadata) - Delete metadata by key
* [DeleteLedgerMetadata](docs/sdks/v2/README.md#deleteledgermetadata) - Delete ledger metadata by key
* [DeleteTransactionMetadata](docs/sdks/v2/README.md#deletetransactionmetadata) - Delete metadata by key
* [ExportLogs](docs/sdks/v2/README.md#exportlogs) - Export logs
* [GetAccount](docs/sdks/v2/README.md#getaccount) - Get account by its address
* [GetBalancesAggregated](docs/sdks/v2/README.md#getbalancesaggregated) - Get the aggregated balances from selected accounts
* [GetInfo](docs/sdks/v2/README.md#getinfo) - Show server information
* [GetLedger](docs/sdks/v2/README.md#getledger) - Get a ledger
* [GetLedgerInfo](docs/sdks/v2/README.md#getledgerinfo) - Get information about a ledger
* [GetMetrics](docs/sdks/v2/README.md#getmetrics) - Read in memory metrics
* [GetTransaction](docs/sdks/v2/README.md#gettransaction) - Get transaction from a ledger by its ID
* [GetVolumesWithBalances](docs/sdks/v2/README.md#getvolumeswithbalances) - Get list of volumes with balances for (account/asset)
* [ImportLogs](docs/sdks/v2/README.md#importlogs)
* [ListAccounts](docs/sdks/v2/README.md#listaccounts) - List accounts from a ledger
* [ListLedgers](docs/sdks/v2/README.md#listledgers) - List ledgers
* [ListLogs](docs/sdks/v2/README.md#listlogs) - List the logs from a ledger
* [ListTransactions](docs/sdks/v2/README.md#listtransactions) - List transactions from a ledger
* [ReadStats](docs/sdks/v2/README.md#readstats) - Get statistics from a ledger
* [RevertTransaction](docs/sdks/v2/README.md#reverttransaction) - Revert a ledger transaction by its ID
* [UpdateLedgerMetadata](docs/sdks/v2/README.md#updateledgermetadata) - Update ledger metadata

### [Orchestration](docs/sdks/orchestration/README.md)


#### [Orchestration.V1](docs/sdks/formanceorchestrationv1/README.md)

* [CancelEvent](docs/sdks/formanceorchestrationv1/README.md#cancelevent) - Cancel a running workflow
* [CreateTrigger](docs/sdks/formanceorchestrationv1/README.md#createtrigger) - Create trigger
* [CreateWorkflow](docs/sdks/formanceorchestrationv1/README.md#createworkflow) - Create workflow
* [DeleteTrigger](docs/sdks/formanceorchestrationv1/README.md#deletetrigger) - Delete trigger
* [DeleteWorkflow](docs/sdks/formanceorchestrationv1/README.md#deleteworkflow) - Delete a flow by id
* [GetInstance](docs/sdks/formanceorchestrationv1/README.md#getinstance) - Get a workflow instance by id
* [GetInstanceHistory](docs/sdks/formanceorchestrationv1/README.md#getinstancehistory) - Get a workflow instance history by id
* [GetInstanceStageHistory](docs/sdks/formanceorchestrationv1/README.md#getinstancestagehistory) - Get a workflow instance stage history
* [GetWorkflow](docs/sdks/formanceorchestrationv1/README.md#getworkflow) - Get a flow by id
* [ListInstances](docs/sdks/formanceorchestrationv1/README.md#listinstances) - List instances of a workflow
* [ListTriggers](docs/sdks/formanceorchestrationv1/README.md#listtriggers) - List triggers
* [ListTriggersOccurrences](docs/sdks/formanceorchestrationv1/README.md#listtriggersoccurrences) - List triggers occurrences
* [ListWorkflows](docs/sdks/formanceorchestrationv1/README.md#listworkflows) - List registered workflows
* [OrchestrationgetServerInfo](docs/sdks/formanceorchestrationv1/README.md#orchestrationgetserverinfo) - Get server info
* [ReadTrigger](docs/sdks/formanceorchestrationv1/README.md#readtrigger) - Read trigger
* [RunWorkflow](docs/sdks/formanceorchestrationv1/README.md#runworkflow) - Run workflow
* [SendEvent](docs/sdks/formanceorchestrationv1/README.md#sendevent) - Send an event to a running workflow

#### [Orchestration.V2](docs/sdks/formancev2/README.md)

* [CancelEvent](docs/sdks/formancev2/README.md#cancelevent) - Cancel a running workflow
* [CreateTrigger](docs/sdks/formancev2/README.md#createtrigger) - Create trigger
* [CreateWorkflow](docs/sdks/formancev2/README.md#createworkflow) - Create workflow
* [DeleteTrigger](docs/sdks/formancev2/README.md#deletetrigger) - Delete trigger
* [DeleteWorkflow](docs/sdks/formancev2/README.md#deleteworkflow) - Delete a flow by id
* [GetInstance](docs/sdks/formancev2/README.md#getinstance) - Get a workflow instance by id
* [GetInstanceHistory](docs/sdks/formancev2/README.md#getinstancehistory) - Get a workflow instance history by id
* [GetInstanceStageHistory](docs/sdks/formancev2/README.md#getinstancestagehistory) - Get a workflow instance stage history
* [GetServerInfo](docs/sdks/formancev2/README.md#getserverinfo) - Get server info
* [GetWorkflow](docs/sdks/formancev2/README.md#getworkflow) - Get a flow by id
* [ListInstances](docs/sdks/formancev2/README.md#listinstances) - List instances of a workflow
* [ListTriggers](docs/sdks/formancev2/README.md#listtriggers) - List triggers
* [ListTriggersOccurrences](docs/sdks/formancev2/README.md#listtriggersoccurrences) - List triggers occurrences
* [ListWorkflows](docs/sdks/formancev2/README.md#listworkflows) - List registered workflows
* [ReadTrigger](docs/sdks/formancev2/README.md#readtrigger) - Read trigger
* [RunWorkflow](docs/sdks/formancev2/README.md#runworkflow) - Run workflow
* [SendEvent](docs/sdks/formancev2/README.md#sendevent) - Send an event to a running workflow
* [TestTrigger](docs/sdks/formancev2/README.md#testtrigger) - Test trigger

### [Payments](docs/sdks/payments/README.md)


#### [Payments.V1](docs/sdks/formancepaymentsv1/README.md)

* [AddAccountToPool](docs/sdks/formancepaymentsv1/README.md#addaccounttopool) - Add an account to a pool
* [ConnectorsTransfer](docs/sdks/formancepaymentsv1/README.md#connectorstransfer) - Transfer funds between Connector accounts
* [CreateAccount](docs/sdks/formancepaymentsv1/README.md#createaccount) - Create an account
* [CreateBankAccount](docs/sdks/formancepaymentsv1/README.md#createbankaccount) - Create a BankAccount in Payments and on the PSP
* [CreatePayment](docs/sdks/formancepaymentsv1/README.md#createpayment) - Create a payment
* [CreatePool](docs/sdks/formancepaymentsv1/README.md#createpool) - Create a Pool
* [CreateTransferInitiation](docs/sdks/formancepaymentsv1/README.md#createtransferinitiation) - Create a TransferInitiation
* [DeletePool](docs/sdks/formancepaymentsv1/README.md#deletepool) - Delete a Pool
* [DeleteTransferInitiation](docs/sdks/formancepaymentsv1/README.md#deletetransferinitiation) - Delete a transfer initiation
* [ForwardBankAccount](docs/sdks/formancepaymentsv1/README.md#forwardbankaccount) - Forward a bank account to a connector
* [GetAccountBalances](docs/sdks/formancepaymentsv1/README.md#getaccountbalances) - Get account balances
* [GetBankAccount](docs/sdks/formancepaymentsv1/README.md#getbankaccount) - Get a bank account created by user on Formance
* [~~GetConnectorTask~~](docs/sdks/formancepaymentsv1/README.md#getconnectortask) - Read a specific task of the connector :warning: **Deprecated**
* [GetConnectorTaskV1](docs/sdks/formancepaymentsv1/README.md#getconnectortaskv1) - Read a specific task of the connector
* [GetPayment](docs/sdks/formancepaymentsv1/README.md#getpayment) - Get a payment
* [GetPool](docs/sdks/formancepaymentsv1/README.md#getpool) - Get a Pool
* [GetPoolBalances](docs/sdks/formancepaymentsv1/README.md#getpoolbalances) - Get pool balances
* [GetTransferInitiation](docs/sdks/formancepaymentsv1/README.md#gettransferinitiation) - Get a transfer initiation
* [InstallConnector](docs/sdks/formancepaymentsv1/README.md#installconnector) - Install a connector
* [ListAllConnectors](docs/sdks/formancepaymentsv1/README.md#listallconnectors) - List all installed connectors
* [ListBankAccounts](docs/sdks/formancepaymentsv1/README.md#listbankaccounts) - List bank accounts created by user on Formance
* [ListConfigsAvailableConnectors](docs/sdks/formancepaymentsv1/README.md#listconfigsavailableconnectors) - List the configs of each available connector
* [~~ListConnectorTasks~~](docs/sdks/formancepaymentsv1/README.md#listconnectortasks) - List tasks from a connector :warning: **Deprecated**
* [ListConnectorTasksV1](docs/sdks/formancepaymentsv1/README.md#listconnectortasksv1) - List tasks from a connector
* [ListPayments](docs/sdks/formancepaymentsv1/README.md#listpayments) - List payments
* [ListPools](docs/sdks/formancepaymentsv1/README.md#listpools) - List Pools
* [ListTransferInitiations](docs/sdks/formancepaymentsv1/README.md#listtransferinitiations) - List Transfer Initiations
* [PaymentsgetAccount](docs/sdks/formancepaymentsv1/README.md#paymentsgetaccount) - Get an account
* [PaymentsgetServerInfo](docs/sdks/formancepaymentsv1/README.md#paymentsgetserverinfo) - Get server info
* [PaymentslistAccounts](docs/sdks/formancepaymentsv1/README.md#paymentslistaccounts) - List accounts
* [~~ReadConnectorConfig~~](docs/sdks/formancepaymentsv1/README.md#readconnectorconfig) - Read the config of a connector :warning: **Deprecated**
* [ReadConnectorConfigV1](docs/sdks/formancepaymentsv1/README.md#readconnectorconfigv1) - Read the config of a connector
* [RemoveAccountFromPool](docs/sdks/formancepaymentsv1/README.md#removeaccountfrompool) - Remove an account from a pool
* [~~ResetConnector~~](docs/sdks/formancepaymentsv1/README.md#resetconnector) - Reset a connector :warning: **Deprecated**
* [ResetConnectorV1](docs/sdks/formancepaymentsv1/README.md#resetconnectorv1) - Reset a connector
* [RetryTransferInitiation](docs/sdks/formancepaymentsv1/README.md#retrytransferinitiation) - Retry a failed transfer initiation
* [ReverseTransferInitiation](docs/sdks/formancepaymentsv1/README.md#reversetransferinitiation) - Reverse a transfer initiation
* [UdpateTransferInitiationStatus](docs/sdks/formancepaymentsv1/README.md#udpatetransferinitiationstatus) - Update the status of a transfer initiation
* [~~UninstallConnector~~](docs/sdks/formancepaymentsv1/README.md#uninstallconnector) - Uninstall a connector :warning: **Deprecated**
* [UninstallConnectorV1](docs/sdks/formancepaymentsv1/README.md#uninstallconnectorv1) - Uninstall a connector
* [UpdateBankAccountMetadata](docs/sdks/formancepaymentsv1/README.md#updatebankaccountmetadata) - Update metadata of a bank account
* [UpdateConnectorConfigV1](docs/sdks/formancepaymentsv1/README.md#updateconnectorconfigv1) - Update the config of a connector
* [UpdateMetadata](docs/sdks/formancepaymentsv1/README.md#updatemetadata) - Update metadata

#### [Payments.V3](docs/sdks/v3/README.md)

* [AddAccountToPool](docs/sdks/v3/README.md#addaccounttopool) - Add an account to a pool
* [ApprovePaymentInitiation](docs/sdks/v3/README.md#approvepaymentinitiation) - Approve a payment initiation
* [CreateAccount](docs/sdks/v3/README.md#createaccount) - Create a formance account object. This object will not be forwarded to the connector. It is only used for internal purposes.

* [CreateBankAccount](docs/sdks/v3/README.md#createbankaccount) - Create a formance bank account object. This object will not be forwarded to the connector until you called the forwardBankAccount method.

* [CreatePayment](docs/sdks/v3/README.md#createpayment) - Create a formance payment object. This object will not be forwarded to the connector. It is only used for internal purposes.

* [CreatePool](docs/sdks/v3/README.md#createpool) - Create a formance pool object
* [DeletePaymentInitiation](docs/sdks/v3/README.md#deletepaymentinitiation) - Delete a payment initiation by ID
* [DeletePool](docs/sdks/v3/README.md#deletepool) - Delete a pool by ID
* [ForwardBankAccount](docs/sdks/v3/README.md#forwardbankaccount) - Forward a Bank Account to a PSP for creation
* [GetAccount](docs/sdks/v3/README.md#getaccount) - Get an account by ID
* [GetAccountBalances](docs/sdks/v3/README.md#getaccountbalances) - Get account balances
* [GetBankAccount](docs/sdks/v3/README.md#getbankaccount) - Get a Bank Account by ID
* [GetConnectorConfig](docs/sdks/v3/README.md#getconnectorconfig) - Get a connector configuration by ID
* [GetConnectorSchedule](docs/sdks/v3/README.md#getconnectorschedule) - Get a connector schedule by ID
* [GetPayment](docs/sdks/v3/README.md#getpayment) - Get a payment by ID
* [GetPaymentInitiation](docs/sdks/v3/README.md#getpaymentinitiation) - Get a payment initiation by ID
* [GetPool](docs/sdks/v3/README.md#getpool) - Get a pool by ID
* [GetPoolBalances](docs/sdks/v3/README.md#getpoolbalances) - Get pool balances
* [GetTask](docs/sdks/v3/README.md#gettask) - Get a task and its result by ID
* [InitiatePayment](docs/sdks/v3/README.md#initiatepayment) - Initiate a payment
* [InstallConnector](docs/sdks/v3/README.md#installconnector) - Install a connector
* [ListAccounts](docs/sdks/v3/README.md#listaccounts) - List all accounts
* [ListBankAccounts](docs/sdks/v3/README.md#listbankaccounts) - List all bank accounts
* [ListConnectorConfigs](docs/sdks/v3/README.md#listconnectorconfigs) - List all connector configurations
* [ListConnectorScheduleInstances](docs/sdks/v3/README.md#listconnectorscheduleinstances) - List all connector schedule instances
* [ListConnectorSchedules](docs/sdks/v3/README.md#listconnectorschedules) - List all connector schedules
* [ListConnectors](docs/sdks/v3/README.md#listconnectors) - List all connectors
* [ListPaymentInitiationAdjustments](docs/sdks/v3/README.md#listpaymentinitiationadjustments) - List all payment initiation adjustments
* [ListPaymentInitiationRelatedPayments](docs/sdks/v3/README.md#listpaymentinitiationrelatedpayments) - List all payments related to a payment initiation
* [ListPaymentInitiations](docs/sdks/v3/README.md#listpaymentinitiations) - List all payment initiations
* [ListPayments](docs/sdks/v3/README.md#listpayments) - List all payments
* [ListPools](docs/sdks/v3/README.md#listpools) - List all pools
* [RejectPaymentInitiation](docs/sdks/v3/README.md#rejectpaymentinitiation) - Reject a payment initiation
* [RemoveAccountFromPool](docs/sdks/v3/README.md#removeaccountfrompool) - Remove an account from a pool
* [ResetConnector](docs/sdks/v3/README.md#resetconnector) - Reset a connector. Be aware that this will delete all data and stop all existing tasks like payment initiations and bank account creations.
* [RetryPaymentInitiation](docs/sdks/v3/README.md#retrypaymentinitiation) - Retry a payment initiation
* [ReversePaymentInitiation](docs/sdks/v3/README.md#reversepaymentinitiation) - Reverse a payment initiation
* [UninstallConnector](docs/sdks/v3/README.md#uninstallconnector) - Uninstall a connector
* [UpdateBankAccountMetadata](docs/sdks/v3/README.md#updatebankaccountmetadata) - Update a bank account's metadata
* [UpdatePaymentMetadata](docs/sdks/v3/README.md#updatepaymentmetadata) - Update a payment's metadata

### [Reconciliation](docs/sdks/reconciliation/README.md)


#### [Reconciliation.V1](docs/sdks/formancereconciliationv1/README.md)

* [CreatePolicy](docs/sdks/formancereconciliationv1/README.md#createpolicy) - Create a policy
* [DeletePolicy](docs/sdks/formancereconciliationv1/README.md#deletepolicy) - Delete a policy
* [GetPolicy](docs/sdks/formancereconciliationv1/README.md#getpolicy) - Get a policy
* [GetReconciliation](docs/sdks/formancereconciliationv1/README.md#getreconciliation) - Get a reconciliation
* [ListPolicies](docs/sdks/formancereconciliationv1/README.md#listpolicies) - List policies
* [ListReconciliations](docs/sdks/formancereconciliationv1/README.md#listreconciliations) - List reconciliations
* [Reconcile](docs/sdks/formancereconciliationv1/README.md#reconcile) - Reconcile using a policy
* [ReconciliationgetServerInfo](docs/sdks/formancereconciliationv1/README.md#reconciliationgetserverinfo) - Get server info

### [~~Search~~](docs/sdks/search/README.md)


#### [~~Search.V1~~](docs/sdks/formancesearchv1/README.md)

* [~~Search~~](docs/sdks/formancesearchv1/README.md#search) - search.v1 :warning: **Deprecated**
* [~~SearchgetServerInfo~~](docs/sdks/formancesearchv1/README.md#searchgetserverinfo) - Get server info :warning: **Deprecated**

### [Wallets](docs/sdks/wallets/README.md)


#### [Wallets.V1](docs/sdks/formancewalletsv1/README.md)

* [ConfirmHold](docs/sdks/formancewalletsv1/README.md#confirmhold) - Confirm a hold
* [CreateBalance](docs/sdks/formancewalletsv1/README.md#createbalance) - Create a balance
* [CreateWallet](docs/sdks/formancewalletsv1/README.md#createwallet) - Create a new wallet
* [CreditWallet](docs/sdks/formancewalletsv1/README.md#creditwallet) - Credit a wallet
* [DebitWallet](docs/sdks/formancewalletsv1/README.md#debitwallet) - Debit a wallet
* [GetBalance](docs/sdks/formancewalletsv1/README.md#getbalance) - Get detailed balance
* [GetHold](docs/sdks/formancewalletsv1/README.md#gethold) - Get a hold
* [GetHolds](docs/sdks/formancewalletsv1/README.md#getholds) - Get all holds for a wallet
* [GetTransactions](docs/sdks/formancewalletsv1/README.md#gettransactions)
* [GetWallet](docs/sdks/formancewalletsv1/README.md#getwallet) - Get a wallet
* [GetWalletSummary](docs/sdks/formancewalletsv1/README.md#getwalletsummary) - Get wallet summary
* [ListBalances](docs/sdks/formancewalletsv1/README.md#listbalances) - List balances of a wallet
* [ListWallets](docs/sdks/formancewalletsv1/README.md#listwallets) - List all wallets
* [UpdateWallet](docs/sdks/formancewalletsv1/README.md#updatewallet) - Update a wallet
* [VoidHold](docs/sdks/formancewalletsv1/README.md#voidhold) - Cancel a hold
* [WalletsgetServerInfo](docs/sdks/formancewalletsv1/README.md#walletsgetserverinfo) - Get server info

### [Webhooks](docs/sdks/webhooks/README.md)


#### [Webhooks.V1](docs/sdks/formancewebhooksv1/README.md)

* [ActivateConfig](docs/sdks/formancewebhooksv1/README.md#activateconfig) - Activate one config
* [ChangeConfigSecret](docs/sdks/formancewebhooksv1/README.md#changeconfigsecret) - Change the signing secret of a config
* [DeactivateConfig](docs/sdks/formancewebhooksv1/README.md#deactivateconfig) - Deactivate one config
* [DeleteConfig](docs/sdks/formancewebhooksv1/README.md#deleteconfig) - Delete one config
* [GetManyConfigs](docs/sdks/formancewebhooksv1/README.md#getmanyconfigs) - Get many configs
* [InsertConfig](docs/sdks/formancewebhooksv1/README.md#insertconfig) - Insert a new config
* [TestConfig](docs/sdks/formancewebhooksv1/README.md#testconfig) - Test one config
* [UpdateConfig](docs/sdks/formancewebhooksv1/README.md#updateconfig) - Update one config

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `sdkerrors.SDKError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `AddMetadataOnTransaction` function may return the following errors:

| Error Type                | Status Code | Content Type     |
| ------------------------- | ----------- | ---------------- |
| sdkerrors.V2ErrorResponse | default     | application/json |
| sdkerrors.SDKError        | 4XX, 5XX    | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/operations"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/sdkerrors"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"log"
	"math/big"
)

func main() {
	ctx := context.Background()

	s := formancesdkgo.New(
		formancesdkgo.WithSecurity(shared.Security{
			ClientID:     formancesdkgo.String("<YOUR_CLIENT_ID_HERE>"),
			ClientSecret: formancesdkgo.String("<YOUR_CLIENT_SECRET_HERE>"),
		}),
	)

	res, err := s.Ledger.V2.AddMetadataOnTransaction(ctx, operations.V2AddMetadataOnTransactionRequest{
		RequestBody: map[string]string{
			"admin": "true",
		},
		DryRun: formancesdkgo.Bool(true),
		ID:     big.NewInt(1234),
		Ledger: "ledger001",
	})
	if err != nil {

		var e *sdkerrors.V2ErrorResponse
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex(serverIndex int)` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| #   | Server                                                | Variables                        | Description                                |
| --- | ----------------------------------------------------- | -------------------------------- | ------------------------------------------ |
| 0   | `http://localhost`                                    |                                  | local server                               |
| 1   | `https://{organization}.{environment}.formance.cloud` | `environment`<br/>`organization` | A per-organization and per-environment API |

If the selected server has variables, you may override its default values using the associated option(s):

| Variable       | Option                                           | Supported Values                                                           | Default           | Description                                                   |
| -------------- | ------------------------------------------------ | -------------------------------------------------------------------------- | ----------------- | ------------------------------------------------------------- |
| `environment`  | `WithEnvironment(environment ServerEnvironment)` | - `"eu.sandbox"`<br/>- `"sandbox"`<br/>- `"eu-west-1"`<br/>- `"us-east-1"` | `"eu.sandbox"`    | The environment name. Defaults to the production environment. |
| `organization` | `WithOrganization(organization string)`          | string                                                                     | `"orgID-stackID"` | The organization name. Defaults to a generic organization.    |

#### Example

```go
package main

import (
	"context"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"log"
)

func main() {
	ctx := context.Background()

	s := formancesdkgo.New(
		formancesdkgo.WithServerIndex(1),
		formancesdkgo.WithEnvironment("us-east-1"),
		formancesdkgo.WithOrganization("<value>"),
		formancesdkgo.WithSecurity(shared.Security{
			ClientID:     formancesdkgo.String("<YOUR_CLIENT_ID_HERE>"),
			ClientSecret: formancesdkgo.String("<YOUR_CLIENT_SECRET_HERE>"),
		}),
	)

	res, err := s.GetVersions(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.GetVersionsResponse != nil {
		// handle response
	}
}

```

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"log"
)

func main() {
	ctx := context.Background()

	s := formancesdkgo.New(
		formancesdkgo.WithServerURL("http://localhost"),
		formancesdkgo.WithSecurity(shared.Security{
			ClientID:     formancesdkgo.String("<YOUR_CLIENT_ID_HERE>"),
			ClientSecret: formancesdkgo.String("<YOUR_CLIENT_SECRET_HERE>"),
		}),
	)

	res, err := s.GetVersions(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.GetVersionsResponse != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name                          | Type   | Scheme                         |
| ----------------------------- | ------ | ------------------------------ |
| `ClientID`<br/>`ClientSecret` | oauth2 | OAuth2 Client Credentials Flow |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"log"
)

func main() {
	ctx := context.Background()

	s := formancesdkgo.New(
		formancesdkgo.WithSecurity(shared.Security{
			ClientID:     formancesdkgo.String("<YOUR_CLIENT_ID_HERE>"),
			ClientSecret: formancesdkgo.String("<YOUR_CLIENT_SECRET_HERE>"),
		}),
	)

	res, err := s.GetVersions(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.GetVersionsResponse != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v3/pkg/retry"
	"log"
	"pkg/models/operations"
)

func main() {
	ctx := context.Background()

	s := formancesdkgo.New(
		formancesdkgo.WithSecurity(shared.Security{
			ClientID:     formancesdkgo.String("<YOUR_CLIENT_ID_HERE>"),
			ClientSecret: formancesdkgo.String("<YOUR_CLIENT_SECRET_HERE>"),
		}),
	)

	res, err := s.GetVersions(ctx, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.GetVersionsResponse != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v3/pkg/retry"
	"log"
)

func main() {
	ctx := context.Background()

	s := formancesdkgo.New(
		formancesdkgo.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		formancesdkgo.WithSecurity(shared.Security{
			ClientID:     formancesdkgo.String("<YOUR_CLIENT_ID_HERE>"),
			ClientSecret: formancesdkgo.String("<YOUR_CLIENT_SECRET_HERE>"),
		}),
	)

	res, err := s.GetVersions(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.GetVersionsResponse != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
