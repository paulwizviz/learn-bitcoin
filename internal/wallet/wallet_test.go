package wallet

import (
	"errors"
	"fmt"
	"testing"
)

func TestMnemonicError(t *testing.T) {
	testcases := []struct {
		bitsize int
		want    error
	}{
		{
			bitsize: 64, // Invalid value
			want:    ErrMnemonic,
		},
		{
			bitsize: 128, // Valid value
			want:    nil,
		},
		{
			bitsize: 129, // Invalid value
			want:    ErrMnemonic,
		},
		{
			bitsize: 256, // Valid value
			want:    nil,
		},
		{
			bitsize: 257, // Invalid value
			want:    ErrMnemonic,
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			if tc.want == nil {
				_, got := Mnemonic(tc.bitsize)
				if got != tc.want {
					t.Errorf("Want: %v Got: %v", tc.want, got)
				}
			} else {
				_, got := Mnemonic(tc.bitsize)
				if !errors.Is(got, ErrMnemonic) {
					t.Errorf("Want: %v Got: %v", tc.want, got)
				}
			}
		})
	}
}

func Example_masterkey() {
	m := "tube resemble retreat short holiday grain attack give rose autumn dog mango"
	mKey, _ := MasterHDKey(m, "test")
	cKey0, _ := mKey.NewChildKey(0)
	cKey1, _ := mKey.NewChildKey(1)
	fmt.Println("Master private key: ", mKey)
	fmt.Println("Child 0 private key: ", cKey0)
	fmt.Println("Child 1 private key: ", cKey1)
	// Output:
	// Master private key:  xprv9s21ZrQH143K2kt5XjSXGSGVFEZccJ1q5efUdD1Y2BkmG5yR2P5bigDZJFzMK19XzNbkYxpQL1rYbnrD8bBzmTEGqWP8ZM4kd8qu1s2gHn2
	// Child 0 private key:  xprv9twC5ZYBAxukZzTx2kCmD3kRmry2uoJ6sRA479FdYFxRTahp4MH14XKBAiCmVKdeVoLUqWLzEkvCHvaw6RdF5zcR9WXD5eHkiT2uhEYQ3Da
	// Child 1 private key:  xprv9twC5ZYBAxukbjobDQYBG5giCfiSfHhPp6vpCaHD2cJ33PPfFDqGccume8Mkd6ktfAo2EXv7r9JEPZq7XZ59TDLQjkkRqmNpSXHKe9ZE69j
}
