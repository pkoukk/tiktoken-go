# tiktoken-go
Go 语言版本的 OpenAI 的 tiktoken。  
帮你把文本转换成 OpenAI 的模型可以识别的 token。  
tiktoken的原项目地址[tiktoken](https://github.com/openai/tiktoken).  

# 用法

## 安装


```bash
go get github.com/pkoukk/tiktoken-go
# 默认的tiktoken需要从openai下载token字典，如果想要离线使用，可以使用以下分支
go get github.com/pkoukk/tiktoken-go@embed
```
## 缓存
Tiktoken-go 和原始的 Tiktoken 库一样，具有相同的缓存机制。  
您可以使用环境变量 TIKTOKEN_CACHE_DIR 来设置缓存目录。  
一旦设置了该变量，tiktoken-go 将使用该目录来缓存令牌字典。  
如果您未设置此环境变量，则 tiktoken-go 将在每次首次初始化编码时下载字典。  

## 例子

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

### 计算chat API消息当中的token消耗
这段代码由@nasa1024根据[官方示例](https://github.com/openai/openai-cookbook/blob/main/examples/How_to_count_tokens_with_tiktoken.ipynb)编写

请注意，消息的token计算方式可能随时会发生改变，以下代码并不一定在将来适用，如果您需要精确的计算，请关注官方文档。

如果您发现这段代码不再适用，欢迎您提PR或Issue。

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

# 与官方 [tiktoken](https://github.com/openai/tiktoken) 的对比

## get token by encoding
[测试结果](./doc/test_result.md#encoding-test-result)

## get token by model  
[测试结果](./doc/test_result.md#model-test-result)

# Benchmark
> 你可以使用 [test](./test) 目录下的文件执行基准测试。 

## Benchmark result
| name        | time/op | os         | cpu      | text                             | times  |
| ----------- | ------- | ---------- | -------- | -------------------------------- | ------ |
| tiktoken-go | 8795ns  | macOS 13.2 | Apple M1 | [UDHR](https://unicode.org/udhr) | 100000 |
| tiktoken    | 8838ns  | macOS 13.2 | Apple M1 | [UDHR](https://unicode.org/udhr) | 100000 |

看上去tiktoken-go的性能基本与原tiktoken一致。  

也许在不同的机器上的测试结果会有所不同。也可能是我的测试方法并不恰当。

如果你有更好的测试方法，或者说你想添加在你机器上的测试结果，欢迎提PR。

# License
[MIT](./LICENSE)
