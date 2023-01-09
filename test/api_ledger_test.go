/*
Formance Stack API

Testing LedgerApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package formance

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    client "./openapi"
)

func Test_formance_LedgerApiService(t *testing.T) {

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)

    t.Run("Test LedgerApiService GetLedgerInfo", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ledger string

        resp, httpRes, err := apiClient.LedgerApi.GetLedgerInfo(context.Background(), ledger).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
