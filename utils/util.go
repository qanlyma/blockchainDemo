//util.go

package utils

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"math/big"
	"os"
	"qanly_chain/constcoe"

	"github.com/mr-tron/base58"
	"golang.org/x/crypto/ripemd160"
)

func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func ToHexInt(num int64) []byte { //将int64转换为字节串类型
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

func FileExists(fileAddr string) bool {
	if _, err := os.Stat(fileAddr); os.IsNotExist(err) {
		return false
	}
	return true
}

func PublicKeyHash(publicKey []byte) []byte { //将公钥转为公钥哈希的函数
	hashedPublicKey := sha256.Sum256(publicKey)
	hasher := ripemd160.New()
	_, err := hasher.Write(hashedPublicKey[:])
	Handle(err)
	publicRipeMd := hasher.Sum(nil)
	return publicRipeMd
}

func CheckSum(ripeMdHash []byte) []byte { //增加检查位生成函数
	firstHash := sha256.Sum256(ripeMdHash)
	secondHash := sha256.Sum256(firstHash[:])
	return secondHash[:constcoe.ChecksumLength]
}

func Base58Encode(input []byte) []byte { //Base256转Base58函数
	encode := base58.Encode(input)
	return []byte(encode)
}

func Base58Decode(input []byte) []byte {
	decode, err := base58.Decode(string(input[:]))
	Handle(err)
	return decode
}

func PubHash2Address(pubKeyHash []byte) []byte { //公钥哈希生成钱包地址
	networkVersionedHash := append([]byte{constcoe.NetworkVersion}, pubKeyHash...)
	checkSum := CheckSum(networkVersionedHash)
	finalHash := append(networkVersionedHash, checkSum...)
	address := Base58Encode(finalHash)
	return address
}

func Address2PubHash(address []byte) []byte { //钱包地址转公钥哈希
	pubKeyHash := Base58Decode(address)
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-constcoe.ChecksumLength]
	return pubKeyHash
}

func Sign(msg []byte, privKey ecdsa.PrivateKey) []byte {
	r, s, err := ecdsa.Sign(rand.Reader, &privKey, msg)
	Handle(err)
	signature := append(r.Bytes(), s.Bytes()...)
	return signature
}

func Verify(msg []byte, pubkey []byte, signature []byte) bool {
	curve := elliptic.P256()
	r := big.Int{}
	s := big.Int{}
	sigLen := len(signature)
	r.SetBytes(signature[:(sigLen / 2)])
	s.SetBytes(signature[(sigLen / 2):])

	x := big.Int{}
	y := big.Int{}
	keyLen := len(pubkey)
	x.SetBytes(pubkey[:(keyLen / 2)])
	y.SetBytes(pubkey[(keyLen / 2):])

	rawPubKey := ecdsa.PublicKey{curve, &x, &y}
	return ecdsa.Verify(&rawPubKey, msg, &r, &s)
}
