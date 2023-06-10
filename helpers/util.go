package helpers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func SanitizeNumber(word string) string {
	re, err := regexp.Compile(`[^\0-9]`)
	if err != nil {
		log.Fatal(err)
	}
	afterSenitize := re.ReplaceAllString(word, "")
	return afterSenitize
}

func PublicKey() string {
	var pemPublicKey = `-----BEGIN PUBLIC KEY-----
	MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA5G9w1Pb9SSXcjqJNGq/h
	It7JwTUOItrXH+KxyIinfCYFIH6alHwM+UYZ7pQlwFQZMoM+Iy4gx1t/MtXRSiaj
	OSTW94F2QZP14HrRsiekFW1rkcX2r2Upt539hUOH/wSIVx5/u9qnn3yVT/SZzRPT
	80gULImfKkWLYCV/MnTw2UMYAst+gKt/RNvczO1veg+Y+1A3p/wpbkeNaopqKUjf
	/KCiQvVl+r0RkqLfNsgNqoggaQ/KZV3dXaFnE+bGgmpzjKN5cjlocDt0vednnvCL
	1UK2ZFjaFIFCzii28H4ONnxHpWNmUUxyuLXJvqgsp7qS59j6XkjAxVt15zMTlBzi
	1wIDAQAB
-----END PUBLIC KEY-----`
	return pemPublicKey
}
func PrivateKey() string {
	var pemPrivateString = `-----BEGIN PRIVATE KEY-----
	MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDkb3DU9v1JJdyO
	ok0ar+Ei3snBNQ4i2tcf4rHIiKd8JgUgfpqUfAz5RhnulCXAVBkygz4jLiDHW38y
	1dFKJqM5JNb3gXZBk/XgetGyJ6QVbWuRxfavZSm3nf2FQ4f/BIhXHn+72qeffJVP
	9JnNE9PzSBQsiZ8qRYtgJX8ydPDZQxgCy36Aq39E29zM7W96D5j7UDen/CluR41q
	imopSN/8oKJC9WX6vRGSot82yA2qiCBpD8plXd1doWcT5saCanOMo3lyOWhwO3S9
	52ee8IvVQrZkWNoUgULOKLbwfg42fEelY2ZRTHK4tcm+qCynupLn2PpeSMDFW3Xn
	MxOUHOLXAgMBAAECggEAaeOKjv2KvVySl39+dE9w8gQJy8i3K8r7i2k+9fD6ih7p
	o31sVEYIkYhAPwpnUXbqUzLpG8+nHCI6nSrmIBQ29ycvin11fsKCaDMmfwnHErOs
	+F6mkfk31EilGyAJq1nDhXa6yS57Iv/SCsUcgiadyhjwWRDWOfcQu1nGU3JHrr19
	2KnyNGFegIItJtiiqVz6tmQsEB8N2xWoLI85HaTQ51l1nfoebT+8KlIycVhqoR2J
	X+SP5EHsDVw4R6t3XK1d4rH05qrjSYybjULdFXPGvFc+IcYZ6QOKCVSIsq3YJUV0
	aj609aY9kqBGUqozXAfRsOwjt/nUl63g0vXZ2GmOkQKBgQD5Q/jEP3KIh5BidVZA
	JewceHoS1+8+Dp8HnLWJUpaGz4n8Sk2p+8j+oevlkmkkJub9Gt094XfYvkGbHb7B
	a0zNk825D23prD7lqyFXjZmkr66eFnTukCXG7nz/6/67og3NglScQCoE3DcV+4Iz
	noN/cUxLi9eTQCPoWyh5l8L4/wKBgQDqm2X4r821IyLlnbZeU0pgMZO+oD5rSLUk
	2dyd7AaSAWz/sEVFeAVcyNXyJ4lGNrdJvJ+twJlOPMK5C00uoAm2yqG6NlHUDwdc
	fOgau8YvxutrLIsPrydWhLtHc8/RSzJr4Zae0TTfqthW20UimpPD+6UftRyo76LS
	ocTFByr+KQKBgFXJUeVgnK9mUIfCMEP1iTQnNoQjst/dslexVD0Fom6VIL0maWI2
	GG+iFIi3Ad6CUP8M7tWsMk3y9KtI6myw3AbodmXZbI9+S0tJwTjbr+Qg3mzj96xf
	CdFUJMsDUnELDcsLrsjzwEJZ889p9t6DEGic+pAJedDgwzrlnKF0XJLLAoGBAJJd
	pw2y6LykkiX88gUBI7rF024vXSHjt5epEBm6YhL/LriKiX0gtv+/ELNF9T/H7Svk
	sR5etYZ5I+b8ZQe8srLG0oVxVDXftnD+QHRFSA0Qplkz7gI3/Wvd3VVjrHjf2DI0
	CJtG3Bza4qO1ovlGxP+VZNxWSu4eq0+Lu05M/YaZAoGBAK7pWxFTNE6tO+OdMtV9
	7DjqPwOraWFWxiw7SXHnX3j6Srb+QHxhsoGsSMwo3ldp/Y3ENxZE8u0+G3x2mZGI
	nParDl/dG/+eqM0Nf5d48UGoNXjdVUqUdXdornLRRH4aHJqziLFRGDua4vstgfHp
	+nm79ciTVItms+SDUuFu8su6
-----END PRIVATE KEY-----`
	return pemPrivateString
}

