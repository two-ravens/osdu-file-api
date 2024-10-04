# FileDeliveryAPI
(*FileDeliveryAPI*)

## Overview

File Delivery API

### Available Operations

* [DownloadURL](#downloadurl) - Gets a URL to download the file

## DownloadURL

Gets a URL for downloading the file associated with the unique `id`.By default, the download URL is valid for `1 Hour` and it is `7 Days` maximum.<p> **Required roles**: `service.file.viewers`. Users added to groups `users.datalake.viewers`,`users.datalake.editors`, `users.datalake.admins`, `users.datalake.ops` would be added to group `service.file.viewers` by default.</p>

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
    res, err := s.FileDeliveryAPI.DownloadURL(ctx, "<id>", "<id>", openapi.String("5M"))
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
| `id`                                                                                                                                                                                                                                       | *string*                                                                                                                                                                                                                                   | :heavy_check_mark:                                                                                                                                                                                                                         | File Metadata record Id.                                                                                                                                                                                                                   |                                                                                                                                                                                                                                            |
| `dataPartitionID`                                                                                                                                                                                                                          | *string*                                                                                                                                                                                                                                   | :heavy_check_mark:                                                                                                                                                                                                                         | Tenant Id                                                                                                                                                                                                                                  |                                                                                                                                                                                                                                            |
| `expiryTime`                                                                                                                                                                                                                               | **string*                                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                                         | The Time for which Signed URL to be valid. Accepted Regex patterns are "^[0-9]+M$", "^[0-9]+H$", "^[0-9]+D$" denoting Integer values in Minutes, Hours, Days respectively. In absence of this parameter the URL would be valid for 1 Hour. | 5M                                                                                                                                                                                                                                         |
| `opts`                                                                                                                                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                                         | The options for this request.                                                                                                                                                                                                              |                                                                                                                                                                                                                                            |

### Response

**[*operations.DownloadURLResponse](../../models/operations/downloadurlresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |