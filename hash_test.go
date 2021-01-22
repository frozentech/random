package random_test

import (
	"fmt"

	"github.com/frozentech/random"
)

func ExampleBlake2B() {
	fmt.Println(random.Blake2B("foobar1234"))
	// Output:
	// 666f6f62617231323334e4a6a0577479b2b4
}

func ExampleHMAC() {
	fmt.Println(random.HMAC([]byte("foobar1234"), []byte("foobar1234")))
	// Output:
	// [226 0 124 72 240 77 240 23 188 138 71 237 130 217 30 165 204 193 125 1 226 165 199 117 7 1 74 233 139 207 212 215]
}

func ExampleMD5() {
	fmt.Println(random.MD5("foobar1234"))
	// Output:
	// 68052bc3a6a843c79bd86621822b06f1
}

func ExampleSHA256() {
	fmt.Println(random.SHA256("foobar1234"))
	// Output:
	// baef1fb2d0dc3a5584b59f99ef3a903891b036c46f1019c393c30c1aeb1e8fcb
}
