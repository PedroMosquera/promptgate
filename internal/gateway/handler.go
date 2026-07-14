package gateway

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/PedroMosquera/promptgate/internal/cache"
	"github.com/PedroMosquera/promptgate/internal/llm"
	"github.com/PedroMosquera/promptgate/internal/ratelimit"
)

type Gateway struct {
	client  llm.Client
	cache   *cache.Cache
	limiter *ratelimit.Limiter
	log     *slog.Logger
}

func New(client llm.Client, c *cache.Cache, l *ratelimit.Limiter, log *slog.Logger) *Gateway {
	if log == nil {
		log = slog.Default()
	}
	return &Gateway{client: client, cache: c, limiter: l, log: log}
}

func (g *Gateway) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /v1/complete", g.handleComplete)
	mux.HandleFunc("GET /healthz", g.handleHealth)
	return mux
}

type completeRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

type completeResponse struct {
	Model      string `json:"model"`
	Completion string `json:"completion"`
	Cached     bool   `json:"cached"`
}

func (g *Gateway) handleComplete(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-API-Key")
	if apiKey == "" {
		writeError(w, http.StatusUnauthorized, "missing X-API-Key header")
		return
	}
	if !g.limiter.Allow(apiKey) {
		writeError(w, http.StatusTooManyRequests, "rate limit exceeded")
		return
	}

	var req completeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if req.Model == "" || req.Prompt == "" {
		writeError(w, http.StatusBadRequest, "model and prompt are required")
		return
	}

	if completion, ok := g.cache.Get(req.Model, req.Prompt); ok {
		writeJSON(w, http.StatusOK, completeResponse{Model: req.Model, Completion: completion, Cached: true})
		return
	}

	resp, err := g.client.Complete(r.Context(), llm.Request{Model: req.Model, Prompt: req.Prompt})
	if err != nil {
		g.log.ErrorContext(r.Context(), "completion failed", "model", req.Model, "err", err)
		writeError(w, http.StatusBadGateway, "upstream error")
		return
	}

	g.cache.Set(req.Model, req.Prompt, resp.Completion)
	writeJSON(w, http.StatusOK, completeResponse{Model: resp.Model, Completion: resp.Completion, Cached: false})
}

func (g *Gateway) handleHealth(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
