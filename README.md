# CryptoCloud

This program reads the dir it is in, fetches the file names and writes them to a .txt. 

Useful links: 
https://gist.github.com/josephspurrier/12cc5ed76d2228a41ceb
https://eli.thegreenplace.net/2019/aes-encryption-of-files-in-go/
https://tutorialedge.net/golang/go-encrypt-decrypt-aes-tutorial/

https://golang.org/pkg/crypto/cipher/

https://en.wikipedia.org/wiki/SEAL_(cipher)
https://de.wikipedia.org/wiki/Galois/Counter_Mode
https://de.wikipedia.org/wiki/Substitutions-Permutations-Netzwerk
https://de.wikipedia.org/wiki/Betriebsmodus_(Kryptographie)
https://de.wikipedia.org/wiki/Advanced_Encryption_Standard

Process: 
1. Read files in dir one by one as []byte data
2. Ask for pass phrase 
3. Encrypt pass phrase to 32 byte slice with sha256 checksum
4. Encrypt data with AES and GCM (Galois/Counter Mode) as operation mode