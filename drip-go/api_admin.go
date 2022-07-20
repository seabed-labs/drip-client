/*
Drip Backend

Drip backend service. All API's have a IP rate limit of 10 requests per second. 

API version: 1.0.0
Contact: mocha@dcaf.so
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package drip

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


type AdminApi interface {

	/*
	AdminVaultPubkeyPathEnablePut Enable a vault

	Enable the specified vault

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param pubkeyPath
	@return ApiAdminVaultPubkeyPathEnablePutRequest
	*/
	AdminVaultPubkeyPathEnablePut(ctx context.Context, pubkeyPath string) ApiAdminVaultPubkeyPathEnablePutRequest

	// AdminVaultPubkeyPathEnablePutExecute executes the request
	//  @return Vault
	AdminVaultPubkeyPathEnablePutExecute(r ApiAdminVaultPubkeyPathEnablePutRequest) (*Vault, *http.Response, error)

	/*
	AdminVaultsGet Get All Vaults

	Get all vaults with filters and expanded properties.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAdminVaultsGetRequest
	*/
	AdminVaultsGet(ctx context.Context) ApiAdminVaultsGetRequest

	// AdminVaultsGetExecute executes the request
	//  @return []ExpandedAdminVault
	AdminVaultsGetExecute(r ApiAdminVaultsGetRequest) ([]ExpandedAdminVault, *http.Response, error)
}

// AdminApiService AdminApi service
type AdminApiService service

type ApiAdminVaultPubkeyPathEnablePutRequest struct {
	ctx context.Context
	ApiService AdminApi
	pubkeyPath string
	tokenId *string
}

func (r ApiAdminVaultPubkeyPathEnablePutRequest) TokenId(tokenId string) ApiAdminVaultPubkeyPathEnablePutRequest {
	r.tokenId = &tokenId
	return r
}

func (r ApiAdminVaultPubkeyPathEnablePutRequest) Execute() (*Vault, *http.Response, error) {
	return r.ApiService.AdminVaultPubkeyPathEnablePutExecute(r)
}

/*
AdminVaultPubkeyPathEnablePut Enable a vault

Enable the specified vault

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pubkeyPath
 @return ApiAdminVaultPubkeyPathEnablePutRequest
*/
func (a *AdminApiService) AdminVaultPubkeyPathEnablePut(ctx context.Context, pubkeyPath string) ApiAdminVaultPubkeyPathEnablePutRequest {
	return ApiAdminVaultPubkeyPathEnablePutRequest{
		ApiService: a,
		ctx: ctx,
		pubkeyPath: pubkeyPath,
	}
}

// Execute executes the request
//  @return Vault
func (a *AdminApiService) AdminVaultPubkeyPathEnablePutExecute(r ApiAdminVaultPubkeyPathEnablePutRequest) (*Vault, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Vault
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdminApiService.AdminVaultPubkeyPathEnablePut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/vault/{pubkeyPath}/enable"
	localVarPath = strings.Replace(localVarPath, "{"+"pubkeyPath"+"}", url.PathEscape(parameterToString(r.pubkeyPath, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tokenId == nil {
		return localVarReturnValue, nil, reportError("tokenId is required and must be specified")
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
	localVarHeaderParams["token-id"] = parameterToString(*r.tokenId, "")
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

type ApiAdminVaultsGetRequest struct {
	ctx context.Context
	ApiService AdminApi
	tokenId *string
	expand *[]string
	enabled *bool
	offset *int32
	limit *int32
}

func (r ApiAdminVaultsGetRequest) TokenId(tokenId string) ApiAdminVaultsGetRequest {
	r.tokenId = &tokenId
	return r
}

func (r ApiAdminVaultsGetRequest) Expand(expand []string) ApiAdminVaultsGetRequest {
	r.expand = &expand
	return r
}

func (r ApiAdminVaultsGetRequest) Enabled(enabled bool) ApiAdminVaultsGetRequest {
	r.enabled = &enabled
	return r
}

func (r ApiAdminVaultsGetRequest) Offset(offset int32) ApiAdminVaultsGetRequest {
	r.offset = &offset
	return r
}

func (r ApiAdminVaultsGetRequest) Limit(limit int32) ApiAdminVaultsGetRequest {
	r.limit = &limit
	return r
}

func (r ApiAdminVaultsGetRequest) Execute() ([]ExpandedAdminVault, *http.Response, error) {
	return r.ApiService.AdminVaultsGetExecute(r)
}

/*
AdminVaultsGet Get All Vaults

Get all vaults with filters and expanded properties.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAdminVaultsGetRequest
*/
func (a *AdminApiService) AdminVaultsGet(ctx context.Context) ApiAdminVaultsGetRequest {
	return ApiAdminVaultsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ExpandedAdminVault
func (a *AdminApiService) AdminVaultsGetExecute(r ApiAdminVaultsGetRequest) ([]ExpandedAdminVault, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ExpandedAdminVault
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdminApiService.AdminVaultsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/vaults"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tokenId == nil {
		return localVarReturnValue, nil, reportError("tokenId is required and must be specified")
	}

	if r.expand != nil {
		t := *r.expand
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("expand", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("expand", parameterToString(t, "multi"))
		}
	}
	if r.enabled != nil {
		localVarQueryParams.Add("enabled", parameterToString(*r.enabled, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
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
	localVarHeaderParams["token-id"] = parameterToString(*r.tokenId, "")
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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