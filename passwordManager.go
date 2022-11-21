package main


//make a vault.aes file that is decrypted only in memory


//use aes-256 to encrypt data

//make a table
//  -show all passwords 
//  -add password
//  -make  encrypted text  comma seporated



import (
	"fmt"
	"encoding/hex"
	"log"
	"os"
	"crypto/aes"
	"crypto/sha256"
	"crypto/cipher"
    "crypto/rand"
	"strings"
	"io"
)

//a function to handle errora
func check(err error) {
	if err != nil{
		log.Fatal(err)
	}
}

//this will parse the string
func showPasswords(text string) {
	sepratedText := strings.Split(text, ",")

	for _, y := range sepratedText{
		fmt.Println(y)
	}
}

//aes encrypt the data(string) and retrun a byte array
func encrypt(data string, key string) string{

	//make  the key a byte array
	bytsKey := []byte(key)

	//def the cipfer
	ciph, err := aes.NewCipher(bytsKey)
	check(err)

	gcm, err := cipher.NewGCM(ciph)

	nonce := make([]byte, gcm.NonceSize())
    // populates our nonce with a cryptographically secure
    // random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
    	fmt.Println(err)
	}

	return string(gcm.Seal(nonce, nonce, []byte(data), nil))
}

//decrypt the data(byte array) and return a string
func decrypt(data string, key string) string{


	//ciphertext, _ := hex.DecodeString(data)
 
    c, err := aes.NewCipher([]byte(key))
    check(err)
 
    /*
    pt := make([]byte, len(data))
    c.Decrypt(pt, ciphertext)
 
    s := string(pt[:])
    return s
    */


    gmc, err := cipher.NewGCM(c)
    check(err)

    nonceSize := gmc.NonceSize()
    if len(data) < nonceSize {
        fmt.Println(err)
    }

    nonce, ciphertext := []byte(data)[:nonceSize], []byte(data)[nonceSize:]
    plaintext, err := gmc.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        fmt.Println(err)
    }
	return string(plaintext)
}

//this will read the file
func getFileContents() string{
	//this will read the file
	data, err := os.ReadFile("vault.aes")
	check(err)

	return string(data)
}

//this will add a new password the the csv string
func addPassword() string{
	var username string
	var password string
	var site string

	fmt.Print("username: ")
	fmt.Scanln(&username)

	fmt.Print("password: ")
	fmt.Scanln(&password)

	fmt.Print("site: ")
	fmt.Scanln(&site)

	master := username + " " + password + " " + site + ","
	return master
} 

//this will make a new file with the new encrypted data
func makeNewVault(text string) {
	file, err := os.Create("vault.aes")
	check(err)

	//convert the string to bytes
	byteData := []byte(text)

	//write the data to the file and close it
	file.Write(byteData)
	file.Close()
}



func main(){

	const masterPassword string  = "d35ed0ad3351706a7810e0c9dbd4b006e532a79f601e814e843056db0952556e"//thisis32bitlongpassphraseimusing

	var password string
	var attempt int = 0
	for {
		if attempt == 3{
			os.Exit(3)
		} else {
			fmt.Print("Password: ")
			fmt.Scanln(&password)

			h := sha256.New()
			h.Write([]byte(password))
			bs := h.Sum(nil)

			if string(hex.EncodeToString(bs)) == masterPassword {
				fmt.Println("fuck")
				break
			}else {
				attempt++
			}
		}
	}


	for {
		fmt.Println(`

			passwords are decrypted in memory

[0] show passwords
[1] add password
[2] make new vault
[3] default
			`)

		fmt.Print(": ")
		var option int

		fmt.Scanln(&option)


		switch (option){
			case 0:
				data := decrypt(getFileContents(), password)
				showPasswords(data)
				continue
			case 1:
				//decrypt the data
				data := decrypt(getFileContents(), password)
				//make a new password string
				newPass := addPassword()
				//add the new password string to the data 
				newData := encrypt(data + newPass, password)

				//write the new encrypted string to the file
				makeNewVault(newData)
			case 2:
				newData := encrypt("password manager 1.0\n", password)
				makeNewVault(newData)
			case 3:
				os.Exit(3) 
		}
	}
}
