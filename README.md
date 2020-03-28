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
1. Walk through path in folder files one by one as []byte data
2. Ask for pass phrase 
3. Encrypt pass phrase to 32 byte slice with sha256 checksum
4. Encrypt data with AES and GCM (Galois/Counter Mode) as operation mode
5. Add sha256 checksum of file to the additional data of aes GCM ciphertext

To-Do:
- Compare function to new file
- File reader to new file + move into package
- Documentation
- Manual
- Error log
- Session log

Notes: 
- File name needs to be passed to aditional data if necessary



# Manual:

1. Copy the the executalbe and the folder "files" in to a separate folder.
2. Copy some files or any kind of data into the folder "files". Make sure that you only use copies of the files/data. Since this is a test there is no guaranty that you data takes no damage.
3. After several runs with different files/data check if the data in the folder "decrypted" has taken any damage.
4. Please sent the files in the folder "logs" to e.kusowenko@gmail.com

!!!Thank you for the participation in this test!!!