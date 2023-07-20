# hashing

## Installation

```bash
go get github.com/go-fires/hashing
```

## Usage

```go
package main

import (
	"github.com/go-fires/hashing/bcrypt"
	"github.com/go-fires/hashing/md5"
	"github.com/go-fires/hashing/sha1"
)

func main() {
	md5.New().Make("hello")
	md5.New().MustMake("hello")
	md5.New().Check("hello", "68656c6c6fda39a3ee5e6b4b0d3255bfef95601890afd80709")

	bcrypt.New().Make("hello")
	bcrypt.New().MustMake("hello")
	bcrypt.New().Check("hello", "68656c6c6fda39a3ee5e6b4b0d3255bfef95601890afd80709")

	sha1.New().Make("hello")
	sha1.New().MustMake("hello")
	sha1.New().Check("hello", "68656c6c6fda39a3ee5e6b4b0d3255bfef95601890afd80709")
}
```

## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.