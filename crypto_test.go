package faker

import (
	// "fmt"
	// "strings"
	"testing"
)

func TestBitcoinAddress(t *testing.T) {
	c := New().Crypto()
	addr := c.BitcoinAddress()
	t.Log(addr)
	panic(addr)
}

// func TestP2PKHAddress(t *testing.T) {
// 	c := New().Crypto()
// 	addr := c.P2PKHAddress()
// 	Expect(t, true, len(addr) >= bitcoinMin)
// 	Expect(t, true, len(addr) <= bitcoinMax)
// 	Expect(t, true, strings.HasPrefix(addr, validBitcoinPrefix["p2pkh"]))
// 	for i := 0; i < len(bannedBitcoin); i++ {
// 		Expect(t, true, !strings.Contains(addr, bannedBitcoin[i]))
// 	}
// }

// func TestP2PKHAddressWithLength(t *testing.T) {
// 	c := New().Crypto()
// 	length := c.Faker.IntBetween(26, 62)
// 	addr := c.P2PKHAddressWithLength(length)
// 	Expect(t, true, len(addr) == length)
// 	Expect(t, true, strings.HasPrefix(addr, validBitcoinPrefix["p2pkh"]))
// }

// func TestP2SHAddress(t *testing.T) {
// 	c := New().Crypto()
// 	addr := c.P2SHAddress()
// 	Expect(t, true, len(addr) >= bitcoinMin)
// 	Expect(t, true, len(addr) <= bitcoinMax)
// 	Expect(t, true, strings.HasPrefix(addr, validBitcoinPrefix["p2sh"]))
// 	for i := 0; i < len(bannedBitcoin); i++ {
// 		Expect(t, true, !strings.Contains(addr, bannedBitcoin[i]))
// 	}
// }

// func TestP2SHAddressWithLength(t *testing.T) {
// 	c := New().Crypto()
// 	length := c.Faker.IntBetween(26, 62)
// 	addr := c.P2SHAddressWithLength(length)
// 	Expect(t, true, len(addr) == length)
// 	Expect(t, true, strings.HasPrefix(addr, validBitcoinPrefix["p2sh"]))
// }

// func TestBech32Address(t *testing.T) {
// 	c := New().Crypto()
// 	addr := c.Bech32Address()
// 	Expect(t, true, len(addr) >= bitcoinMin)
// 	Expect(t, true, len(addr) <= bitcoinMax)
// 	Expect(t, true, strings.HasPrefix(addr, validBitcoinPrefix["bech32"]))
// 	for i := 0; i < len(bannedBitcoin); i++ {
// 		Expect(t, true, !strings.Contains(addr, bannedBitcoin[i]))
// 	}
// }

// func TestBech32AddressWithLength(t *testing.T) {
// 	c := New().Crypto()
// 	length := c.Faker.IntBetween(26, 62)
// 	addr := c.Bech32AddressWithLength(length)
// 	Expect(t, true, len(addr) == length)
// 	Expect(t, true, strings.HasPrefix(addr, validBitcoinPrefix["bech32"]))
// }

// func TestEtheriumAddress(t *testing.T) {
// 	c := New().Crypto()
// 	addr := c.EtheriumAddress()
// 	Expect(t, true, len(addr) == ethLen)
// 	Expect(t, true, strings.HasPrefix(addr, ethPrefix))
// }

// func TestAlgorithmRange(t *testing.T) {
// 	for k, tc := range []TestCaseAlnum{
// 		{
// 			// The Description of the test case
// 			desc:     "Test Get Digit 0-9",
// 			localInt: 0,
// 			// Our anticipated result
// 			assert: func(t *testing.T, a int, b int) {
// 				Expect(t, true, a == int('0'))
// 				Expect(t, true, b == int('9'))
// 			},
// 		},
// 		{
// 			desc:     "Test Get Uppercase A-Z",
// 			localInt: 1,
// 			assert: func(t *testing.T, a int, b int) {
// 				Expect(t, true, a == int('A'))
// 				Expect(t, true, b == int('Z'))
// 			},
// 		},
// 		{
// 			desc:     "Test Get Lowercase a-z",
// 			localInt: 2,
// 			assert: func(t *testing.T, a int, b int) {
// 				Expect(t, true, a == int('a'))
// 				Expect(t, true, b == int('z'))
// 			},
// 		},
// 	} {
// 		t.Run(fmt.Sprintf("case=%d/description=%s", k, tc.desc), func(t *testing.T) {
// 			// Use our mock here instead of using a seed.
// 			gen := GeneratorMock{}
// 			gen.local = tc.localInt
// 			// populate the generator with our mock as it is an interface.
// 			c := Faker{Generator: gen}
// 			a, b := c.Crypto().algorithmRange()
// 			tc.assert(t, a, b)
// 		})
// 	}
// }

// func TestBitcoinAddressGeneration(t *testing.T) {
// 	for k, tc := range []BitcoinAddressTestCase{
// 		{
// 			// The Description of the test case
// 			desc:           "TestBech32Address",
// 			addressType:    0,
// 			// Our anticipated result
// 			expectedPrefix: "bc1",
// 		},
// 		{
// 			// The Description of the test case
// 			desc:           "TestP2SHAddress",
// 			addressType:    1,
// 			// Our anticipated result
// 			expectedPrefix: "3",
// 		},
// 		{
// 			// The Description of the test case
// 			desc:           "TestP2PKHAddress",
// 			addressType:    2,
// 			// Our anticipated result
// 			expectedPrefix: "1",
// 		},
// 	} {
// 		t.Run(fmt.Sprintf("case=%d/description=%s", k, tc.desc), func(t *testing.T) {
// 			// Use our mock here instead of using a seed.
// 			gen := GeneratorMock{}
// 			gen.local = tc.addressType
// 			// populate the generator with our mock as it is an interface.
// 			c := Faker{Generator: gen}
// 			rs := c.Crypto().BitcoinAddress()
// 			Expect(t, true, strings.HasPrefix(rs, tc.expectedPrefix))
// 			Expect(t, true, len(rs) >= bitcoinMin)
// 			Expect(t, true, len(rs) <= bitcoinMax)

// 			// check if result is a valid base58 string
// 			addr := rs[len(tc.expectedPrefix):]
// 			_, _, err := base58CheckDecode(addr)
// 			t.Log(rs)
// 			t.Log(addr)
// 			t.Log(err)
// 			if err != nil {
// 				t.Errorf("Expected valid base58 string, got %s", err)
// 			}
// 			Expect(t, nil, err)
// 		})
// 	}
// }

// type BitcoinAddressTestCase struct {
// 	desc           string
// 	addressType    int
// 	expectedPrefix string
// }