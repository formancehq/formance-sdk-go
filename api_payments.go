/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: v0.2.2
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


type PaymentsApi interface {

	/*
	GetAllConnectors Get all installed connectors

	Get all installed connectors

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetAllConnectorsRequest
	*/
	GetAllConnectors(ctx context.Context) ApiGetAllConnectorsRequest

	// GetAllConnectorsExecute executes the request
	//  @return ListConnectorsResponse
	GetAllConnectorsExecute(r ApiGetAllConnectorsRequest) (*ListConnectorsResponse, *http.Response, error)

	/*
	GetAllConnectorsConfigs Get all available connectors configs

	Get all available connectors configs

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetAllConnectorsConfigsRequest
	*/
	GetAllConnectorsConfigs(ctx context.Context) ApiGetAllConnectorsConfigsRequest

	// GetAllConnectorsConfigsExecute executes the request
	//  @return ListConnectorsConfigsResponse
	GetAllConnectorsConfigsExecute(r ApiGetAllConnectorsConfigsRequest) (*ListConnectorsConfigsResponse, *http.Response, error)

	/*
	GetConnectorTask Read a specific task of the connector

	Get a specific task associated to the connector

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param connector The connector code
	@param taskId The task id
	@return ApiGetConnectorTaskRequest
	*/
	GetConnectorTask(ctx context.Context, connector string, taskId string) ApiGetConnectorTaskRequest

	// GetConnectorTaskExecute executes the request
	//  @return ConnectorTask
	GetConnectorTaskExecute(r ApiGetConnectorTaskRequest) (*ConnectorTask, *http.Response, error)

	/*
	GetPayment Returns a payment.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param paymentId The payment id
	@return ApiGetPaymentRequest
	*/
	GetPayment(ctx context.Context, paymentId string) ApiGetPaymentRequest

	// GetPaymentExecute executes the request
	//  @return Payment
	GetPaymentExecute(r ApiGetPaymentRequest) (*Payment, *http.Response, error)

	/*
	InstallConnector Install connector

	Install connector

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param connector The connector code
	@return ApiInstallConnectorRequest
	*/
	InstallConnector(ctx context.Context, connector string) ApiInstallConnectorRequest

	// InstallConnectorExecute executes the request
	InstallConnectorExecute(r ApiInstallConnectorRequest) (*http.Response, error)

	/*
	ListConnectorTasks List connector tasks

	List all tasks associated with this connector.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param connector The connector code
	@return ApiListConnectorTasksRequest
	*/
	ListConnectorTasks(ctx context.Context, connector string) ApiListConnectorTasksRequest

	// ListConnectorTasksExecute executes the request
	//  @return []ConnectorTask
	ListConnectorTasksExecute(r ApiListConnectorTasksRequest) ([]ConnectorTask, *http.Response, error)

	/*
	ListPayments Returns a list of payments.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListPaymentsRequest
	*/
	ListPayments(ctx context.Context) ApiListPaymentsRequest

	// ListPaymentsExecute executes the request
	//  @return ListPaymentsResponse
	ListPaymentsExecute(r ApiListPaymentsRequest) (*ListPaymentsResponse, *http.Response, error)

	/*
	ReadConnectorConfig Read connector config

	Read connector config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param connector The connector code
	@return ApiReadConnectorConfigRequest
	*/
	ReadConnectorConfig(ctx context.Context, connector string) ApiReadConnectorConfigRequest

	// ReadConnectorConfigExecute executes the request
	//  @return ConnectorConfig
	ReadConnectorConfigExecute(r ApiReadConnectorConfigRequest) (*ConnectorConfig, *http.Response, error)

	/*
	ResetConnector Reset connector

	Reset connector. Will remove the connector and ALL PAYMENTS generated with it.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param connector The connector code
	@return ApiResetConnectorRequest
	*/
	ResetConnector(ctx context.Context, connector string) ApiResetConnectorRequest

	// ResetConnectorExecute executes the request
	ResetConnectorExecute(r ApiResetConnectorRequest) (*http.Response, error)

	/*
	UninstallConnector Uninstall connector

	Uninstall  connector

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param connector The connector code
	@return ApiUninstallConnectorRequest
	*/
	UninstallConnector(ctx context.Context, connector string) ApiUninstallConnectorRequest

	// UninstallConnectorExecute executes the request
	UninstallConnectorExecute(r ApiUninstallConnectorRequest) (*http.Response, error)
}

