package http_test

import (
	"net/http"
	test "net/http/httptest"
	"testing"

	httpServer "github.com/kecci/goscription/internal/http"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCORS(t *testing.T) {
	e := echo.New()
	req := test.NewRequest(echo.GET, "/", nil)
	res := test.NewRecorder()
	c := e.NewContext(req, res)
	m := httpServer.InitMiddleware()

	h := m.CORS(echo.HandlerFunc(func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	}))

	err := h(c)
	assert.NoError(t, err)
	assert.Equal(t, "*", res.Header().Get("Access-Control-Allow-Origin"))
}
