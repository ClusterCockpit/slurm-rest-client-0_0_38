# Go API client for openapi

API to access and control Slurm.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.0.38
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://www.schedmd.com/](https://www.schedmd.com/)

## Known Issues
- JSON Data Type Mismatch
A data type mismatch exists between the Slurm JSON payload and the Go representation in `V0038MetaSlurmVersion`. The Go implementation expects `*string`, while the JSON payload provides numerical values.

```go
type V0038MetaSlurmVersion struct {
    Major *string `json:"major,omitempty"`
    Micro *string `json:"micro,omitempty"`
    Minor *string `json:"minor,omitempty"`
}
```

```json
"Slurm": {
    "version": {
         "major": 22,
         "micro": 2,
         "minor": 5
    }
}
```
Another mismatch arises when `V0038JobResponseProperties.jobs.exit_code` contains values like `4294967295`, exceeding the maximum limit of `int32`, which is `2147483647`. This results in an unmarshaling error in Go.

- Empty Map for `V0038NodeAllocationSockets.Cores`
The field `V0038NodeAllocationSockets.Cores` appears to be an empty map. Additionally, the results from Slurm, particularly the `cpus_used` value being `0`, appear anomalous and require further investigation. This behavior may potentially stem from a conversion issue within the library.

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
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
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost/slurm/v0.0.38*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*OpenapiAPI* | [**OpenapiGet**](docs/OpenapiAPI.md#openapiget) | **Get** /openapi | Retrieve OpenAPI Specification
*OpenapiAPI* | [**OpenapiJsonGet**](docs/OpenapiAPI.md#openapijsonget) | **Get** /openapi.json | Retrieve OpenAPI Specification
*OpenapiAPI* | [**OpenapiV3Get**](docs/OpenapiAPI.md#openapiv3get) | **Get** /openapi/v3 | Retrieve OpenAPI Specification
*OpenapiAPI* | [**OpenapiYamlGet**](docs/OpenapiAPI.md#openapiyamlget) | **Get** /openapi.yaml | Retrieve OpenAPI Specification
*SlurmAPI* | [**CancelJob**](docs/SlurmAPI.md#canceljob) | **Delete** /job/{job_id} | cancel or signal job
*SlurmAPI* | [**Diag**](docs/SlurmAPI.md#diag) | **Get** /diag/ | get diagnostics
*SlurmAPI* | [**GetJob**](docs/SlurmAPI.md#getjob) | **Get** /job/{job_id} | get job info
*SlurmAPI* | [**GetJobs**](docs/SlurmAPI.md#getjobs) | **Get** /jobs/ | get list of jobs
*SlurmAPI* | [**GetNode**](docs/SlurmAPI.md#getnode) | **Get** /node/{node_name} | get node info
*SlurmAPI* | [**GetNodes**](docs/SlurmAPI.md#getnodes) | **Get** /nodes/ | get all node info
*SlurmAPI* | [**GetPartition**](docs/SlurmAPI.md#getpartition) | **Get** /partition/{partition_name} | get partition info
*SlurmAPI* | [**GetPartitions**](docs/SlurmAPI.md#getpartitions) | **Get** /partitions/ | get all partition info
*SlurmAPI* | [**GetReservation**](docs/SlurmAPI.md#getreservation) | **Get** /reservation/{reservation_name} | get reservation info
*SlurmAPI* | [**GetReservations**](docs/SlurmAPI.md#getreservations) | **Get** /reservations/ | get all reservation info
*SlurmAPI* | [**Ping**](docs/SlurmAPI.md#ping) | **Get** /ping/ | ping test
*SlurmAPI* | [**SlurmctldGetLicenses**](docs/SlurmAPI.md#slurmctldgetlicenses) | **Get** /licenses/ | get all Slurm tracked license info
*SlurmAPI* | [**SubmitJob**](docs/SlurmAPI.md#submitjob) | **Post** /job/submit | submit new job
*SlurmAPI* | [**UpdateJob**](docs/SlurmAPI.md#updatejob) | **Post** /job/{job_id} | update job


## Documentation For Models

 - [V0038Diag](docs/V0038Diag.md)
 - [V0038DiagRpcm](docs/V0038DiagRpcm.md)
 - [V0038DiagRpcu](docs/V0038DiagRpcu.md)
 - [V0038DiagStatistics](docs/V0038DiagStatistics.md)
 - [V0038Error](docs/V0038Error.md)
 - [V0038JobProperties](docs/V0038JobProperties.md)
 - [V0038JobResources](docs/V0038JobResources.md)
 - [V0038JobResponseProperties](docs/V0038JobResponseProperties.md)
 - [V0038JobSubmission](docs/V0038JobSubmission.md)
 - [V0038JobSubmissionResponse](docs/V0038JobSubmissionResponse.md)
 - [V0038JobsResponse](docs/V0038JobsResponse.md)
 - [V0038License](docs/V0038License.md)
 - [V0038Licenses](docs/V0038Licenses.md)
 - [V0038Meta](docs/V0038Meta.md)
 - [V0038MetaPlugin](docs/V0038MetaPlugin.md)
 - [V0038MetaSlurm](docs/V0038MetaSlurm.md)
 - [V0038MetaSlurmVersion](docs/V0038MetaSlurmVersion.md)
 - [V0038Node](docs/V0038Node.md)
 - [V0038NodeAllocation](docs/V0038NodeAllocation.md)
 - [V0038NodeAllocationSockets](docs/V0038NodeAllocationSockets.md)
 - [V0038NodesResponse](docs/V0038NodesResponse.md)
 - [V0038Partition](docs/V0038Partition.md)
 - [V0038PartitionsResponse](docs/V0038PartitionsResponse.md)
 - [V0038Ping](docs/V0038Ping.md)
 - [V0038Pings](docs/V0038Pings.md)
 - [V0038Reservation](docs/V0038Reservation.md)
 - [V0038ReservationPurgeCompleted](docs/V0038ReservationPurgeCompleted.md)
 - [V0038ReservationsResponse](docs/V0038ReservationsResponse.md)
 - [V0038Signal](docs/V0038Signal.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### user

- **Type**: API key
- **API key parameter name**: X-SLURM-USER-NAME
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-SLURM-USER-NAME and passed in as the auth context for each request.

Example

```golang
auth := context.WithValue(
		context.Background(),
		sw.ContextAPIKeys,
		map[string]sw.APIKey{
			"X-SLURM-USER-NAME": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### token

- **Type**: API key
- **API key parameter name**: X-SLURM-USER-TOKEN
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-SLURM-USER-TOKEN and passed in as the auth context for each request.

Example

```golang
auth := context.WithValue(
		context.Background(),
		sw.ContextAPIKeys,
		map[string]sw.APIKey{
			"X-SLURM-USER-TOKEN": {Key: "API_KEY_STRING"},
		},
	)
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

sales@schedmd.com

