package Plugins

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"net"
)

var PluginList = map[string]interface{}{
	"21":      FtpScan,
	"22":      SshScan,
	"135":     Findnet,
	"139":     NetBIOS,
	"445":     SmbScan,
	"1433":    MssqlScan,
	"1521":    OracleScan,
	"3306":    MysqlScan,
	"3389":    RdpScan,
	"5432":    PostgresScan,
	"6379":    RedisScan,
	"9000":    FcgiScan,
	"11211":   MemcachedScan,
	"27017":   MongodbScan,
	"1000001": MS17010,
	"1000002": SmbGhost,
	"1000003": WebTitle,
	"1000004": SmbScan2,
	"1000005": WmiExec,
}

func ReadBytes(conn net.Conn) (result []byte, err error) {
	size := 4096
	buf := make([]byte, size)
	for {
		count, err := conn.Read(buf)
		if err != nil {
			break
		}
		result = append(result, buf[0:count]...)
		if count < size {
			break
		}
	}
	if len(result) > 0 {
		err = nil
	}
	return result, err
}

var key = "0123456789abcdef"

func AesEncrypt(orig string, key string) string {
	origData := []byte(orig)
	k := []byte(key)
	block, _ := aes.NewCipher(k)
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	cryted := make([]byte, len(origData))
	blockMode.CryptBlocks(cryted, origData)
	return base64.StdEncoding.EncodeToString(cryted)
}
func AesDecrypt(cryted string, key string) string {
	crytedByte, _ := base64.StdEncoding.DecodeString(cryted)
	k := []byte(key)
	block, _ := aes.NewCipher(k)
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	orig := make([]byte, len(crytedByte))
	blockMode.CryptBlocks(orig, crytedByte)
	orig = PKCS7UnPadding(orig)
	return string(orig)
}

// The block size for AES encryption must be 128 bits (byte[16]),
// and the key length can be any of the following: 128 bits (byte[16]), 192 bits (byte[24]), or 256 bits (byte[32]).
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
