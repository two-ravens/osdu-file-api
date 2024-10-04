# FileLocationAPI
(*FileLocationAPI*)

## Overview

File Location API

### Available Operations

* [GetLocationFile](#getlocationfile) - Get a location in Landing Zone to upload a file.

## GetLocationFile

Gets a temporary signed URL to upload a file (Service does not upload the file by itself, User needs to use this URL to upload the file). The generated URL is time bound and by default expires by 7 days maximum. <p> User will receive a FileSource in the response.This is the relative path where the uploaded file will persist. Once the file is uploaded, FileSource can then be used to post metadata of the file.</p> <p> **Required roles**: `service.file.editors`. Users added to groups `users.datalake.editors`, `users.datalake.admins`, `users.datalake.ops` would be added to group `service.file.editors` by default.</p>

### Example Usage

```go
package main

import(
	"openapi"
	"context"
	"log"
)

func main() {
    s := openapi.New(
        openapi.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.FileLocationAPI.GetLocationFile(ctx, "<id>", openapi.String("5M"))
    if err != nil {
        log.Fatal(err)
    }
    if res.Body != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                  | Type                                                                                                                                                                                                                                       | Required                                                                                                                                                                                                                                   | Description                                                                                                                                                                                                                                | Example                                                                                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                                                         | The context to use for the request.                                                                                                                                                                                                        |                                                                                                                                                                                                                                            |
| `dataPartitionID`                                                                                                                                                                                                                          | *string*                                                                                                                                                                                                                                   | :heavy_check_mark:                                                                                                                                                                                                                         | Tenant Id                                                                                                                                                                                                                                  |                                                                                                                                                                                                                                            |
| `expiryTime`                                                                                                                                                                                                                               | **string*                                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                                         | The Time for which Signed URL to be valid. Accepted Regex patterns are "^[0-9]+M$", "^[0-9]+H$", "^[0-9]+D$" denoting Integer values in Minutes, Hours, Days respectively. In absence of this parameter the URL would be valid for 1 Hour. | 5M                                                                                                                                                                                                                                         |
| `opts`                                                                                                                                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                                         | The options for this request.                                                                                                                                                                                                              |                                                                                                                                                                                                                                            |

### Response

**[*operations.GetLocationFileResponse](../../models/operations/getlocationfileresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |