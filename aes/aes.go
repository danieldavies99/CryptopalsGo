package aes

import (
	"crypto/aes"
	"errors"

	"github.com/danieldavies99/cryptopals/xor"
)

func chunkByteSlice(input []byte, chunkSize int) [][]byte {
	output := [][]byte{}
	for i := 0; i < len(input)/chunkSize; i++ {
		output = append(output, input[i*chunkSize:i*chunkSize+chunkSize])
	}
	return output
}

func DecryptAes128Ecb(key []byte, data []byte) ([]byte, error) {
	if len(data)%len(key) != 0 {
		return nil, errors.New("Key length must be multiple of data length")
	}

	cipher, _ := aes.NewCipher(key)
	decrypted := make([]byte, len(data))
	size := 16

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		cipher.Decrypt(decrypted[bs:be], data[bs:be])
	}

	return decrypted, nil
}

func EncryptAes128Ecb(key []byte, data []byte) ([]byte, error) {
	if len(data)%len(key) != 0 {
		return nil, errors.New("Key length must be multiple of data length")
	}

	cipher, _ := aes.NewCipher(key)
	encrypted := make([]byte, len(data))
	size := 16

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		cipher.Encrypt(encrypted[bs:be], data[bs:be])
	}

	return encrypted, nil
}

func DecryptAes128Cbc(
	key []byte,
	IV []byte,
	input []byte,
) ([]byte, error) {
	blockSize := 16
	previousCT := IV

	blocks := chunkByteSlice(input, blockSize)
	output := []byte{}
	for _, block := range blocks {
		decryptedRes, err := DecryptAes128Ecb(key, block)
		if err != nil {
			return nil, err
		}

		xorRes, err := xor.XOR(previousCT, decryptedRes)
		previousCT = block
		output = append(output, xorRes...)
		if err != nil {
			return nil, err
		}
	}

	return output, nil
}

func EncryptAes128Cbc(
	key []byte,
	IV []byte,
	input []byte,
) ([]byte, error) {
	blockSize := 16
	previousCT := IV

	blocks := chunkByteSlice(input, blockSize)
	output := []byte{}
	for _, block := range blocks {
		xorRes, err := xor.XOR(previousCT, block)
		if err != nil {
			return nil, err
		}

		previousCT, err = EncryptAes128Ecb(key, xorRes)

		output = append(output, previousCT...)
		if err != nil {
			return nil, err
		}
	}

	return output, nil
}
