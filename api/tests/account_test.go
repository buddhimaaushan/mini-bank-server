package api_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/buddhimaaushan/mini_bank/api"
	mockdb "github.com/buddhimaaushan/mini_bank/db/mock"
	"github.com/buddhimaaushan/mini_bank/db/sqlc"
	"github.com/buddhimaaushan/mini_bank/utils"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

func TestGetAccountAPI(t *testing.T) {
	// Create a new account
	account := randomAccount()

	testCases := []struct {
		name          string
		accountID     int64
		accountType   string
		balance       int64
		accStatus     sqlc.Status
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name:        "OK",
			accountID:   account.ID,
			accountType: account.Type,
			balance:     account.Balance,
			accStatus:   account.AccStatus,
			// Set up expectations
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(account, nil)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
				recoderBodyMatchAccount(t, recoder.Body, account)
			},
		},
		{
			name:        "NotFound",
			accountID:   account.ID,
			accountType: account.Type,
			balance:     account.Balance,
			accStatus:   account.AccStatus,
			// Set up expectations
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(sqlc.Account{}, pgx.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recoder.Code)
			},
		},
		{
			name:        "InternalError",
			accountID:   account.ID,
			accountType: account.Type,
			balance:     account.Balance,
			accStatus:   account.AccStatus,
			// Set up expectations
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(sqlc.Account{}, pgx.ErrTxClosed)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
		{
			name:        "InvalidID",
			accountID:   0,
			accountType: account.Type,
			balance:     account.Balance,
			accStatus:   account.AccStatus,
			// Set up expectations
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAccount(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Build mock
			store := mockdb.NewMockStore(ctrl)

			// Build stub
			tc.buildStubs(store)

			// Start test server
			server := api.NewServer(store)
			recoder := httptest.NewRecorder()

			// Make request
			url := fmt.Sprintf("/accounts/%d", tc.accountID)
			req, err := http.NewRequest("GET", url, nil)
			require.NoError(t, err)

			// Listen and serve Response
			server.Router.ServeHTTP(recoder, req)
			tc.checkResponse(t, recoder)
		})
	}

}

func randomAccount() sqlc.Account {
	return sqlc.Account{
		ID:      int64(utils.RandomInt(1, 1000)),
		Type:    utils.RandomString(6),
		Balance: int64(utils.RandomInt(100, 10000)),
	}
}

func recoderBodyMatchAccount(t *testing.T, body *bytes.Buffer, account sqlc.Account) {
	// Read the body
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	// Unmarshal the body
	var gotAccount sqlc.Account
	err = json.Unmarshal(data, &gotAccount)
	require.NoError(t, err)

	// Compare the body
	require.Equal(t, account, gotAccount)

}
