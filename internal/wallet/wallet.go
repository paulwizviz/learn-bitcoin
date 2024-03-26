// Package wallet contains utilities to create, orchestrate and manage wallets.
package wallet

import (
	"errors"
	"fmt"

	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

var (
	ErrMnemonic    = errors.New("mnemonic error")
	ErrMasterHDKey = errors.New("master hd key")
)

// Mnemonic takes a bitsize and returns mnemonics if no error
func Mnemonic(bitsize int) (string, error) {
	e, err := bip39.NewEntropy(bitsize)
	if err != nil {
		return "", fmt.Errorf("%w-%s", ErrMnemonic, err.Error())
	}
	m, err := bip39.NewMnemonic(e)
	if err != nil {
		return "", fmt.Errorf("%w-%s", ErrMnemonic, err.Error())
	}
	return m, nil
}

// MasterHDKey takes mnemonic and passphrase, and returns a master
// private key if no error
func MasterHDKey(mnemonic, passphrase string) (*bip32.Key, error) {
	seed := bip39.NewSeed(mnemonic, passphrase)
	key, err := bip32.NewMasterKey(seed)
	if err != nil {
		return nil, fmt.Errorf("%w-%s", ErrMasterHDKey, err.Error())
	}
	return key, nil
}
