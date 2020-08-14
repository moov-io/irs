# \IrsFilesApi

All URIs are relative to *https://local.moov.io:8208*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Convert**](IrsFilesApi.md#Convert) | **Post** /convert | Convert irs file
[**Health**](IrsFilesApi.md#Health) | **Get** /health | health irs service
[**Print**](IrsFilesApi.md#Print) | **Post** /print | Print irs file with specific format
[**Validator**](IrsFilesApi.md#Validator) | **Post** /validator | Validate irs file



## Convert

> *os.File Convert(ctx, optional)

Convert irs file

Convert from original irs file to new irs file

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConvertOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConvertOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **optional.String**| print irs file type | [default to json]
 **generate** | **optional.Bool**| generate new trailer record | [default to false]
 **file** | **optional.Interface of *os.File****optional.*os.File**| irs file to upload | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/octet-stream, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Health

> string Health(ctx, )

health irs service

Check the irs service to check if running

### Required Parameters

This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Print

> string Print(ctx, optional)

Print irs file with specific format

Print irs file with requested file format.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PrintOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PrintOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **optional.String**| print irs file type | [default to json]
 **file** | **optional.Interface of *os.File****optional.*os.File**| irs file to upload | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Validator

> string Validator(ctx, optional)

Validate irs file

Validation of irs file.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidatorOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidatorOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | **optional.Interface of *os.File****optional.*os.File**| irs file to upload | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

