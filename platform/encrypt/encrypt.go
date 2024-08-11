package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
)

// Encrypt is used to encrypt the data using AES-256 algorithm, input: []byte data, output: []byte encrypted data, []byte key, error
func Encrypt(data []byte) ([]byte, []byte, error) {
	key := make([]byte, 32)
	if _, err := rand.Reader.Read(key); err != nil {
		return nil, nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Reader.Read(nonce); err != nil {
		return nil, nil, err
	}
	cipherText := gcm.Seal(nonce, nonce, data, nil)
	fmt.Println("cipherText: ", cipherText, "key: ", key, "data: ", data)
	return cipherText, key, nil
}

// Decrypt is used to decrypt the data using AES-256 algorithm
// input: encrypted data in byte array, key in byte array
// output: decrypted data in byte array
func Decrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	decryptedData, err := gcm.Open(nil, data[:gcm.NonceSize()], data[gcm.NonceSize():], nil)
	if err != nil {
		return nil, err
	}
	return decryptedData, nil
}
