package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DavidKorochik/issues4missions-backend/pkg/models"
	"github.com/DavidKorochik/issues4missions-backend/pkg/util"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func randomIssue() models.Issue {
	return models.Issue{
		Title:       util.RandomTitle(),
		Description: util.RandomDescription(),
		Status:      "open",
		IsCompleted: false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   gorm.DeletedAt{},
	}
}

func requireBodyMatchIssue(t *testing.T, body *bytes.Buffer, issue models.Issue) {
	var gotIssue models.Issue

	bodyBuffer, err := ioutil.ReadAll(body)
	require.NoError(t, err)
	require.NotEmpty(t, bodyBuffer)

	err = json.Unmarshal(bodyBuffer, &gotIssue)
	require.NoError(t, err)

	require.Equal(t, issue, gotIssue)
}

func TestCreateIssue(t *testing.T) {
	issue := randomIssue()

	body := models.Issue{
		Title:       issue.Title,
		Description: issue.Description,
	}

	bodyBuffer, err := json.Marshal(&body)
	require.NoError(t, err)

	recorder := httptest.NewRecorder()
	r := newTestServer(recorder)
	request, err := http.NewRequest(http.MethodPost, "/api/issues", bytes.NewBuffer(bodyBuffer))
	require.NoError(t, err)

	r.ServeHTTP(recorder, request)

	require.Equal(t, http.StatusOK, recorder.Code)
	requireBodyMatchIssue(t, recorder.Body, issue)
}
