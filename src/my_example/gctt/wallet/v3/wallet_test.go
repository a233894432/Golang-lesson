package main

import (
	"testing"
)

// func TestWallet(t *testing.T) {

// 	wallet := Wallet{}

// 	wallet.Deposit(Bitcoin(10))

// 	got := wallet.Balance()

// 	fmt.Println("address of balance in test is", &wallet.balance)

// 	want := Bitcoin(10)

// 	if got != want {
// 		t.Errorf("got %d want %d", got, want)
// 	}
// }

// func TestWallet(t *testing.T) {

// 	t.Run("Deposit", func(t *testing.T) {
// 		wallet := Wallet{}

// 		wallet.Deposit(Bitcoin(10))

// 		got := wallet.Balance()

// 		want := Bitcoin(10)

// 		if got != want {
// 			t.Errorf("got %s want %s", got, want)
// 		}
// 	})

// 	t.Run("Withdraw", func(t *testing.T) {
// 		wallet := Wallet{balance: Bitcoin(20)}

// 		wallet.Withdraw(10)

// 		got := wallet.Balance()

// 		want := Bitcoin(10)

// 		if got != want {
// 			t.Errorf("got %s want %s", got, want)
// 		}
// 	})

// }

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

}
