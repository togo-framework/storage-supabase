// Package supastorage is a Supabase Storage driver for togo (implements
// togo.Storage). Blank-import to store blobs in a Supabase bucket; overrides the
// default filesystem storage. Install: `togo install togo-framework/storage-supabase`.
package supastorage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/togo-framework/togo"
)

func init() {
	togo.RegisterProviderFunc("storage-supabase", togo.PriorityService+10, func(k *togo.Kernel) error {
		url := strings.TrimRight(os.Getenv("SUPABASE_URL"), "/")
		key := os.Getenv("SUPABASE_SERVICE_KEY")
		if url == "" || key == "" {
			if k.Log != nil {
				k.Log.Warn("storage-supabase: SUPABASE_URL/SUPABASE_SERVICE_KEY not set; skipping")
			}
			return nil
		}
		bucket := os.Getenv("SUPABASE_STORAGE_BUCKET")
		if bucket == "" {
			bucket = "public"
		}
		k.Storage = &store{base: url, key: key, bucket: bucket, client: &http.Client{Timeout: 30 * time.Second}}
		return nil
	})
}

type store struct {
	base, key, bucket string
	client            *http.Client
}

func (s *store) objectURL(path string) string {
	return fmt.Sprintf("%s/storage/v1/object/%s/%s", s.base, s.bucket, strings.TrimPrefix(path, "/"))
}

func (s *store) do(method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(context.Background(), method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+s.key)
	if body != nil {
		req.Header.Set("Content-Type", "application/octet-stream")
		req.Header.Set("x-upsert", "true")
	}
	return s.client.Do(req)
}

func (s *store) Put(path string, data []byte) error {
	resp, err := s.do(http.MethodPost, s.objectURL(path), bytes.NewReader(data))
	if err != nil {
		return err
	}
	return drain(resp)
}

func (s *store) Get(path string) ([]byte, error) {
	resp, err := s.do(http.MethodGet, s.objectURL(path), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("storage-supabase: get status %d", resp.StatusCode)
	}
	return io.ReadAll(resp.Body)
}

func (s *store) Delete(path string) error {
	resp, err := s.do(http.MethodDelete, s.objectURL(path), nil)
	if err != nil {
		return err
	}
	return drain(resp)
}

// Path returns the public object URL.
func (s *store) Path(path string) string {
	return fmt.Sprintf("%s/storage/v1/object/public/%s/%s", s.base, s.bucket, strings.TrimPrefix(path, "/"))
}

func drain(resp *http.Response) error {
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("storage-supabase: status %d: %s", resp.StatusCode, string(b))
	}
	_, _ = io.Copy(io.Discard, resp.Body)
	return nil
}
