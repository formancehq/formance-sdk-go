/*
Formance Stack API

Testing ScopesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package formance

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	client "github.com/formancehq/formance-sdk-go"
)

func Test_formance_ScopesApiService(t *testing.T) {

	configuration := client.NewConfiguration()
	apiClient := client.NewAPIClient(configuration)

	t.Run("Test ScopesApiService AddTransientScope", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var scopeId string
		var transientScopeId string

		httpRes, err := apiClient.ScopesApi.AddTransientScope(context.Background(), scopeId, transientScopeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopesApiService CreateScope", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ScopesApi.CreateScope(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopesApiService DeleteScope", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var scopeId string

		httpRes, err := apiClient.ScopesApi.DeleteScope(context.Background(), scopeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopesApiService DeleteTransientScope", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var scopeId string
		var transientScopeId string

		httpRes, err := apiClient.ScopesApi.DeleteTransientScope(context.Background(), scopeId, transientScopeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopesApiService ListScopes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ScopesApi.ListScopes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopesApiService ReadScope", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var scopeId string

		resp, httpRes, err := apiClient.ScopesApi.ReadScope(context.Background(), scopeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopesApiService UpdateScope", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var scopeId string

		resp, httpRes, err := apiClient.ScopesApi.UpdateScope(context.Background(), scopeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
