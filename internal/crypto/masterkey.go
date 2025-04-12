package crypto

import (
	"bufio"
	"crypto/sha256"
	"crypto/subtle"
	"errors"
	"golang.org/x/crypto/pbkdf2"
	"io/ioutil"
	"os"
	"strings"
)

var (
	pepper   = "PoGoPassPepper!" // Replace this BEFORE you run the app, Or replace the values in .env
	salt     = "PoGoPassSalt1234" // Replace this BEFORE you run the app, Or replace the values in .env
	hashFile = "masterkey.hash"
)

func init() {
	loadEnv(".env")

	// Override defaults if .env has them
	if v := os.Getenv("PEPPER"); v != "" {
		pepper = v
	}
	if v := os.Getenv("SALT"); v != "" {
		salt = v
	}
	if v := os.Getenv("HASH_FILE"); v != "" {
		hashFile = v
	}
}

func loadEnv(path string) {
	file, err := os.Open(path)
	if err != nil {
		return // silently ignore if .env is missing
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Skip comments and empty lines
		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue
		}

		// Parse key=value
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])
		os.Setenv(key, val)
	}
}

// DeriveMasterKey derives the master key using PBKDF2
func DeriveMasterKey(password string) ([]byte, error) {
	peppered := password + pepper
	return pbkdf2.Key([]byte(peppered), []byte(salt), 100000, 32, sha256.New), nil
}

// SaveMasterKeyHash saves the hashed master key to a file
func SaveMasterKeyHash(hash []byte) error {
	err := ioutil.WriteFile(hashFile, hash, 0600)
	if err != nil {
		return errors.New("failed to save hash: " + err.Error())
	}
	return nil
}

// LoadMasterKeyHash loads the hashed master key from a file
func LoadMasterKeyHash() ([]byte, error) {
	hash, err := ioutil.ReadFile(hashFile)
	if err != nil {
		return nil, errors.New("failed to load hash: " + err.Error())
	}
	return hash, nil
}

// VerifyMasterKey verifies the entered password against the saved hash
func VerifyMasterKey(password string) bool {
	storedHash, err := LoadMasterKeyHash()
	if err != nil {
		return false
	}

	derivedHash, err := DeriveMasterKey(password)
	if err != nil {
		return false
	}

	return subtle.ConstantTimeCompare(derivedHash, storedHash) == 1
}
