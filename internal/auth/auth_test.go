package auth_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
	"github.com/stretchr/testify/require"
)

func TestGetAPIKey(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	req.Header.Add("Authorization", "ApiKey somerandomvalue123")
	_, err := auth.GetAPIKey(req.Header)
	require.NoError(t, err)
}
