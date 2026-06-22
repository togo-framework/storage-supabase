# storage-supabase

Supabase Storage driver for togo (implements `togo.Storage`). Install it to store
blobs in a Supabase bucket instead of the local filesystem.

```bash
togo install togo-framework/storage-supabase
```

Env: `SUPABASE_URL`, `SUPABASE_SERVICE_KEY`, `SUPABASE_STORAGE_BUCKET` (default `public`).
