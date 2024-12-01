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

func IsAmulet(text string) (bool, int) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))

	result := int(C.check_amulet(cstr))
	return result >= 4, result
}
