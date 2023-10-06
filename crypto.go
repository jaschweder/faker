package faker

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"crypto/sha256"
)

type Crypto struct {
	Faker *Faker
}

// BitcoinAddress returns a random Bitcoin address of either Bech32, P2PKH, or P2SH type.
func (c Crypto) BitcoinAddress() string {
	// Generate a random byte slice of length 20
	bytes := make([]byte, 20)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}

	// Choose a random address type
	types := []string{"bech32", "p2pkh", "p2sh"}
	addressType := types[c.Faker.IntBetween(0, len(types)-1)]

	// Generate the address based on the chosen type
	switch addressType {
	case "bech32":
		return c.bech32Address(bytes)
	case "p2pkh":
		return c.p2pkhAddress(bytes)
	case "p2sh":
		return c.p2shAddress(bytes)
	default:
		panic(fmt.Sprintf("Invalid address type: %s", addressType))
	}
}

// doubleSha256 calculates the SHA-256 hash of the input, and then calculates the SHA-256 hash of the result.
func doubleSha256(input []byte) []byte {
	hasher := sha256.New()
	hasher.Write(input)
	hash := hasher.Sum(nil)

	hasher.Reset()
	hasher.Write(hash)
	return hasher.Sum(nil)
}

// base58Encode encodes the input byte slice as a base58 string.
func base58Encode(input []byte) string {
	alphabet := "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

	// Convert the input bytes to a big integer
	num := new(big.Int).SetBytes(input)

	// Encode the big integer as a base58 string
	var result []byte
	for num.Sign() > 0 {
		mod := new(big.Int)
		num.DivMod(num, big.NewInt(58), mod)
		result = append([]byte{alphabet[mod.Int64()]}, result...)
	}

	// Add leading zeros for each zero byte in the input
	for _, b := range input {
		if b != 0x00 {
			break
		}
		result = append([]byte{alphabet[0]}, result...)
	}

	return string(result)
}

// p2pkhAddress generates a P2PKH bitcoin address.
func (c Crypto) p2pkhAddress(bytes []byte) string {
	// Prepend the version byte (0x00) to the public key hash
	versionedBytes := append([]byte{0x00}, bytes...)

	// Calculate the checksum by hashing the versioned bytes twice
	checksum := doubleSha256(versionedBytes)[:4]

	// Concatenate the versioned bytes and the checksum
	addressBytes := append(versionedBytes, checksum...)

	// Encode the address bytes as a base58 string
	return base58Encode(addressBytes)
}

// P2PKHAddress generates a P2PKH bitcoin address.
func (c Crypto) P2PKHAddress() string {
	return c.p2pkhAddress(c.Faker.RandomBytes(20))
}

func (c Crypto) p2shAddress(bytes []byte) string {
	// Prepend the version byte (0x05) to the redeem script hash
	versionedBytes := append([]byte{0x05}, bytes...)

	// Calculate the checksum by hashing the versioned bytes twice
	checksum := doubleSha256(versionedBytes)[:4]

	// Concatenate the versioned bytes and the checksum
	addressBytes := append(versionedBytes, checksum...)

	// Encode the address bytes as a base58 string
	return base58Encode(addressBytes)
}

// P2SHAddress generates a P2SH bitcoin address.
func (c Crypto) P2SHAddress() string {
	return c.p2shAddress(c.Faker.RandomBytes(20))
}

// bech32Address generates a Bech32 bitcoin address
func (c Crypto) bech32Address(bytes []byte) string {
	// Define the Bech32 character set
	const charset = "qpzry9x8gf2tvdw0s3jn54khce6mua7l"

	// Convert the input bytes to a 5-bit binary representation
	var bits []uint5
	for _, b := range bytes {
		bits = append(bits, uint5(b>>5), uint5(b&0x1f))
	}

	// Calculate the checksum
	checksum := bech32Checksum(bits)

	// Concatenate the input bytes and the checksum
	allBits := append(bits, checksum...)

	// Convert the 5-bit binary representation to a Bech32 string
	var bech32Chars []rune
	for _, b := range allBits {
		bech32Chars = append(bech32Chars, rune(charset[b]))
	}

	// Return the Bech32 string
	return "bc1" + string(bech32Chars)
}

// bech32Checksum calculates the Bech32 checksum for the given 5-bit binary representation.
func bech32Checksum(bits []uint5) []uint5 {
	// Define the generator polynomial
	const generator = 0x3b6a57b2

	// Append 6 zero bits to the input
	bits = append(bits, make([]uint5, 6)...)

	// Calculate the initial value of the checksum
	var checksum uint32 = 1
	for _, b := range bits {
		// Extract the most significant bit of the checksum
		msb := checksum >> 25

		// Shift the checksum left by 1 bit and add the current bit
		checksum = (checksum&0x1ffffff)<<1 | uint32(b)

		// If the most significant bit is 1, XOR the checksum with the generator polynomial
		if msb == 1 {
			checksum ^= generator
		}
	}

	// Convert the checksum to a 5-bit binary representation
	var result []uint5
	for i := 0; i < 6; i++ {
		result = append(result, uint5(checksum>>(25-5*i)))
	}

	return result
}

// uint5 is an unsigned 5-bit integer.
type uint5 uint8

// Bech32Address generates a Bech32 bitcoin address
func (c Crypto) Bech32Address() string {
	return c.bech32Address(c.Faker.RandomBytes(20))
}

// // EtheriumAddress returns a hexadecimal ethereum address of 42 characters.
// func (c Crypto) EtheriumAddress() string {
// }

// // P2PKHAddress generates a P2PKH bitcoin address.
// func (c Crypto) P2PKHAddress() string {
// }

// // P2PKHAddressWithLength generates a P2PKH bitcoin address with specified length.
// func (c Crypto) P2PKHAddressWithLength(length int) string {
// }

// // P2SHAddress generates a P2SH bitcoin address.
// func (c Crypto) P2SHAddress() string {
// }

// // P2SHAddressWithLength generates a P2PKH bitcoin address with specified length.
// func (c Crypto) P2SHAddressWithLength(length int) string {
// }

// // Bech32Address generates a Bech32 bitcoin address
// func (c Crypto) Bech32Address() string {
// }

// // Bech32AddressWithLength generates a Bech32 bitcoin address with specified length.
// func (c Crypto) Bech32AddressWithLength(length int) string {
// }