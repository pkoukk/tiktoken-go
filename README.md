# tiktoken-go
[简体中文](./README_zh-hans.md)

OpenAI's tiktoken in Go.  
Tiktoken is a fast BPE tokeniser for use with OpenAI's models.  
This is a port of the original [tiktoken](https://github.com/openai/tiktoken).  

# Usage

## Install

```bash
go get github.com/zhenghaoz/tiktoken-go
```

## Example

### get token by encoding

```go
package main

import (
    "fmt"
    "github.com/zhenghaoz/tiktoken-go"
)

func main() (num_tokens int) {
    text = "Hello, world!"
    encoding = "gpt2"

	tke, err := tiktoken.GetEncoding(encoding)
	if err != nil {
		err = fmt.Errorf("GetEncoding: %v", err)
		return
	}

    // encode
	token := tke.Encode(text, nil, nil)

    // num_tokens
    num_tokens = len(token)
}
```

### get token by Model

```go
package main

import (
    "fmt"
    "github.com/zhenghaoz/tiktoken-go"
)

func main() (num_tokens int) {
    text = "Hello, world!"
    encoding = "davinci"

   tkm, err := tiktoken.EncodingForModel(model)
	if err != nil {
		err = fmt.Errorf(": %v", err)
		return
	}

	 // encode
	token := tke.Encode(text, nil, nil)

    // num_tokens
    num_tokens = len(token)
}
```

# compare with original [tiktoken](https://github.com/openai/tiktoken)

## get token by encoding  

| python | go |
| --- | --- |
|


# License
[MIT](./LICENSE)