// PaymentsApiService PaymentsApi service
type PaymentsApiService service

type ApiGetAllConnectorsRequest struct {
	ctx context.Context
	ApiService PaymentsApi
}

func (r ApiGetAllConnectorsRequest) Execute() (*ListConnectorsResponse, *http.Response, error) {
	return r.ApiService.GetAllConnectorsExecute(r)
}

/*
GetAllConnectors Get all installed connectors

Get all installed connectors

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAllConnectorsRequest
*/
func (a *PaymentsApiService) GetAllConnectors(ctx context.Context) ApiGetAllConnectorsRequest {
	return ApiGetAllConnectorsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListConnectorsResponse
func (a *PaymentsApiService) GetAllConnectorsExecute(r ApiGetAllConnectorsRequest) (*ListConnectorsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListConnectorsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentsApiService.GetAllConnectors")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/payments/connectors"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetAllConnectorsConfigsRequest struct {
	ctx context.Context
	ApiService PaymentsApi
}

func (r ApiGetAllConnectorsConfigsRequest) Execute() (*ListConnectorsConfigsResponse, *http.Response, error) {
	return r.ApiService.GetAllConnectorsConfigsExecute(r)
}

/*
GetAllConnectorsConfigs Get all available connectors configs

Get all available connectors configs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAllConnectorsConfigsRequest
*/
func (a *PaymentsApiService) GetAllConnectorsConfigs(ctx context.Context) ApiGetAllConnectorsConfigsRequest {
	return ApiGetAllConnectorsConfigsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListConnectorsConfigsResponse
func (a *PaymentsApiService) GetAllConnectorsConfigsExecute(r ApiGetAllConnectorsConfigsRequest) (*ListConnectorsConfigsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListConnectorsConfigsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentsApiService.GetAllConnectorsConfigs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/payments/connectors/configs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetConnectorTaskRequest struct {
	ctx context.Context
	ApiService PaymentsApi
	connector string
	taskId string
}

func (r ApiGetConnectorTaskRequest) Execute() (*ConnectorTask, *http.Response, error) {
	return r.ApiService.GetConnectorTaskExecute(r)
}

/*
GetConnectorTask Read a specific task of the connector

Get a specific task associated to the connector

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param connector The connector code
 @param taskId The task id
 @return ApiGetConnectorTaskRequest
*/
func (a *PaymentsApiService) GetConnectorTask(ctx context.Context, connector string, taskId string) ApiGetConnectorTaskRequest {
	return ApiGetConnectorTaskRequest{
		ApiService: a,
		ctx: ctx,
		connector: connector,
		taskId: taskId,
	}
}

// Execute executes the request
//  @return ConnectorTask
func (a *PaymentsApiService) GetConnectorTaskExecute(r ApiGetConnectorTaskRequest) (*ConnectorTask, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConnectorTask
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentsApiService.GetConnectorTask")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/payments/connectors/{connector}/tasks/{taskId}"
	localVarPath = strings.Replace(localVarPath, "{"+"connector"+"}", url.PathEscape(parameterValueToString(r.connector, "connector")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"taskId"+"}", url.PathEscape(parameterValueToString(r.taskId, "taskId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPaymentRequest struct {
	ctx context.Context
	ApiService PaymentsApi
	paymentId string
}

func (r ApiGetPaymentRequest) Execute() (*Payment, *http.Response, error) {
	return r.ApiService.GetPaymentExecute(r)
}

/*
GetPayment Returns a payment.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param paymentId The payment id
 @return ApiGetPaymentRequest
*/
func (a *PaymentsApiService) GetPayment(ctx context.Context, paymentId string) ApiGetPaymentRequest {
	return ApiGetPaymentRequest{
		ApiService: a,
		ctx: ctx,
		paymentId: paymentId,
	}
}

// Execute executes the request
//  @return Payment
func (a *PaymentsApiService) GetPaymentExecute(r ApiGetPaymentRequest) (*Payment, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Payment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentsApiService.GetPayment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/payments/payments/{paymentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"paymentId"+"}", url.PathEscape(parameterValueToString(r.paymentId, "paymentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiInstallConnectorRequest struct {
	ctx context.Context
	ApiService PaymentsApi
	connector string
	connectorConfig *ConnectorConfig
}

func (r ApiInstallConnectorRequest) ConnectorConfig(connectorConfig ConnectorConfig) ApiInstallConnectorRequest {
	r.connectorConfig = &connectorConfig
	return r
}

func (r ApiInstallConnectorRequest) Execute() (*http.Response, error) {
	return r.ApiService.InstallConnectorExecute(r)
}

/*
InstallConnector Install connector

Install connector

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param connector The connector code
 @return ApiInstallConnectorRequest
*/
func (a *PaymentsApiService) InstallConnector(ctx context.Context, connector string) ApiInstallConnectorRequest {
	return ApiInstallConnectorRequest{
		ApiService: a,
		ctx: ctx,
		connector: connector,
	}
}

// Execute executes the request
func (a *PaymentsApiService) InstallConnectorExecute(r ApiInstallConnectorRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentsApiService.InstallConnector")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/payments/connectors/{connector}"
	localVarPath = strings.Replace(localVarPath, "{"+"connector"+"}", url.PathEscape(parameterValueToString(r.connector, "connector")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.connectorConfig == nil {
		return nil, reportError("connectorConfig is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.connectorConfig
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiListConnectorTasksRequest struct {
	ctx context.Context
	ApiService PaymentsApi
	connector string
}

func (r ApiListConnectorTasksRequest) Execute() ([]ConnectorTask, *http.Response, error) {
	return r.ApiService.ListConnectorTasksExecute(r)
}

/*
ListConnectorTasks List connector tasks

List all tasks associated with this connector.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param connector The connector code
 @return ApiListConnectorTasksRequest
*/
func (a *PaymentsApiService) ListConnectorTasks(ctx context.Context, connector string) ApiListConnectorTasksRequest {
	return ApiListConnectorTasksRequest{
		ApiService: a,
		ctx: ctx,
		connector: connector,
	}
}

// Execute executes the request
//  @return []ConnectorTask
func (a *PaymentsApiService) ListConnectorTasksExecute(r ApiListConnectorTasksRequest) ([]ConnectorTask, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ConnectorTask
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentsApiService.ListConnectorTasks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/payments/connectors/{connector}/tasks"
	localVarPath = strings.Replace(localVarPath, "{"+"connector"+"}", url.PathEscape(parameterValueToString(r.connector, "connector")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListPaymentsRequest struct {
	ctx context.Context
	ApiService PaymentsApi
	limit *int32
	skip *int32
	sort *[]string
}

// Limit the number of payments to return, pagination can be achieved in conjunction with &#39;skip&#39; parameter.
func (r ApiListPaymentsRequest) Limit(limit int32) ApiListPaymentsRequest {
	r.limit = &limit
	return r
}

// How many payments to skip, pagination can be achieved in conjunction with &#39;limit&#39; parameter.
func (r ApiListPaymentsRequest) Skip(skip int32) ApiListPaymentsRequest {
	r.skip = &skip
	return r
}

// Field used to sort payments (Default is by date).
func (r ApiListPaymentsRequest) Sort(sort []string) ApiListPaymentsRequest {
	r.sort = &sort
	return r
}

func (r ApiListPaymentsRequest) Execute() (*ListPaymentsResponse, *http.Response, error) {
	return r.ApiService.ListPaymentsExecute(r)
}

/*
ListPayments Returns a list of payments.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListPaymentsRequest
*/
func (a *PaymentsApiService) ListPayments(ctx context.Context) ApiListPaymentsRequest {
	return ApiListPaymentsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListPaymentsResponse
func (a *PaymentsApiService) ListPaymentsExecute(r ApiListPaymentsRequest) (*ListPaymentsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListPaymentsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentsApiService.ListPayments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/payments/payments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
	    parameterAddToQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.skip != nil {
	    parameterAddToQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		t := *r.sort
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToQuery(localVarQueryParams, "sort", s.Index(i), "multi")
			}
		} else {
			parameterAddToQuery(localVarQueryParams, "sort", t, "multi")
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiReadConnectorConfigRequest struct {
	ctx context.Context
	ApiService PaymentsApi
	connector string
}

func (r ApiReadConnectorConfigRequest) Execute() (*ConnectorConfig, *http.Response, error) {
	return r.ApiService.ReadConnectorConfigExecute(r)
}

/*
ReadConnectorConfig Read connector config

Read connector config

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param connector The connector code
 @return ApiReadConnectorConfigRequest
*/
func (a *PaymentsApiService) ReadConnectorConfig(ctx context.Context, connector string) ApiReadConnectorConfigRequest {
	return ApiReadConnectorConfigRequest{
		ApiService: a,
		ctx: ctx,
		connector: connector,
	}
}

// Execute executes the request
//  @return ConnectorConfig
func (a *PaymentsApiService) ReadConnectorConfigExecute(r ApiReadConnectorConfigRequest) (*ConnectorConfig, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConnectorConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentsApiService.ReadConnectorConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/payments/connectors/{connector}/config"
	localVarPath = strings.Replace(localVarPath, "{"+"connector"+"}", url.PathEscape(parameterValueToString(r.connector, "connector")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiResetConnectorRequest struct {
	ctx context.Context
	ApiService PaymentsApi
	connector string
}

func (r ApiResetConnectorRequest) Execute() (*http.Response, error) {
	return r.ApiService.ResetConnectorExecute(r)
}

/*
ResetConnector Reset connector

Reset connector. Will remove the connector and ALL PAYMENTS generated with it.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param connector The connector code
 @return ApiResetConnectorRequest
*/
func (a *PaymentsApiService) ResetConnector(ctx context.Context, connector string) ApiResetConnectorRequest {
	return ApiResetConnectorRequest{
		ApiService: a,
		ctx: ctx,
		connector: connector,
	}
}

// Execute executes the request
func (a *PaymentsApiService) ResetConnectorExecute(r ApiResetConnectorRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentsApiService.ResetConnector")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/payments/connectors/{connector}/reset"
	localVarPath = strings.Replace(localVarPath, "{"+"connector"+"}", url.PathEscape(parameterValueToString(r.connector, "connector")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiUninstallConnectorRequest struct {
	ctx context.Context
	ApiService PaymentsApi
	connector string
}

func (r ApiUninstallConnectorRequest) Execute() (*http.Response, error) {
	return r.ApiService.UninstallConnectorExecute(r)
}

/*
UninstallConnector Uninstall connector

Uninstall  connector

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param connector The connector code
 @return ApiUninstallConnectorRequest
*/
func (a *PaymentsApiService) UninstallConnector(ctx context.Context, connector string) ApiUninstallConnectorRequest {
	return ApiUninstallConnectorRequest{
		ApiService: a,
		ctx: ctx,
		connector: connector,
	}
}

// Execute executes the request
func (a *PaymentsApiService) UninstallConnectorExecute(r ApiUninstallConnectorRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentsApiService.UninstallConnector")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/payments/connectors/{connector}"
	localVarPath = strings.Replace(localVarPath, "{"+"connector"+"}", url.PathEscape(parameterValueToString(r.connector, "connector")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
