package pointer

import "testing"

func TestPointer(t *testing.T) {

	t.Run("wallet collect Bitcoin", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10.0)
		get := wallet.Balance()
		want := 10
		if get != float64(want) {
			t.Errorf("got %f want %d", get, want)
		}
	})
}
