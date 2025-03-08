# Goavatar Identicon Generator in Go

This package provides a simple way to generate unique, symmetric identicons based on an input string (e.g., an email address or username). It uses an **MD5 hash** to create a deterministic pattern and color scheme, then mirrors the design for a visually appealing avatar.

## Installation

To use this package in your Go project, install it via:

```sh
go get github.com/MuhammadSaim/goavatar
```

Then, import it in your Go code:

```go
import "github.com/MuhammadSaim/goavatar"
```


## Usage

### **Basic Example**

```go
package main

import (
    "github.com/MuhammadSaim/goavatar"
)

func main() {
    goavatar.Make("example@example.com", "avatar.png")
}
```

This will generate a unique identicon for the input string and save it as `avatar.png`.

## Package Documentation

### **Generate Identicon**

```go
func Make(input, filename string)
```

- `input`: A string used to generate a unique identicon (e.g., email, username).
- `filename`: The name of the output image file.


## License
This project is open-source under the MIT License.


## Contributing
Contributions are welcome! Feel free to open a pull request or create an issue.

## Author
👤 **Muhammad Saim**
