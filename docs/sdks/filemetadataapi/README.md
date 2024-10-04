# FileMetadataAPI
(*FileMetadataAPI*)

## Overview

File Metadata API

### Available Operations

* [PostFilesMetadata](#postfilesmetadata) - Creates a metadata for a file
* [GetFileMetadataByID](#getfilemetadatabyid) - Gets metadata record for the given id
* [DeleteFileMetadataByID](#deletefilemetadatabyid) - Deletes metadata record & file associated with that record for the given id

## PostFilesMetadata

This API creates a metadata record for a file that is already uploaded. The Metadata is linked to the file via `FileSource` provided in the request body. <p> If `FileSource` attribute is missing in the request body or there is no file present, then the request fails with an error. </p><p> When metadata is successfully updated in the system, it returns the `Id` of the file metadata record. </p><p> **Required roles**: `service.file.editors`. Users added to groups `users.datalake.editors`, `users.datalake.admins`, `users.datalake.ops` would be added to group `service.file.editors` by default.</p>

### Example Usage

```go
package main

import(
	"openapi"
	"context"
	"openapi/models/components"
	"log"
)

func main() {
    s := openapi.New(
        openapi.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.FileMetadataAPI.PostFilesMetadata(ctx, "<id>", components.FileMetadata{
        Kind: "osdu:wks:dataset--File.Generic:1.0.0",
        Legal: components.Legal{
            OtherRelevantDataCountries: []string{
                "<value>",
                "<value>",
                "<value>",
            },
        },
        Data: components.FileData{
            DatasetProperties: components.DatasetProperties{
                FileSourceInfo: components.FileSourceInfo{
                    FileSource: "<value>",
                },
            },
        },
        ACL: components.ACL{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Body != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `dataPartitionID`                                                  | *string*                                                           | :heavy_check_mark:                                                 | Tenant Id                                                          |
| `fileMetadata`                                                     | [components.FileMetadata](../../models/components/filemetadata.md) | :heavy_check_mark:                                                 | N/A                                                                |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.PostFilesMetadataResponse](../../models/operations/postfilesmetadataresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetFileMetadataByID

Gets the latest version of File metadata record identified by the given id. <p> **Required roles**: `service.file.editors`. Users added to groups `users.datalake.editors`, `users.datalake.admins`, `users.datalake.ops` would be added to group `service.file.editors` by default.</p>

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
    res, err := s.FileMetadataAPI.GetFileMetadataByID(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Body != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | File metadata record Id.                                 |
| `dataPartitionID`                                        | *string*                                                 | :heavy_check_mark:                                       | Tenant Id                                                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetFileMetadataByIDResponse](../../models/operations/getfilemetadatabyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteFileMetadataByID

Deletes the File metadata record identified by the given id and file associated with that metadata record. <p> **Required roles**: `users.datalake.editors`  or `users.datalake.admins`.</p>

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
    res, err := s.FileMetadataAPI.DeleteFileMetadataByID(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | File metadata record Id.                                 |
| `dataPartitionID`                                        | *string*                                                 | :heavy_check_mark:                                       | Tenant Id                                                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteFileMetadataByIDResponse](../../models/operations/deletefilemetadatabyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |