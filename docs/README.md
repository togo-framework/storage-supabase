# storage-supabase — documentation

  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />

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

Environment variables read by this plugin (extracted from the source):

| Env var | Notes |
|---|---|
| `G` | _see provider docs_ |
| `SUPABASE_SERVICE_KEY` | _see provider docs_ |
| `SUPABASE_STORAGE_BUCKET` | _see provider docs_ |
| `SUPABASE_URL` | _see provider docs_ |

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
- README: ../README.md
