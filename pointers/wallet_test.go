package pointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, expected Bitcoin) {
		t.Helper()
		actual := wallet.Balance()
		assert.Equal(t, expected, actual, "they should be equal")
	}

	assertError := func(t *testing.T, theError error, errString string) {
		t.Helper()
		assert.EqualError(t, theError, errString)
	}

	assertNoError := func(t *testing.T, theError error) {
		t.Helper()
		assert.Nil(t, theError)
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(10)

		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertError(t, err, ErrInsufficientFunds.Error())
	})
}
