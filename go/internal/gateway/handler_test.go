package gateway

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PedroMosquera/promptgate/internal/cache"
	"github.com/PedroMosquera/promptgate/internal/llm"
	"github.com/PedroMosquera/promptgate/internal/ratelimit"
)

func testGateway(client llm.Client, limiter *ratelimit.Limiter) *Gateway {
	if limiter == nil {
		limiter = ratelimit.New(1000, 1000)
	}
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	return New(client, cache.New(), limiter, log)
}

func complete(g *Gateway, apiKey, body string) *httptest.ResponseRecorder {
	r := httptest.NewRequest(http.MethodPost, "/v1/complete", strings.NewReader(body))
	if apiKey != "" {
		r.Header.Set("X-API-Key", apiKey)
	}
	w := httptest.NewRecorder()
	g.Routes().ServeHTTP(w, r)
	return w
}

func decode(t *testing.T, w *httptest.ResponseRecorder) completeResponse {
	t.Helper()
	var resp completeResponse
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("decode response %q: %v", w.Body.String(), err)
	}
	return resp
}

func TestComplete_HappyPath(t *testing.T) {
	g := testGateway(llm.NewMock(), nil)
	w := complete(g, "key1", `{"model":"gpt-x","prompt":"hi"}`)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200 (body: %s)", w.Code, w.Body.String())
	}
	resp := decode(t, w)
	if resp.Cached {
		t.Error("first request should not be cached")
	}
	if resp.Model != "gpt-x" || resp.Completion == "" {
		t.Errorf("unexpected response: %+v", resp)
	}
}

func TestComplete_CacheHit(t *testing.T) {
	g := testGateway(llm.NewMock(), nil)
	body := `{"model":"gpt-x","prompt":"hi"}`

	complete(g, "key1", body)
	w := complete(g, "key1", body)

	resp := decode(t, w)
	if !resp.Cached {
		t.Error("second identical request should be served from cache")
	}
}

// TestComplete_DifferentModelsDoNotCollide: the same prompt sent to two models
// must return each model's own completion, not the first one cached.
func TestComplete_DifferentModelsDoNotCollide(t *testing.T) {
	g := testGateway(llm.NewMock(), nil)

	first := decode(t, complete(g, "key1", `{"model":"model-a","prompt":"same"}`))
	second := decode(t, complete(g, "key1", `{"model":"model-b","prompt":"same"}`))

	if second.Cached {
		t.Error("different model should be a cache miss, not a hit")
	}
	if second.Completion == first.Completion {
		t.Errorf("model-b returned model-a's cached completion: %q", second.Completion)
	}
}

func TestComplete_RateLimited(t *testing.T) {
	g := testGateway(llm.NewMock(), ratelimit.New(0, 2))
	body := `{"model":"m","prompt":"p"}`

	complete(g, "key1", body)
	complete(g, "key1", body)
	w := complete(g, "key1", body)

	if w.Code != http.StatusTooManyRequests {
		t.Errorf("status = %d, want 429", w.Code)
	}
}

func TestComplete_Validation(t *testing.T) {
	tests := []struct {
		name   string
		apiKey string
		body   string
		want   int
	}{
		{"missing api key", "", `{"model":"m","prompt":"p"}`, http.StatusUnauthorized},
		{"malformed json", "k", `{`, http.StatusBadRequest},
		{"missing prompt", "k", `{"model":"m"}`, http.StatusBadRequest},
		{"missing model", "k", `{"prompt":"p"}`, http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := testGateway(llm.NewMock(), nil)
			w := complete(g, tt.apiKey, tt.body)
			if w.Code != tt.want {
				t.Errorf("status = %d, want %d", w.Code, tt.want)
			}
		})
	}
}

type erroringClient struct{ err error }

func (e erroringClient) Complete(context.Context, llm.Request) (llm.Response, error) {
	return llm.Response{}, e.err
}

func TestComplete_UpstreamError(t *testing.T) {
	g := testGateway(erroringClient{err: io.ErrUnexpectedEOF}, nil)

	w := complete(g, "key1", `{"model":"m","prompt":"p"}`)
	if w.Code != http.StatusBadGateway {
		t.Errorf("status = %d, want 502", w.Code)
	}
}

func TestHealthz(t *testing.T) {
	g := testGateway(llm.NewMock(), nil)
	r := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	w := httptest.NewRecorder()
	g.Routes().ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
}