func EncryptRSAKey(word string) string {
	pembytes := PublicKey()
	data, _ := pem.Decode([]byte(pembytes))
	publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	key := publicKeyImported.(*rsa.PublicKey)
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, key, []byte(word))
	if err != nil {
		fmt.Println("Error encrypting message:", err)
		os.Exit(1)
	}
	ct := base64.StdEncoding.EncodeToString(ciphertext)
	return ct
}
func DecryptRSAKey(words string) string {
	pembytePrvt := PrivateKey()
	dataPrvt, _ := pem.Decode([]byte(pembytePrvt))
	privateKeyImported, err := x509.ParsePKCS8PrivateKey(dataPrvt.Bytes)
	if err != nil {
		fmt.Println(err)
	}
	key := privateKeyImported.(*rsa.PrivateKey)
	ct, _ := base64.StdEncoding.DecodeString(words)
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, key, []byte(ct))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Plaintext:")
	fmt.Println(string(plaintext))
	return string(plaintext)
}
func OneWayEncrypt(message []byte) string {
	hash_byte := sha256.Sum256(message)
	hash_str := hex.EncodeToString(hash_byte[:])
	return hash_str
}

func SeparateWord(word string) string {
	var mergeArr []string
	var merge string
	var wordMerge string
	separate := strings.Split(word, " ")
	if len(separate) > 1 {
		for _, item := range separate {
			merge = item[0:1]
			mergeArr = append(mergeArr, merge)
		}

		wordMerge = strings.Join(mergeArr, "")
	} else {
		wordMerge = word[0:1]
	}
	return wordMerge
}

func GenerateRandomNumber(len int) (int, error) {
	maxLimit := int64(int(math.Pow10(len)) - 1)
	lowLimit := int(math.Pow10(len - 1))

	randomNumber, err := rand.Int(rand.Reader, big.NewInt(maxLimit))
	if err != nil {
		return 0, err
	}
	randomNumberInt := int(randomNumber.Int64())

	// Handling integers between 0, 10^(n-1) .. for n=4, handling cases between (0, 999)
	if randomNumberInt <= lowLimit {
		randomNumberInt += lowLimit
	}

	// Never likely to occur, kust for safe side.
	if randomNumberInt > int(maxLimit) {
		randomNumberInt = int(maxLimit)
	}
	return randomNumberInt, nil
}

func IntegerToRoman(numberStr string) string {
	number, _ := strconv.Atoi(numberStr)
	romanMap := map[int]string{
		1: "I", 4: "IV", 5: "V", 9: "IX", 10: "X", 40: "XL", 50: "L",
		90: "XC", 100: "C", 400: "CD", 500: "D", 900: "CM", 1000: "M",
	}
	// create a slice of slices
	rows := len(romanMap)
	matrix := make([][]string, rows)
	var key_slice []int
	for k, _ := range romanMap {
		key_slice = append(key_slice, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(key_slice)))
	row := 0
	for _, key := range key_slice {
		// convert int key to string key
		skey := strconv.Itoa(key)
		matrix[row] = []string{skey, romanMap[key]}
		row++

	}
	result := ""
	for _, item := range matrix {
		// convert string to int
		den, err := strconv.Atoi(item[0])
		if err != nil {
			panic(err)
		}
		sym := item[1]
		for number >= den {
			result += sym
			number -= den
		}
	}
	return result
}
