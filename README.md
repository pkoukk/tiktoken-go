# tiktoken-go
[简体中文](./README_zh-hans.md)

OpenAI's tiktoken in Go. 

Tiktoken is a fast BPE tokeniser for use with OpenAI's models.

This is a port of the original [tiktoken](https://github.com/openai/tiktoken).  

# Usage

## Install

```bash
go get github.com/pkoukk/tiktoken-go
# default tiktoken need download token dictionary from openai website, 
# if you want use this lib offline, use embed branch instead
go get github.com/pkoukk/tiktoken-go@embed
```
## Cache
Tiktoken-go has the same cache mechanism as the original Tiktoken library.  

You can set the cache directory by using the environment variable TIKTOKEN_CACHE_DIR.   

Once this variable is set, tiktoken-go will use this directory to cache the token dictionary.   

If you don't set this environment variable, tiktoken-go will download the dictionary each time you initialize an encoding for the first time.  

## Example

### get token by encoding

```go
package main

import (
    "fmt"
    "github.com/pkoukk/tiktoken-go"
)

func main()  {
	text := "Hello, world!"
	encoding := "cl100k_base"

	tke, err := tiktoken.GetEncoding(encoding)
	if err != nil {
		err = fmt.Errorf("getEncoding: %v", err)
		return
	}

	// encode
	token := tke.Encode(text, nil, nil)

	//tokens
	fmt.Println((token))
	// num_tokens
	fmt.Println(len(token))
}
```

### get token by Model

```go
package main

import (
    "fmt"
    "github.com/pkoukk/tiktoken-go"
)

func main()  {
	text := "Hello, world!"
	encoding := "gpt-3.5-turbo"

	tkm, err := tiktoken.EncodingForModel(encoding)
	if err != nil {
		err = fmt.Errorf("getEncoding: %v", err)
		return
	}

	// encode
	token := tkm.Encode(text, nil, nil)

	// tokens
	fmt.Println(token)
	// num_tokens
	fmt.Println(len(token))
}
```
### counting tokens for chat API calls
Below is an example function for counting tokens for messages passed to gpt-3.5-turbo-0301 or gpt-4-0314.

The following code was written by @nasa1024 based on [openai-cookbook](https://github.com/openai/openai-cookbook/blob/main/examples/How_to_count_tokens_with_tiktoken.ipynb)     examples.

Please note that the token calculation method for the message may change at any time, so this code may not necessarily be applicable in the future.

If you need accurate calculation, please refer to the official documentation.

If you find that this code is no longer applicable, please feel free to submit a PR or Issue.


```go
package main

import (
	"fmt"

	"github.com/pkoukk/tiktoken-go"
	"github.com/sashabaranov/go-openai"
)

func NumTokensFromMessages(messages []openai.ChatCompletionMessage, model string) (num_tokens int) {
	tkm, err := tiktoken.EncodingForModel(model)
	if err != nil {
		err = fmt.Errorf("EncodingForModel: %v", err)
		fmt.Println(err)
		return
	}

	var tokens_per_message int
	var tokens_per_name int
	if model == "gpt-3.5-turbo-0301" || model == "gpt-3.5-turbo" {
		tokens_per_message = 4
		tokens_per_name = -1
	} else if model == "gpt-4-0314" || model == "gpt-4" {
		tokens_per_message = 3
		tokens_per_name = 1
	} else {
		fmt.Println("Warning: model not found. Using cl100k_base encoding.")
		tokens_per_message = 3
		tokens_per_name = 1
	}

	for _, message := range messages {
		num_tokens += tokens_per_message
		num_tokens += len(tkm.Encode(message.Content, nil, nil))
		num_tokens += len(tkm.Encode(message.Role, nil, nil))
		num_tokens += len(tkm.Encode(message.Name,nil,nil))
		if message.Name != "" {
			num_tokens += tokens_per_name
		}
	}
	num_tokens += 3
	return num_tokens
}
```


# available encodings
 | Encoding name           | OpenAI models                                        |
 | ----------------------- | ---------------------------------------------------- |
 | `cl100k_base`           | `gpt-4`, `gpt-3.5-turbo`, `text-embedding-ada-002`   |
 | `p50k_base`             | Codex models, `text-davinci-002`, `text-davinci-003` |
 | `r50k_base` (or `gpt2`) | GPT-3 models like `davinci`                          |



# available models
| Model name                   | OpenAI models |
| ---------------------------- | ------------- |
| gpt-4                        | cl100k_base   |
| gpt-4-*                      | cl100k_base   |
| gpt-3.5-turbo                | cl100k_base   |
| gpt-3.5-turbo-*              | cl100k_base   |
| text-davinci-003             | p50k_base     |
| text-davinci-002             | p50k_base     |
| text-davinci-001             | r50k_base     |
| text-curie-001               | r50k_base     |
| text-babbage-001             | r50k_base     |
| text-ada-001                 | r50k_base     |
| davinci                      | r50k_base     |
| curie                        | r50k_base     |
| babbage                      | r50k_base     |
| ada                          | r50k_base     |
| code-davinci-002             | p50k_base     |
| code-davinci-001             | p50k_base     |
| code-cushman-002             | p50k_base     |
| code-cushman-001             | p50k_base     |
| davinci-codex                | p50k_base     |
| cushman-codex                | p50k_base     |
| text-davinci-edit-001        | p50k_edit     |
| code-davinci-edit-001        | p50k_edit     |
| text-embedding-ada-002       | cl100k_base   |
| text-similarity-davinci-001  | r50k_base     |
| text-similarity-curie-001    | r50k_base     |
| text-similarity-babbage-001  | r50k_base     |
| text-similarity-ada-001      | r50k_base     |
| text-search-davinci-doc-001  | r50k_base     |
| text-search-curie-doc-001    | r50k_base     |
| text-search-babbage-doc-001  | r50k_base     |
| text-search-ada-doc-001      | r50k_base     |
| code-search-babbage-code-001 | r50k_base     |
| code-search-ada-code-001     | r50k_base     |
| gpt2                         | gpt2          |



# Test
> you can run test in [test](./test) folder

# compare with original [tiktoken](https://github.com/openai/tiktoken)

## get token by encoding
[result](./doc/test_result.md#encoding-test-result)

## get token by model  
[result](./doc/test_result.md#model-test-result)



# Benchmark
> you can run benchmark in [test](./test) folder

## Benchmark result
| name        | time/op | os         | cpu      | text                             | times  |
| ----------- | ------- | ---------- | -------- | -------------------------------- | ------ |
| tiktoken-go | 8795ns  | macOS 13.2 | Apple M1 | [UDHR](https://unicode.org/udhr) | 100000 |
| tiktoken    | 8838ns  | macOS 13.2 | Apple M1 | [UDHR](https://unicode.org/udhr) | 100000 |

It looks like the performance is almost the same.   

Maybe the difference is due to the difference in the performance of the machine.

Or maybe my benchmark method is not appropriate.  

If you have better benchmark method or if you want add your benchmark result, please feel free to submit a PR.

# License
[MIT](./LICENSE)
