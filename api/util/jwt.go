// Implementation of JWT
// TODO:
//   Implement timeouts
package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const SECRET_ENV = "SEEKERS_GUILD_API_SECRET"

func generateSignature(header, payload string) (string, error) {
    // Get secret
    secret := os.Getenv(SECRET_ENV)
    if secret == "" {
        return "", fmt.Errorf("secret key not defined in env")   
    }

    // Hash header + payload to generate token signature
    hsh := hmac.New(sha256.New, []byte(secret))
    hsh.Write([]byte(header + string(payload)))
    return base64.StdEncoding.EncodeToString(hsh.Sum(nil)), nil
}

func generateToken(claims map[string]string, header string) (string, error) {
    payload, err := json.Marshal(claims)
    if err != nil {
        return "", err
    }

    // Get signature
    sig, err := generateSignature(header, string(payload))
    if err != nil {
        return "", err
    }

    // Base64-encode header and payload
    header64 := base64.StdEncoding.EncodeToString([]byte(header))
    payload64 := base64.StdEncoding.EncodeToString([]byte(payload))

    // Create JWT of form <header>.<payload>.<signature>
    return header64 + "." + payload64 + "." + sig, nil
}

func validateToken(token string) (bool, error) {
    // Split token into components
    splitToken := strings.Split(token, ".")
    if len(splitToken) != 3 {
        return false, nil
    }

    // Decode from base64
    header, err := base64.StdEncoding.DecodeString(splitToken[0])
    if err != nil {
        return false, err
    }
    payload, err := base64.StdEncoding.DecodeString(splitToken[1])
    if err != nil {
        return false, err
    }

    // Get signature
    sig, err := generateSignature(string(header), string(payload))
    if err != nil {
        return false, err
    }

    // Perform comparison
    // TODO: Constant time comparison
    if sig != splitToken[2] {
        return false, nil
    }

    return true, nil
}
