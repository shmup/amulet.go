package amulet

/*
#cgo LDFLAGS: -lcrypto
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <openssl/evp.h>

int check_amulet(const char* text) {
    if (strlen(text) > 64) return 0;

    EVP_MD_CTX *mdctx = EVP_MD_CTX_new();
    unsigned char hash[EVP_MAX_MD_SIZE];
    unsigned int hash_len;
    char hex[EVP_MAX_MD_SIZE*2 + 1];

    EVP_DigestInit_ex(mdctx, EVP_sha256(), NULL);
    EVP_DigestUpdate(mdctx, text, strlen(text));
    EVP_DigestFinal_ex(mdctx, hash, &hash_len);
    EVP_MD_CTX_free(mdctx);

    for(int i = 0; i < hash_len; i++) {
        sprintf(hex + (i * 2), "%02x", hash[i]);
    }

    // Look for 4+ consecutive 8s
    for(int i = 0; i < strlen(hex)-3; i++) {
        if(hex[i] == '8') {
            int count = 1;
            while(i+count < strlen(hex) && hex[i+count] == '8') count++;
            if(count >= 4) return count;
        }
    }
    return 0;
}
*/
import "C"
import "unsafe"

// IsAmulet determines if a string is an "amulet" by checking its SHA-256 hash.
// A string is considered an amulet if its hash contains 4 or more consecutive '8' characters.
// The input string must not exceed 64 characters in length.
//
// Example:
//
//	isAmulet, count := IsAmulet("test")
//
// Returns true and the count of consecutive 8s if it's an amulet,
// false and 0 otherwise.
func IsAmulet(text string) (bool, int) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))

	result := int(C.check_amulet(cstr))
	return result >= 4, result
}
