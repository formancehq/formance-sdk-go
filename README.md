# github.com/formancehq/formance-sdk-go

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/formancehq/formance-sdk-go
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	"github.com/formancehq/formance-sdk-go"
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "",
        }),
    )

    ctx := context.Background()
    res, err := s.GetVersions(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetVersionsResponse != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations

### [Formance SDK](docs/sdks/formance/README.md)

* [GetVersions](docs/sdks/formance/README.md#getversions) - Show stack version information

### [Auth](docs/sdks/auth/README.md)

* [AddScopeToClient](docs/sdks/auth/README.md#addscopetoclient) - Add scope to client
* [AddTransientScope](docs/sdks/auth/README.md#addtransientscope) - Add a transient scope to a scope
* [CreateClient](docs/sdks/auth/README.md#createclient) - Create client
* [CreateScope](docs/sdks/auth/README.md#createscope) - Create scope
* [CreateSecret](docs/sdks/auth/README.md#createsecret) - Add a secret to a client
* [DeleteClient](docs/sdks/auth/README.md#deleteclient) - Delete client
* [DeleteScope](docs/sdks/auth/README.md#deletescope) - Delete scope
* [DeleteScopeFromClient](docs/sdks/auth/README.md#deletescopefromclient) - Delete scope from client
* [DeleteSecret](docs/sdks/auth/README.md#deletesecret) - Delete a secret from a client
* [DeleteTransientScope](docs/sdks/auth/README.md#deletetransientscope) - Delete a transient scope from a scope
* [GetServerInfo](docs/sdks/auth/README.md#getserverinfo) - Get server info
* [ListClients](docs/sdks/auth/README.md#listclients) - List clients
* [ListScopes](docs/sdks/auth/README.md#listscopes) - List scopes
* [ListUsers](docs/sdks/auth/README.md#listusers) - List users
* [ReadClient](docs/sdks/auth/README.md#readclient) - Read client
* [ReadScope](docs/sdks/auth/README.md#readscope) - Read scope
* [ReadUser](docs/sdks/auth/README.md#readuser) - Read user
* [UpdateClient](docs/sdks/auth/README.md#updateclient) - Update client
* [UpdateScope](docs/sdks/auth/README.md#updatescope) - Update scope

### [Flows](docs/sdks/flows/README.md)

* [FlowsgetServerInfo](docs/sdks/flows/README.md#flowsgetserverinfo) - Get server info
* [CancelEvent](docs/sdks/flows/README.md#cancelevent) - Cancel a running workflow
* [CreateWorkflow](docs/sdks/flows/README.md#createworkflow) - Create workflow
* [GetInstance](docs/sdks/flows/README.md#getinstance) - Get a workflow instance by id
* [GetInstanceHistory](docs/sdks/flows/README.md#getinstancehistory) - Get a workflow instance history by id
* [GetInstanceStageHistory](docs/sdks/flows/README.md#getinstancestagehistory) - Get a workflow instance stage history
* [GetWorkflow](docs/sdks/flows/README.md#getworkflow) - Get a flow by id
* [ListInstances](docs/sdks/flows/README.md#listinstances) - List instances of a workflow
* [ListWorkflows](docs/sdks/flows/README.md#listworkflows) - List registered workflows
* [RunWorkflow](docs/sdks/flows/README.md#runworkflow) - Run workflow
* [SendEvent](docs/sdks/flows/README.md#sendevent) - Send an event to a running workflow

### [Ledger](docs/sdks/ledger/README.md)

* [CreateTransactions](docs/sdks/ledger/README.md#createtransactions) - Create a new batch of transactions to a ledger
* [AddMetadataOnTransaction](docs/sdks/ledger/README.md#addmetadataontransaction) - Set the metadata of a transaction by its ID
* [AddMetadataToAccount](docs/sdks/ledger/README.md#addmetadatatoaccount) - Add metadata to an account
* [CountAccounts](docs/sdks/ledger/README.md#countaccounts) - Count the accounts from a ledger
* [CountTransactions](docs/sdks/ledger/README.md#counttransactions) - Count the transactions from a ledger
* [CreateTransaction](docs/sdks/ledger/README.md#createtransaction) - Create a new transaction to a ledger
* [GetAccount](docs/sdks/ledger/README.md#getaccount) - Get account by its address
* [GetBalances](docs/sdks/ledger/README.md#getbalances) - Get the balances from a ledger's account
* [GetBalancesAggregated](docs/sdks/ledger/README.md#getbalancesaggregated) - Get the aggregated balances from selected accounts
* [GetInfo](docs/sdks/ledger/README.md#getinfo) - Show server information
* [GetLedgerInfo](docs/sdks/ledger/README.md#getledgerinfo) - Get information about a ledger
* [GetMapping](docs/sdks/ledger/README.md#getmapping) - Get the mapping of a ledger
* [GetTransaction](docs/sdks/ledger/README.md#gettransaction) - Get transaction from a ledger by its ID
* [ListAccounts](docs/sdks/ledger/README.md#listaccounts) - List accounts from a ledger
* [ListLogs](docs/sdks/ledger/README.md#listlogs) - List the logs from a ledger
* [ListTransactions](docs/sdks/ledger/README.md#listtransactions) - List transactions from a ledger
* [ReadStats](docs/sdks/ledger/README.md#readstats) - Get statistics from a ledger
* [RevertTransaction](docs/sdks/ledger/README.md#reverttransaction) - Revert a ledger transaction by its ID
* [~~RunScript~~](docs/sdks/ledger/README.md#runscript) - Execute a Numscript :warning: **Deprecated**
* [UpdateMapping](docs/sdks/ledger/README.md#updatemapping) - Update the mapping of a ledger

### [Payments](docs/sdks/payments/README.md)

* [ConnectorsStripeTransfer](docs/sdks/payments/README.md#connectorsstripetransfer) - Transfer funds between Stripe accounts
* [ConnectorsTransfer](docs/sdks/payments/README.md#connectorstransfer) - Transfer funds between Connector accounts
* [GetConnectorTask](docs/sdks/payments/README.md#getconnectortask) - Read a specific task of the connector
* [GetPayment](docs/sdks/payments/README.md#getpayment) - Get a payment
* [InstallConnector](docs/sdks/payments/README.md#installconnector) - Install a connector
* [ListAllConnectors](docs/sdks/payments/README.md#listallconnectors) - List all installed connectors
* [ListConfigsAvailableConnectors](docs/sdks/payments/README.md#listconfigsavailableconnectors) - List the configs of each available connector
* [ListConnectorTasks](docs/sdks/payments/README.md#listconnectortasks) - List tasks from a connector
* [ListConnectorsTransfers](docs/sdks/payments/README.md#listconnectorstransfers) - List transfers and their statuses
* [ListPayments](docs/sdks/payments/README.md#listpayments) - List payments
* [PaymentsgetServerInfo](docs/sdks/payments/README.md#paymentsgetserverinfo) - Get server info
* [PaymentslistAccounts](docs/sdks/payments/README.md#paymentslistaccounts) - List accounts
* [ReadConnectorConfig](docs/sdks/payments/README.md#readconnectorconfig) - Read the config of a connector
* [ResetConnector](docs/sdks/payments/README.md#resetconnector) - Reset a connector
* [UninstallConnector](docs/sdks/payments/README.md#uninstallconnector) - Uninstall a connector
* [UpdateMetadata](docs/sdks/payments/README.md#updatemetadata) - Update metadata

### [Search](docs/sdks/search/README.md)

* [Search](docs/sdks/search/README.md#search) - Search
* [SearchgetServerInfo](docs/sdks/search/README.md#searchgetserverinfo) - Get server info

### [Wallets](docs/sdks/wallets/README.md)

* [ConfirmHold](docs/sdks/wallets/README.md#confirmhold) - Confirm a hold
* [CreateBalance](docs/sdks/wallets/README.md#createbalance) - Create a balance
* [CreateWallet](docs/sdks/wallets/README.md#createwallet) - Create a new wallet
* [CreditWallet](docs/sdks/wallets/README.md#creditwallet) - Credit a wallet
* [DebitWallet](docs/sdks/wallets/README.md#debitwallet) - Debit a wallet
* [GetBalance](docs/sdks/wallets/README.md#getbalance) - Get detailed balance
* [GetHold](docs/sdks/wallets/README.md#gethold) - Get a hold
* [GetHolds](docs/sdks/wallets/README.md#getholds) - Get all holds for a wallet
* [GetTransactions](docs/sdks/wallets/README.md#gettransactions)
* [GetWallet](docs/sdks/wallets/README.md#getwallet) - Get a wallet
* [GetWalletSummary](docs/sdks/wallets/README.md#getwalletsummary) - Get wallet summary
* [ListBalances](docs/sdks/wallets/README.md#listbalances) - List balances of a wallet
* [ListWallets](docs/sdks/wallets/README.md#listwallets) - List all wallets
* [UpdateWallet](docs/sdks/wallets/README.md#updatewallet) - Update a wallet
* [VoidHold](docs/sdks/wallets/README.md#voidhold) - Cancel a hold
* [WalletsgetServerInfo](docs/sdks/wallets/README.md#walletsgetserverinfo) - Get server info

### [Webhooks](docs/sdks/webhooks/README.md)

* [ActivateConfig](docs/sdks/webhooks/README.md#activateconfig) - Activate one config
* [ChangeConfigSecret](docs/sdks/webhooks/README.md#changeconfigsecret) - Change the signing secret of a config
* [DeactivateConfig](docs/sdks/webhooks/README.md#deactivateconfig) - Deactivate one config
* [DeleteConfig](docs/sdks/webhooks/README.md#deleteconfig) - Delete one config
* [GetManyConfigs](docs/sdks/webhooks/README.md#getmanyconfigs) - Get many configs
* [InsertConfig](docs/sdks/webhooks/README.md#insertconfig) - Insert a new config
* [TestConfig](docs/sdks/webhooks/README.md#testconfig) - Test one config
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
