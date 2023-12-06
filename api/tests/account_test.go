package api_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/buddhimaaushan/mini_bank/api"
	"github.com/buddhimaaushan/mini_bank/db"
	mockdb "github.com/buddhimaaushan/mini_bank/db/mock"
	"github.com/buddhimaaushan/mini_bank/db/sqlc"
	"github.com/buddhimaaushan/mini_bank/utils"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func TestGetAccountAPI(t *testing.T) {

	// Create users
	user1 := utils.CreateRandomUser()
	user2 := utils.CreateRandomUser()

	// Create a new account
	AccountTxResult := createRandomAccount([]sqlc.User{user1, user2})

	testCases := []struct {
		name          string
		accountID     int64
		accountType   string
		balance       int64
		accStatus     sqlc.Status
		accHolders    []sqlc.AccountHolder
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name:        "OK",
			accountID:   AccountTxResult.Account.ID,
			accountType: AccountTxResult.Account.Type,
			balance:     AccountTxResult.Account.Balance,
			accStatus:   AccountTxResult.Account.AccStatus,
			accHolders:  AccountTxResult.AccountHolders,
			// Set up expectations
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(AccountTxResult.Account.ID)).Times(1).Return(AccountTxResult.Account, nil)

				arg := sqlc.GetAccountHoldersByAccountIDParams{
					AccID:  AccountTxResult.Account.ID,
					Limit:  10,
					Offset: 0,
				}

				store.EXPECT().GetAccountHoldersByAccountID(gomock.Any(), gomock.Eq(arg)).Times(1).Return(AccountTxResult.AccountHolders, nil)

			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				arg := api.AccountResponse{
					ID:             AccountTxResult.Account.ID,
					Type:           AccountTxResult.Account.Type,
					Balance:        AccountTxResult.Account.Balance,
					AccountHolders: AccountTxResult.AccountHolders,
					Status:         AccountTxResult.Account.AccStatus,
					CreatedAt:      AccountTxResult.Account.CreatedAt,
				}
				require.Equal(t, http.StatusOK, recoder.Code)
				recoderBodyMatchAccount(t, recoder.Body, arg)
			},
		},
		{
			name:        "NotFound",
			accountID:   AccountTxResult.Account.ID,
			accountType: AccountTxResult.Account.Type,
			balance:     AccountTxResult.Account.Balance,
			accStatus:   AccountTxResult.Account.AccStatus,
			accHolders:  AccountTxResult.AccountHolders,
			// Set up expectations
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(AccountTxResult.Account.ID)).Times(1).Return(sqlc.Account{}, pgx.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recoder.Code)
			},
		},
		{
			name:        "InternalError",
			accountID:   AccountTxResult.Account.ID,
			accountType: AccountTxResult.Account.Type,
			balance:     AccountTxResult.Account.Balance,
			accStatus:   AccountTxResult.Account.AccStatus,
			accHolders:  AccountTxResult.AccountHolders,
			// Set up expectations
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(AccountTxResult.Account.ID)).Times(1).Return(sqlc.Account{}, pgx.ErrTxClosed)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
		{
			name:        "InvalidID",
			accountID:   0,
			accountType: AccountTxResult.Account.Type,
			balance:     AccountTxResult.Account.Balance,
			accStatus:   AccountTxResult.Account.AccStatus,
			accHolders:  AccountTxResult.AccountHolders,
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
			server := newTestServer(t, store)
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

func createRandomAccount(users []sqlc.User) db.AccountTxResult {
	var accountHolders = make([]sqlc.AccountHolder, len(users))

	account := sqlc.Account{
		ID:      int64(utils.RandomInt(1, 1000)),
		Type:    utils.RandomString(6),
		Balance: int64(utils.RandomInt(100, 10000)),
	}

	for _, user := range users {
		accountHolders = append(accountHolders, sqlc.AccountHolder{
			AccID:  account.ID,
			UserID: user.ID,
			CreatedAt: pgtype.Timestamptz{
				Time:  time.Now(),
				Valid: true,
			},
		})
	}

	return db.AccountTxResult{
		Account:        account,
		AccountHolders: accountHolders,
	}
}

func recoderBodyMatchAccount(t *testing.T, body *bytes.Buffer, account api.AccountResponse) {
	// Read the body
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	// Unmarshal the body
	var gotAccount api.AccountResponse
	err = json.Unmarshal(data, &gotAccount)
	require.NoError(t, err)

	// Compare the body
	require.Equal(t, account.ID, gotAccount.ID)
	require.Equal(t, account.Type, gotAccount.Type)
	require.Equal(t, account.Balance, gotAccount.Balance)

	// Compare account holders
	for i, expected := range account.AccountHolders {
		require.Equal(t, expected.AccID, gotAccount.AccountHolders[i].AccID)
		require.Equal(t, expected.UserID, gotAccount.AccountHolders[i].UserID)
		require.True(t, expected.CreatedAt.Time.Equal(gotAccount.AccountHolders[i].CreatedAt.Time))
	}

}
