package pointers

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("expected %s, got %s", want, got)
		}
	}

	assertError := func(t testing.TB, err error, want string) {
		t.Helper()
		if err == nil {
			t.Fatal("expected error but got nil")
		}

		if err.Error() != want {
			t.Errorf("got %q, want %q", err, want)
		}
	}

	assertNoError := func(t testing.TB, got error) {
		if got != nil {
			t.Fatal("Not expecting any error")
		}
	}
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))

	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Insufficient balance", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(20))
		assertError(t, err, ErrInsufficientBalance)
		assertBalance(t, wallet, Bitcoin(10))
	})

}
