package main

import "os"
import "fmt"
import "strings"

func isAlpha(plainChar byte) bool {

	if plainChar >= 65 && plainChar <= 90 {
		return true
	}
	return false

}

func getEncryptedChar(plainChar byte, keyChar byte) string {

	encryptedChar := ((plainChar - 65) + (keyChar - 65)) % 26

	return string(encryptedChar + 65)
}

func getDecryptedChar(cipherChar int, keyChar int) string {

	decryptedChar := (cipherChar - keyChar) % 26

	for decryptedChar < 0 {
		decryptedChar += 26
	}

	return string(decryptedChar + 65)
}

func vigenere(key string, inputtext string, encrypt bool) string {

	resultText := ""
	keyIndex := 0
	count := 0

	keyUpper := strings.ToUpper(key)
	inputUpper := strings.ToUpper(inputtext)

	for i := 0; i < len(inputUpper); i++ {

		if isAlpha(inputUpper[i]) {

			keyIndex = count % len(keyUpper)

			if encrypt == true {
				resultText += getEncryptedChar(inputUpper[i], keyUpper[keyIndex])
			} else {
				resultText += getDecryptedChar(int(inputUpper[i]), int(keyUpper[keyIndex]))
			}
			count++

		} else {

			resultText += string(inputUpper[i])

		}

	}

	return resultText

}
