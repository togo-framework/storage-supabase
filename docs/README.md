# storage-supabase — documentation

Supabase Storage driver for togo

## Overview

Package supastorage is a Supabase Storage driver for togo (implements
togo.Storage). Blank-import to store blobs in a Supabase bucket; overrides the
default filesystem storage. Install: `togo install togo-framework/storage-supabase`.

## Install

```bash
togo install togo-framework/storage-supabase
```

Set `STORAGE_DRIVER=supabase`.

## Configuration

Environment variables read by this plugin (extracted from the source — see the gateway/provider docs for each value):

| Env var |
|---|
| `SUPABASE_SERVICE_KEY"` |
| `SUPABASE_STORAGE_BUCKET"` |
| `SUPABASE_URL"` |

## Usage

```go
st := k.Storage
st.Put(ctx, "path/file.txt", data)
b, _ := st.Get(ctx, "path/file.txt")
url := st.Path("path/file.txt")
```

## Links

- Marketplace: https://to-go.dev/marketplace
- Source: https://github.com/togo-framework/storage-supabase
- Full README: ../README.md
