package main

import (
	"io/ioutil"
	"mutant/object"
	"mutant/mutil"
	"mutant/runner"
	"os"
	"fmt"
)

func main() {
	signedCode, _ := ioutil.ReadFile(os.Args[1])
	byteCode, _ := runner.Decode(signedCode)

	for i := 0; i < len(byteCode.Constants); i++ {
		obfuscatedObj := byteCode.Constants[i]
		deobfuscatedObj, _ := mutil.DecryptObject(obfuscatedObj, len(byteCode.Instructions))
		
		if(deobfuscatedObj.Type() == object.STRING_OBJ) {
			fmt.Println(deobfuscatedObj.Inspect());
		}
	}
}
