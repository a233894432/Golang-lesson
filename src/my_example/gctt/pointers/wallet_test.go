package main

import (
	"testing"
)

// func TestWallet(t *testing.T) {

// 	wallet := Wallet{}

// 	wallet.Deposit(10)

// 	got := wallet.Balance()
// 	fmt.Println("address of balance in test is", &wallet.balance)
// 	want := 10

// 	if got != want {
// 		t.Errorf("got %d want %d", got, want)
// 	}
// }

// part 01
// func TestWallet(t *testing.T) {
// 	wallet := Wallet{}

// 	wallet.Deposit(Bitcoin(10))

// 	got := wallet.Balance()

// 	want := Bitcoin(10)

// 	if got != want {
// 		t.Errorf("got %d want %d", got, want)
// 	}
// }

// part02

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

// part 03
// func TestWalletTwo(t *testing.T) {

// 	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
// 		got := wallet.Balance()

// 		if got != want {
// 			t.Errorf("got %s want %s", got, want)
// 		}
// 	}
// 	assertError := func(t *testing.T, got error, want string) {
// 		if got == nil {
// 			t.Fatal("didn't get an error but wanted one")
// 		}

// 		if got.Error() != want {
// 			t.Errorf("got '%s', want '%s'", got, want)
// 		}
// 	}

// 	t.Run("Deposit", func(t *testing.T) {
// 		wallet := Wallet{}
// 		wallet.Deposit(Bitcoin(10))
// 		assertBalance(t, wallet, Bitcoin(10))
// 	})

// 	t.Run("Withdraw", func(t *testing.T) {
// 		wallet := Wallet{balance: Bitcoin(20)}
// 		wallet.Withdraw(Bitcoin(10))
// 		assertBalance(t, wallet, Bitcoin(10))
// 	})

// 	t.Run("Withdraw insufficient funds", func(t *testing.T) {
// 		startingBalance := Bitcoin(20)
// 		wallet := Wallet{startingBalance}
// 		err := wallet.Withdraw(Bitcoin(100))

// 		assertBalance(t, wallet, startingBalance)

// 		if err == nil {
// 			t.Error("wanted an error but didn't get one")
// 		}
// 	})

// 	t.Run("Withdraw insufficient funds", func(t *testing.T) {
// 		wallet := Wallet{Bitcoin(20)}
// 		err := wallet.Withdraw(Bitcoin(100))

// 		assertBalance(t, wallet, Bitcoin(20))
// 		assertError(t, err, "cannot withdraw, insufficient funds")
// 	})

// }

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))

		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, Bitcoin(20))
		assertError(t, err, InsufficientFundsError)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	got := wallet.Balance()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	if got != nil {
		t.Fatal("got an error but didnt want one")
	}
}

func assertError(t *testing.T, got error, want error) {
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
