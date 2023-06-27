# Encoding Test Result
| python tiktoken                                          | golang tiktoken-go                                       |
| :------------------------------------------------------- | :------------------------------------------------------- |
| text: hallo world!, encoding: cl100k_base, token: 4      | text: hallo world!, encoding: cl100k_base, token: 4      |
| text: hallo world!, encoding: p50k_base, token: 4        | text: hallo world!, encoding: p50k_base, token: 4        |
| text: hallo world!, encoding: r50k_base, token: 4        | text: hallo world!, encoding: r50k_base, token: 4        |
| text: 你好世界！, encoding: cl100k_base, token: 6        | text: 你好世界！, encoding: cl100k_base, token: 6        |
| text: 你好世界！, encoding: p50k_base, token: 11         | text: 你好世界！, encoding: p50k_base, token: 11         |
| text: 你好世界！, encoding: r50k_base, token: 11         | text: 你好世界！, encoding: r50k_base, token: 11         |
| text: こんにちは世界！, encoding: cl100k_base, token: 5  | text: こんにちは世界！, encoding: cl100k_base, token: 5  |
| text: こんにちは世界！, encoding: p50k_base, token: 13   | text: こんにちは世界！, encoding: p50k_base, token: 13   |
| text: こんにちは世界！, encoding: r50k_base, token: 13   | text: こんにちは世界！, encoding: r50k_base, token: 13   |
| text: 안녕하세요 세계!, encoding: cl100k_base, token: 10 | text: 안녕하세요 세계!, encoding: cl100k_base, token: 10 |
| text: 안녕하세요 세계!, encoding: p50k_base, token: 21   | text: 안녕하세요 세계!, encoding: p50k_base, token: 21   |
| text: 안녕하세요 세계!, encoding: r50k_base, token: 21   | text: 안녕하세요 세계!, encoding: r50k_base, token: 21   |
| text: Привет мир!, encoding: cl100k_base, token: 6       | text: Привет мир!, encoding: cl100k_base, token: 6       |
| text: Привет мир!, encoding: p50k_base, token: 12        | text: Привет мир!, encoding: p50k_base, token: 12        |
| text: Привет мир!, encoding: r50k_base, token: 12        | text: Привет мир!, encoding: r50k_base, token: 12        |
| text: ¡Hola mundo!, encoding: cl100k_base, token: 4      | text: ¡Hola mundo!, encoding: cl100k_base, token: 4      |
| text: ¡Hola mundo!, encoding: p50k_base, token: 7        | text: ¡Hola mundo!, encoding: p50k_base, token: 7        |
| text: ¡Hola mundo!, encoding: r50k_base, token: 7        | text: ¡Hola mundo!, encoding: r50k_base, token: 7        |
| text: Hallo Welt!, encoding: cl100k_base, token: 3       | text: Hallo Welt!, encoding: cl100k_base, token: 3       |
| text: Hallo Welt!, encoding: p50k_base, token: 5         | text: Hallo Welt!, encoding: p50k_base, token: 5         |
| text: Hallo Welt!, encoding: r50k_base, token: 5         | text: Hallo Welt!, encoding: r50k_base, token: 5         |
| text: Bonjour le monde!, encoding: cl100k_base, token: 4 | text: Bonjour le monde!, encoding: cl100k_base, token: 4 |
| text: Bonjour le monde!, encoding: p50k_base, token: 7   | text: Bonjour le monde!, encoding: p50k_base, token: 7   |
| text: Bonjour le monde!, encoding: r50k_base, token: 7   | text: Bonjour le monde!, encoding: r50k_base, token: 7   |
| text: Ciao mondo!, encoding: cl100k_base, token: 4       | text: Ciao mondo!, encoding: cl100k_base, token: 4       |
| text: Ciao mondo!, encoding: p50k_base, token: 5         | text: Ciao mondo!, encoding: p50k_base, token: 5         |
| text: Ciao mondo!, encoding: r50k_base, token: 5         | text: Ciao mondo!, encoding: r50k_base, token: 5         |
| text: Hej världen!, encoding: cl100k_base, token: 7      | text: Hej världen!, encoding: cl100k_base, token: 7      |
| text: Hej världen!, encoding: p50k_base, token: 8        | text: Hej världen!, encoding: p50k_base, token: 8        |
| text: Hej världen!, encoding: r50k_base, token: 8        | text: Hej världen!, encoding: r50k_base, token: 8        |
| text: Hallo wereld!, encoding: cl100k_base, token: 3     | text: Hallo wereld!, encoding: cl100k_base, token: 3     |
| text: Hallo wereld!, encoding: p50k_base, token: 5       | text: Hallo wereld!, encoding: p50k_base, token: 5       |
| text: Hallo wereld!, encoding: r50k_base, token: 5       | text: Hallo wereld!, encoding: r50k_base, token: 5       |
| text: Hallo verden!, encoding: cl100k_base, token: 4     | text: Hallo verden!, encoding: cl100k_base, token: 4     |
| text: Hallo verden!, encoding: p50k_base, token: 5       | text: Hallo verden!, encoding: p50k_base, token: 5       |
| text: Hallo verden!, encoding: r50k_base, token: 5       | text: Hallo verden!, encoding: r50k_base, token: 5       |
| text: Hallo wereld!, encoding: cl100k_base, token: 3     | text: Hallo wereld!, encoding: cl100k_base, token: 3     |
| text: Hallo wereld!, encoding: p50k_base, token: 5       | text: Hallo wereld!, encoding: p50k_base, token: 5       |
| text: Hallo wereld!, encoding: r50k_base, token: 5       | text: Hallo wereld!, encoding: r50k_base, token: 5       |
| text: Hallo verden!, encoding: cl100k_base, token: 4     | text: Hallo verden!, encoding: cl100k_base, token: 4     |
| text: Hallo verden!, encoding: p50k_base, token: 5       | text: Hallo verden!, encoding: p50k_base, token: 5       |
| text: Hallo verden!, encoding: r50k_base, token: 5       | text: Hallo verden!, encoding: r50k_base, token: 5       |

# Model Test Result
| python tiktoken                                                       | golang tiktoken-go                                                    |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| text: hallo world!, model: gpt-4, token: 4                            | text: hallo world!, model: gpt-4, token: 4                            |
| text: hallo world!, model: gpt-3.5-turbo, token: 4                    | text: hallo world!, model: gpt-3.5-turbo, token: 4                    |
| text: hallo world!, model: text-davinci-003, token: 4                 | text: hallo world!, model: text-davinci-003, token: 4                 |
| text: hallo world!, model: text-davinci-002, token: 4                 | text: hallo world!, model: text-davinci-002, token: 4                 |
| text: hallo world!, model: text-davinci-001, token: 4                 | text: hallo world!, model: text-davinci-001, token: 4                 |
| text: hallo world!, model: text-curie-001, token: 4                   | text: hallo world!, model: text-curie-001, token: 4                   |
| text: hallo world!, model: text-babbage-001, token: 4                 | text: hallo world!, model: text-babbage-001, token: 4                 |
| text: hallo world!, model: text-ada-001, token: 4                     | text: hallo world!, model: text-ada-001, token: 4                     |
| text: hallo world!, model: davinci, token: 4                          | text: hallo world!, model: davinci, token: 4                          |
| text: hallo world!, model: curie, token: 4                            | text: hallo world!, model: curie, token: 4                            |
| text: hallo world!, model: babbage, token: 4                          | text: hallo world!, model: babbage, token: 4                          |
| text: hallo world!, model: ada, token: 4                              | text: hallo world!, model: ada, token: 4                              |
| text: hallo world!, model: code-davinci-002, token: 4                 | text: hallo world!, model: code-davinci-002, token: 4                 |
| text: hallo world!, model: code-davinci-001, token: 4                 | text: hallo world!, model: code-davinci-001, token: 4                 |
| text: hallo world!, model: code-cushman-002, token: 4                 | text: hallo world!, model: code-cushman-002, token: 4                 |
| text: hallo world!, model: code-cushman-001, token: 4                 | text: hallo world!, model: code-cushman-001, token: 4                 |
| text: hallo world!, model: davinci-codex, token: 4                    | text: hallo world!, model: davinci-codex, token: 4                    |
| text: hallo world!, model: cushman-codex, token: 4                    | text: hallo world!, model: cushman-codex, token: 4                    |
| text: hallo world!, model: text-davinci-edit-001, token: 4            | text: hallo world!, model: text-davinci-edit-001, token: 4            |
| text: hallo world!, model: code-davinci-edit-001, token: 4            | text: hallo world!, model: code-davinci-edit-001, token: 4            |
| text: hallo world!, model: text-embedding-ada-002, token: 4           | text: hallo world!, model: text-embedding-ada-002, token: 4           |
| text: hallo world!, model: text-similarity-davinci-001, token: 4      | text: hallo world!, model: text-similarity-davinci-001, token: 4      |
| text: 你好世界！, model: gpt-4, token: 6                              | text: 你好世界！, model: gpt-4, token: 6                              |
| text: 你好世界！, model: gpt-3.5-turbo, token: 6                      | text: 你好世界！, model: gpt-3.5-turbo, token: 6                      |
| text: 你好世界！, model: text-davinci-003, token: 11                  | text: 你好世界！, model: text-davinci-003, token: 11                  |
| text: 你好世界！, model: text-davinci-002, token: 11                  | text: 你好世界！, model: text-davinci-002, token: 11                  |
| text: 你好世界！, model: text-davinci-001, token: 11                  | text: 你好世界！, model: text-davinci-001, token: 11                  |
| text: 你好世界！, model: text-curie-001, token: 11                    | text: 你好世界！, model: text-curie-001, token: 11                    |
| text: 你好世界！, model: text-babbage-001, token: 11                  | text: 你好世界！, model: text-babbage-001, token: 11                  |
| text: 你好世界！, model: text-ada-001, token: 11                      | text: 你好世界！, model: text-ada-001, token: 11                      |
| text: 你好世界！, model: davinci, token: 11                           | text: 你好世界！, model: davinci, token: 11                           |
| text: 你好世界！, model: curie, token: 11                             | text: 你好世界！, model: curie, token: 11                             |
| text: 你好世界！, model: babbage, token: 11                           | text: 你好世界！, model: babbage, token: 11                           |
| text: 你好世界！, model: ada, token: 11                               | text: 你好世界！, model: ada, token: 11                               |
| text: 你好世界！, model: code-davinci-002, token: 11                  | text: 你好世界！, model: code-davinci-002, token: 11                  |
| text: 你好世界！, model: code-davinci-001, token: 11                  | text: 你好世界！, model: code-davinci-001, token: 11                  |
| text: 你好世界！, model: code-cushman-002, token: 11                  | text: 你好世界！, model: code-cushman-002, token: 11                  |
| text: 你好世界！, model: code-cushman-001, token: 11                  | text: 你好世界！, model: code-cushman-001, token: 11                  |
| text: 你好世界！, model: davinci-codex, token: 11                     | text: 你好世界！, model: davinci-codex, token: 11                     |
| text: 你好世界！, model: cushman-codex, token: 11                     | text: 你好世界！, model: cushman-codex, token: 11                     |
| text: 你好世界！, model: text-davinci-edit-001, token: 11             | text: 你好世界！, model: text-davinci-edit-001, token: 11             |
| text: 你好世界！, model: code-davinci-edit-001, token: 11             | text: 你好世界！, model: code-davinci-edit-001, token: 11             |
| text: 你好世界！, model: text-embedding-ada-002, token: 6             | text: 你好世界！, model: text-embedding-ada-002, token: 6             |
| text: 你好世界！, model: text-similarity-davinci-001, token: 11       | text: 你好世界！, model: text-similarity-davinci-001, token: 11       |
| text: こんにちは世界！, model: gpt-4, token: 5                        | text: こんにちは世界！, model: gpt-4, token: 5                        |
| text: こんにちは世界！, model: gpt-3.5-turbo, token: 5                | text: こんにちは世界！, model: gpt-3.5-turbo, token: 5                |
| text: こんにちは世界！, model: text-davinci-003, token: 13            | text: こんにちは世界！, model: text-davinci-003, token: 13            |
| text: こんにちは世界！, model: text-davinci-002, token: 13            | text: こんにちは世界！, model: text-davinci-002, token: 13            |
| text: こんにちは世界！, model: text-davinci-001, token: 13            | text: こんにちは世界！, model: text-davinci-001, token: 13            |
| text: こんにちは世界！, model: text-curie-001, token: 13              | text: こんにちは世界！, model: text-curie-001, token: 13              |
| text: こんにちは世界！, model: text-babbage-001, token: 13            | text: こんにちは世界！, model: text-babbage-001, token: 13            |
| text: こんにちは世界！, model: text-ada-001, token: 13                | text: こんにちは世界！, model: text-ada-001, token: 13                |
| text: こんにちは世界！, model: davinci, token: 13                     | text: こんにちは世界！, model: davinci, token: 13                     |
| text: こんにちは世界！, model: curie, token: 13                       | text: こんにちは世界！, model: curie, token: 13                       |
| text: こんにちは世界！, model: babbage, token: 13                     | text: こんにちは世界！, model: babbage, token: 13                     |
| text: こんにちは世界！, model: ada, token: 13                         | text: こんにちは世界！, model: ada, token: 13                         |
| text: こんにちは世界！, model: code-davinci-002, token: 13            | text: こんにちは世界！, model: code-davinci-002, token: 13            |
| text: こんにちは世界！, model: code-davinci-001, token: 13            | text: こんにちは世界！, model: code-davinci-001, token: 13            |
| text: こんにちは世界！, model: code-cushman-002, token: 13            | text: こんにちは世界！, model: code-cushman-002, token: 13            |
| text: こんにちは世界！, model: code-cushman-001, token: 13            | text: こんにちは世界！, model: code-cushman-001, token: 13            |
| text: こんにちは世界！, model: davinci-codex, token: 13               | text: こんにちは世界！, model: davinci-codex, token: 13               |
| text: こんにちは世界！, model: cushman-codex, token: 13               | text: こんにちは世界！, model: cushman-codex, token: 13               |
| text: こんにちは世界！, model: text-davinci-edit-001, token: 13       | text: こんにちは世界！, model: text-davinci-edit-001, token: 13       |
| text: こんにちは世界！, model: code-davinci-edit-001, token: 13       | text: こんにちは世界！, model: code-davinci-edit-001, token: 13       |
| text: こんにちは世界！, model: text-embedding-ada-002, token: 5       | text: こんにちは世界！, model: text-embedding-ada-002, token: 5       |
| text: こんにちは世界！, model: text-similarity-davinci-001, token: 13 | text: こんにちは世界！, model: text-similarity-davinci-001, token: 13 |
| text: 안녕하세요 세계!, model: gpt-4, token: 10                       | text: 안녕하세요 세계!, model: gpt-4, token: 10                       |
| text: 안녕하세요 세계!, model: gpt-3.5-turbo, token: 10               | text: 안녕하세요 세계!, model: gpt-3.5-turbo, token: 10               |
| text: 안녕하세요 세계!, model: text-davinci-003, token: 21            | text: 안녕하세요 세계!, model: text-davinci-003, token: 21            |
| text: 안녕하세요 세계!, model: text-davinci-002, token: 21            | text: 안녕하세요 세계!, model: text-davinci-002, token: 21            |
| text: 안녕하세요 세계!, model: text-davinci-001, token: 21            | text: 안녕하세요 세계!, model: text-davinci-001, token: 21            |
| text: 안녕하세요 세계!, model: text-curie-001, token: 21              | text: 안녕하세요 세계!, model: text-curie-001, token: 21              |
| text: 안녕하세요 세계!, model: text-babbage-001, token: 21            | text: 안녕하세요 세계!, model: text-babbage-001, token: 21            |
| text: 안녕하세요 세계!, model: text-ada-001, token: 21                | text: 안녕하세요 세계!, model: text-ada-001, token: 21                |
| text: 안녕하세요 세계!, model: davinci, token: 21                     | text: 안녕하세요 세계!, model: davinci, token: 21                     |
| text: 안녕하세요 세계!, model: curie, token: 21                       | text: 안녕하세요 세계!, model: curie, token: 21                       |
| text: 안녕하세요 세계!, model: babbage, token: 21                     | text: 안녕하세요 세계!, model: babbage, token: 21                     |
| text: 안녕하세요 세계!, model: ada, token: 21                         | text: 안녕하세요 세계!, model: ada, token: 21                         |
| text: 안녕하세요 세계!, model: code-davinci-002, token: 21            | text: 안녕하세요 세계!, model: code-davinci-002, token: 21            |
| text: 안녕하세요 세계!, model: code-davinci-001, token: 21            | text: 안녕하세요 세계!, model: code-davinci-001, token: 21            |
| text: 안녕하세요 세계!, model: code-cushman-002, token: 21            | text: 안녕하세요 세계!, model: code-cushman-002, token: 21            |
| text: 안녕하세요 세계!, model: code-cushman-001, token: 21            | text: 안녕하세요 세계!, model: code-cushman-001, token: 21            |
| text: 안녕하세요 세계!, model: davinci-codex, token: 21               | text: 안녕하세요 세계!, model: davinci-codex, token: 21               |
| text: 안녕하세요 세계!, model: cushman-codex, token: 21               | text: 안녕하세요 세계!, model: cushman-codex, token: 21               |
| text: 안녕하세요 세계!, model: text-davinci-edit-001, token: 21       | text: 안녕하세요 세계!, model: text-davinci-edit-001, token: 21       |
| text: 안녕하세요 세계!, model: code-davinci-edit-001, token: 21       | text: 안녕하세요 세계!, model: code-davinci-edit-001, token: 21       |
| text: 안녕하세요 세계!, model: text-embedding-ada-002, token: 10      | text: 안녕하세요 세계!, model: text-embedding-ada-002, token: 10      |
| text: 안녕하세요 세계!, model: text-similarity-davinci-001, token: 21 | text: 안녕하세요 세계!, model: text-similarity-davinci-001, token: 21 |
| text: Привет мир!, model: gpt-4, token: 6                             | text: Привет мир!, model: gpt-4, token: 6                             |
| text: Привет мир!, model: gpt-3.5-turbo, token: 6                     | text: Привет мир!, model: gpt-3.5-turbo, token: 6                     |
| text: Привет мир!, model: text-davinci-003, token: 12                 | text: Привет мир!, model: text-davinci-003, token: 12                 |
| text: Привет мир!, model: text-davinci-002, token: 12                 | text: Привет мир!, model: text-davinci-002, token: 12                 |
| text: Привет мир!, model: text-davinci-001, token: 12                 | text: Привет мир!, model: text-davinci-001, token: 12                 |
| text: Привет мир!, model: text-curie-001, token: 12                   | text: Привет мир!, model: text-curie-001, token: 12                   |
| text: Привет мир!, model: text-babbage-001, token: 12                 | text: Привет мир!, model: text-babbage-001, token: 12                 |
| text: Привет мир!, model: text-ada-001, token: 12                     | text: Привет мир!, model: text-ada-001, token: 12                     |
| text: Привет мир!, model: davinci, token: 12                          | text: Привет мир!, model: davinci, token: 12                          |
| text: Привет мир!, model: curie, token: 12                            | text: Привет мир!, model: curie, token: 12                            |
| text: Привет мир!, model: babbage, token: 12                          | text: Привет мир!, model: babbage, token: 12                          |
| text: Привет мир!, model: ada, token: 12                              | text: Привет мир!, model: ada, token: 12                              |
| text: Привет мир!, model: code-davinci-002, token: 12                 | text: Привет мир!, model: code-davinci-002, token: 12                 |
| text: Привет мир!, model: code-davinci-001, token: 12                 | text: Привет мир!, model: code-davinci-001, token: 12                 |
| text: Привет мир!, model: code-cushman-002, token: 12                 | text: Привет мир!, model: code-cushman-002, token: 12                 |
| text: Привет мир!, model: code-cushman-001, token: 12                 | text: Привет мир!, model: code-cushman-001, token: 12                 |
| text: Привет мир!, model: davinci-codex, token: 12                    | text: Привет мир!, model: davinci-codex, token: 12                    |
| text: Привет мир!, model: cushman-codex, token: 12                    | text: Привет мир!, model: cushman-codex, token: 12                    |
| text: Привет мир!, model: text-davinci-edit-001, token: 12            | text: Привет мир!, model: text-davinci-edit-001, token: 12            |
| text: Привет мир!, model: code-davinci-edit-001, token: 12            | text: Привет мир!, model: code-davinci-edit-001, token: 12            |
| text: Привет мир!, model: text-embedding-ada-002, token: 6            | text: Привет мир!, model: text-embedding-ada-002, token: 6            |
| text: Привет мир!, model: text-similarity-davinci-001, token: 12      | text: Привет мир!, model: text-similarity-davinci-001, token: 12      |
| text: ¡Hola mundo!, model: gpt-4, token: 4                            | text: ¡Hola mundo!, model: gpt-4, token: 4                            |
| text: ¡Hola mundo!, model: gpt-3.5-turbo, token: 4                    | text: ¡Hola mundo!, model: gpt-3.5-turbo, token: 4                    |
| text: ¡Hola mundo!, model: text-davinci-003, token: 7                 | text: ¡Hola mundo!, model: text-davinci-003, token: 7                 |
| text: ¡Hola mundo!, model: text-davinci-002, token: 7                 | text: ¡Hola mundo!, model: text-davinci-002, token: 7                 |
| text: ¡Hola mundo!, model: text-davinci-001, token: 7                 | text: ¡Hola mundo!, model: text-davinci-001, token: 7                 |
| text: ¡Hola mundo!, model: text-curie-001, token: 7                   | text: ¡Hola mundo!, model: text-curie-001, token: 7                   |
| text: ¡Hola mundo!, model: text-babbage-001, token: 7                 | text: ¡Hola mundo!, model: text-babbage-001, token: 7                 |
| text: ¡Hola mundo!, model: text-ada-001, token: 7                     | text: ¡Hola mundo!, model: text-ada-001, token: 7                     |
| text: ¡Hola mundo!, model: davinci, token: 7                          | text: ¡Hola mundo!, model: davinci, token: 7                          |
| text: ¡Hola mundo!, model: curie, token: 7                            | text: ¡Hola mundo!, model: curie, token: 7                            |
| text: ¡Hola mundo!, model: babbage, token: 7                          | text: ¡Hola mundo!, model: babbage, token: 7                          |
| text: ¡Hola mundo!, model: ada, token: 7                              | text: ¡Hola mundo!, model: ada, token: 7                              |
| text: ¡Hola mundo!, model: code-davinci-002, token: 7                 | text: ¡Hola mundo!, model: code-davinci-002, token: 7                 |
| text: ¡Hola mundo!, model: code-davinci-001, token: 7                 | text: ¡Hola mundo!, model: code-davinci-001, token: 7                 |
| text: ¡Hola mundo!, model: code-cushman-002, token: 7                 | text: ¡Hola mundo!, model: code-cushman-002, token: 7                 |
| text: ¡Hola mundo!, model: code-cushman-001, token: 7                 | text: ¡Hola mundo!, model: code-cushman-001, token: 7                 |
| text: ¡Hola mundo!, model: davinci-codex, token: 7                    | text: ¡Hola mundo!, model: davinci-codex, token: 7                    |
| text: ¡Hola mundo!, model: cushman-codex, token: 7                    | text: ¡Hola mundo!, model: cushman-codex, token: 7                    |
| text: ¡Hola mundo!, model: text-davinci-edit-001, token: 7            | text: ¡Hola mundo!, model: text-davinci-edit-001, token: 7            |
| text: ¡Hola mundo!, model: code-davinci-edit-001, token: 7            | text: ¡Hola mundo!, model: code-davinci-edit-001, token: 7            |
| text: ¡Hola mundo!, model: text-embedding-ada-002, token: 4           | text: ¡Hola mundo!, model: text-embedding-ada-002, token: 4           |
| text: ¡Hola mundo!, model: text-similarity-davinci-001, token: 7      | text: ¡Hola mundo!, model: text-similarity-davinci-001, token: 7      |
| text: Hallo Welt!, model: gpt-4, token: 3                             | text: Hallo Welt!, model: gpt-4, token: 3                             |
| text: Hallo Welt!, model: gpt-3.5-turbo, token: 3                     | text: Hallo Welt!, model: gpt-3.5-turbo, token: 3                     |
| text: Hallo Welt!, model: text-davinci-003, token: 5                  | text: Hallo Welt!, model: text-davinci-003, token: 5                  |
| text: Hallo Welt!, model: text-davinci-002, token: 5                  | text: Hallo Welt!, model: text-davinci-002, token: 5                  |
| text: Hallo Welt!, model: text-davinci-001, token: 5                  | text: Hallo Welt!, model: text-davinci-001, token: 5                  |
| text: Hallo Welt!, model: text-curie-001, token: 5                    | text: Hallo Welt!, model: text-curie-001, token: 5                    |
| text: Hallo Welt!, model: text-babbage-001, token: 5                  | text: Hallo Welt!, model: text-babbage-001, token: 5                  |
| text: Hallo Welt!, model: text-ada-001, token: 5                      | text: Hallo Welt!, model: text-ada-001, token: 5                      |
| text: Hallo Welt!, model: davinci, token: 5                           | text: Hallo Welt!, model: davinci, token: 5                           |
| text: Hallo Welt!, model: curie, token: 5                             | text: Hallo Welt!, model: curie, token: 5                             |
| text: Hallo Welt!, model: babbage, token: 5                           | text: Hallo Welt!, model: babbage, token: 5                           |
| text: Hallo Welt!, model: ada, token: 5                               | text: Hallo Welt!, model: ada, token: 5                               |
| text: Hallo Welt!, model: code-davinci-002, token: 5                  | text: Hallo Welt!, model: code-davinci-002, token: 5                  |
| text: Hallo Welt!, model: code-davinci-001, token: 5                  | text: Hallo Welt!, model: code-davinci-001, token: 5                  |
| text: Hallo Welt!, model: code-cushman-002, token: 5                  | text: Hallo Welt!, model: code-cushman-002, token: 5                  |
| text: Hallo Welt!, model: code-cushman-001, token: 5                  | text: Hallo Welt!, model: code-cushman-001, token: 5                  |
| text: Hallo Welt!, model: davinci-codex, token: 5                     | text: Hallo Welt!, model: davinci-codex, token: 5                     |
| text: Hallo Welt!, model: cushman-codex, token: 5                     | text: Hallo Welt!, model: cushman-codex, token: 5                     |
| text: Hallo Welt!, model: text-davinci-edit-001, token: 5             | text: Hallo Welt!, model: text-davinci-edit-001, token: 5             |
| text: Hallo Welt!, model: code-davinci-edit-001, token: 5             | text: Hallo Welt!, model: code-davinci-edit-001, token: 5             |
| text: Hallo Welt!, model: text-embedding-ada-002, token: 3            | text: Hallo Welt!, model: text-embedding-ada-002, token: 3            |
| text: Hallo Welt!, model: text-similarity-davinci-001, token: 5       | text: Hallo Welt!, model: text-similarity-davinci-001, token: 5       |
| text: Bonjour le monde!, model: gpt-4, token: 4                       | text: Bonjour le monde!, model: gpt-4, token: 4                       |
| text: Bonjour le monde!, model: gpt-3.5-turbo, token: 4               | text: Bonjour le monde!, model: gpt-3.5-turbo, token: 4               |
| text: Bonjour le monde!, model: text-davinci-003, token: 7            | text: Bonjour le monde!, model: text-davinci-003, token: 7            |
| text: Bonjour le monde!, model: text-davinci-002, token: 7            | text: Bonjour le monde!, model: text-davinci-002, token: 7            |
| text: Bonjour le monde!, model: text-davinci-001, token: 7            | text: Bonjour le monde!, model: text-davinci-001, token: 7            |
| text: Bonjour le monde!, model: text-curie-001, token: 7              | text: Bonjour le monde!, model: text-curie-001, token: 7              |
| text: Bonjour le monde!, model: text-babbage-001, token: 7            | text: Bonjour le monde!, model: text-babbage-001, token: 7            |
| text: Bonjour le monde!, model: text-ada-001, token: 7                | text: Bonjour le monde!, model: text-ada-001, token: 7                |
| text: Bonjour le monde!, model: davinci, token: 7                     | text: Bonjour le monde!, model: davinci, token: 7                     |
| text: Bonjour le monde!, model: curie, token: 7                       | text: Bonjour le monde!, model: curie, token: 7                       |
| text: Bonjour le monde!, model: babbage, token: 7                     | text: Bonjour le monde!, model: babbage, token: 7                     |
| text: Bonjour le monde!, model: ada, token: 7                         | text: Bonjour le monde!, model: ada, token: 7                         |
| text: Bonjour le monde!, model: code-davinci-002, token: 7            | text: Bonjour le monde!, model: code-davinci-002, token: 7            |
| text: Bonjour le monde!, model: code-davinci-001, token: 7            | text: Bonjour le monde!, model: code-davinci-001, token: 7            |
| text: Bonjour le monde!, model: code-cushman-002, token: 7            | text: Bonjour le monde!, model: code-cushman-002, token: 7            |
| text: Bonjour le monde!, model: code-cushman-001, token: 7            | text: Bonjour le monde!, model: code-cushman-001, token: 7            |
| text: Bonjour le monde!, model: davinci-codex, token: 7               | text: Bonjour le monde!, model: davinci-codex, token: 7               |
| text: Bonjour le monde!, model: cushman-codex, token: 7               | text: Bonjour le monde!, model: cushman-codex, token: 7               |
| text: Bonjour le monde!, model: text-davinci-edit-001, token: 7       | text: Bonjour le monde!, model: text-davinci-edit-001, token: 7       |
| text: Bonjour le monde!, model: code-davinci-edit-001, token: 7       | text: Bonjour le monde!, model: code-davinci-edit-001, token: 7       |
| text: Bonjour le monde!, model: text-embedding-ada-002, token: 4      | text: Bonjour le monde!, model: text-embedding-ada-002, token: 4      |
| text: Bonjour le monde!, model: text-similarity-davinci-001, token: 7 | text: Bonjour le monde!, model: text-similarity-davinci-001, token: 7 |
| text: Ciao mondo!, model: gpt-4, token: 4                             | text: Ciao mondo!, model: gpt-4, token: 4                             |
| text: Ciao mondo!, model: gpt-3.5-turbo, token: 4                     | text: Ciao mondo!, model: gpt-3.5-turbo, token: 4                     |
| text: Ciao mondo!, model: text-davinci-003, token: 5                  | text: Ciao mondo!, model: text-davinci-003, token: 5                  |
| text: Ciao mondo!, model: text-davinci-002, token: 5                  | text: Ciao mondo!, model: text-davinci-002, token: 5                  |
| text: Ciao mondo!, model: text-davinci-001, token: 5                  | text: Ciao mondo!, model: text-davinci-001, token: 5                  |
| text: Ciao mondo!, model: text-curie-001, token: 5                    | text: Ciao mondo!, model: text-curie-001, token: 5                    |
| text: Ciao mondo!, model: text-babbage-001, token: 5                  | text: Ciao mondo!, model: text-babbage-001, token: 5                  |
| text: Ciao mondo!, model: text-ada-001, token: 5                      | text: Ciao mondo!, model: text-ada-001, token: 5                      |
| text: Ciao mondo!, model: davinci, token: 5                           | text: Ciao mondo!, model: davinci, token: 5                           |
| text: Ciao mondo!, model: curie, token: 5                             | text: Ciao mondo!, model: curie, token: 5                             |
| text: Ciao mondo!, model: babbage, token: 5                           | text: Ciao mondo!, model: babbage, token: 5                           |
| text: Ciao mondo!, model: ada, token: 5                               | text: Ciao mondo!, model: ada, token: 5                               |
| text: Ciao mondo!, model: code-davinci-002, token: 5                  | text: Ciao mondo!, model: code-davinci-002, token: 5                  |
| text: Ciao mondo!, model: code-davinci-001, token: 5                  | text: Ciao mondo!, model: code-davinci-001, token: 5                  |
| text: Ciao mondo!, model: code-cushman-002, token: 5                  | text: Ciao mondo!, model: code-cushman-002, token: 5                  |
| text: Ciao mondo!, model: code-cushman-001, token: 5                  | text: Ciao mondo!, model: code-cushman-001, token: 5                  |
| text: Ciao mondo!, model: davinci-codex, token: 5                     | text: Ciao mondo!, model: davinci-codex, token: 5                     |
| text: Ciao mondo!, model: cushman-codex, token: 5                     | text: Ciao mondo!, model: cushman-codex, token: 5                     |
| text: Ciao mondo!, model: text-davinci-edit-001, token: 5             | text: Ciao mondo!, model: text-davinci-edit-001, token: 5             |
| text: Ciao mondo!, model: code-davinci-edit-001, token: 5             | text: Ciao mondo!, model: code-davinci-edit-001, token: 5             |
| text: Ciao mondo!, model: text-embedding-ada-002, token: 4            | text: Ciao mondo!, model: text-embedding-ada-002, token: 4            |
| text: Ciao mondo!, model: text-similarity-davinci-001, token: 5       | text: Ciao mondo!, model: text-similarity-davinci-001, token: 5       |
| text: Hej världen!, model: gpt-4, token: 7                            | text: Hej världen!, model: gpt-4, token: 7                            |
| text: Hej världen!, model: gpt-3.5-turbo, token: 7                    | text: Hej världen!, model: gpt-3.5-turbo, token: 7                    |
| text: Hej världen!, model: text-davinci-003, token: 8                 | text: Hej världen!, model: text-davinci-003, token: 8                 |
| text: Hej världen!, model: text-davinci-002, token: 8                 | text: Hej världen!, model: text-davinci-002, token: 8                 |
| text: Hej världen!, model: text-davinci-001, token: 8                 | text: Hej världen!, model: text-davinci-001, token: 8                 |
| text: Hej världen!, model: text-curie-001, token: 8                   | text: Hej världen!, model: text-curie-001, token: 8                   |
| text: Hej världen!, model: text-babbage-001, token: 8                 | text: Hej världen!, model: text-babbage-001, token: 8                 |
| text: Hej världen!, model: text-ada-001, token: 8                     | text: Hej världen!, model: text-ada-001, token: 8                     |
| text: Hej världen!, model: davinci, token: 8                          | text: Hej världen!, model: davinci, token: 8                          |
| text: Hej världen!, model: curie, token: 8                            | text: Hej världen!, model: curie, token: 8                            |
| text: Hej världen!, model: babbage, token: 8                          | text: Hej världen!, model: babbage, token: 8                          |
| text: Hej världen!, model: ada, token: 8                              | text: Hej världen!, model: ada, token: 8                              |
| text: Hej världen!, model: code-davinci-002, token: 8                 | text: Hej världen!, model: code-davinci-002, token: 8                 |
| text: Hej världen!, model: code-davinci-001, token: 8                 | text: Hej världen!, model: code-davinci-001, token: 8                 |
| text: Hej världen!, model: code-cushman-002, token: 8                 | text: Hej världen!, model: code-cushman-002, token: 8                 |
| text: Hej världen!, model: code-cushman-001, token: 8                 | text: Hej världen!, model: code-cushman-001, token: 8                 |
| text: Hej världen!, model: davinci-codex, token: 8                    | text: Hej världen!, model: davinci-codex, token: 8                    |
| text: Hej världen!, model: cushman-codex, token: 8                    | text: Hej världen!, model: cushman-codex, token: 8                    |
| text: Hej världen!, model: text-davinci-edit-001, token: 8            | text: Hej världen!, model: text-davinci-edit-001, token: 8            |
| text: Hej världen!, model: code-davinci-edit-001, token: 8            | text: Hej världen!, model: code-davinci-edit-001, token: 8            |
| text: Hej världen!, model: text-embedding-ada-002, token: 7           | text: Hej världen!, model: text-embedding-ada-002, token: 7           |
| text: Hej världen!, model: text-similarity-davinci-001, token: 8      | text: Hej världen!, model: text-similarity-davinci-001, token: 8      |
| text: Hallo wereld!, model: gpt-4, token: 3                           | text: Hallo wereld!, model: gpt-4, token: 3                           |
| text: Hallo wereld!, model: gpt-3.5-turbo, token: 3                   | text: Hallo wereld!, model: gpt-3.5-turbo, token: 3                   |
| text: Hallo wereld!, model: text-davinci-003, token: 5                | text: Hallo wereld!, model: text-davinci-003, token: 5                |
| text: Hallo wereld!, model: text-davinci-002, token: 5                | text: Hallo wereld!, model: text-davinci-002, token: 5                |
| text: Hallo wereld!, model: text-davinci-001, token: 5                | text: Hallo wereld!, model: text-davinci-001, token: 5                |
| text: Hallo wereld!, model: text-curie-001, token: 5                  | text: Hallo wereld!, model: text-curie-001, token: 5                  |
| text: Hallo wereld!, model: text-babbage-001, token: 5                | text: Hallo wereld!, model: text-babbage-001, token: 5                |
| text: Hallo wereld!, model: text-ada-001, token: 5                    | text: Hallo wereld!, model: text-ada-001, token: 5                    |
| text: Hallo wereld!, model: davinci, token: 5                         | text: Hallo wereld!, model: davinci, token: 5                         |
| text: Hallo wereld!, model: curie, token: 5                           | text: Hallo wereld!, model: curie, token: 5                           |
| text: Hallo wereld!, model: babbage, token: 5                         | text: Hallo wereld!, model: babbage, token: 5                         |
| text: Hallo wereld!, model: ada, token: 5                             | text: Hallo wereld!, model: ada, token: 5                             |
| text: Hallo wereld!, model: code-davinci-002, token: 5                | text: Hallo wereld!, model: code-davinci-002, token: 5                |
| text: Hallo wereld!, model: code-davinci-001, token: 5                | text: Hallo wereld!, model: code-davinci-001, token: 5                |
| text: Hallo wereld!, model: code-cushman-002, token: 5                | text: Hallo wereld!, model: code-cushman-002, token: 5                |
| text: Hallo wereld!, model: code-cushman-001, token: 5                | text: Hallo wereld!, model: code-cushman-001, token: 5                |
| text: Hallo wereld!, model: davinci-codex, token: 5                   | text: Hallo wereld!, model: davinci-codex, token: 5                   |
| text: Hallo wereld!, model: cushman-codex, token: 5                   | text: Hallo wereld!, model: cushman-codex, token: 5                   |
| text: Hallo wereld!, model: text-davinci-edit-001, token: 5           | text: Hallo wereld!, model: text-davinci-edit-001, token: 5           |
| text: Hallo wereld!, model: code-davinci-edit-001, token: 5           | text: Hallo wereld!, model: code-davinci-edit-001, token: 5           |
| text: Hallo wereld!, model: text-embedding-ada-002, token: 3          | text: Hallo wereld!, model: text-embedding-ada-002, token: 3          |
| text: Hallo wereld!, model: text-similarity-davinci-001, token: 5     | text: Hallo wereld!, model: text-similarity-davinci-001, token: 5     |
| text: Hallo verden!, model: gpt-4, token: 4                           | text: Hallo verden!, model: gpt-4, token: 4                           |
| text: Hallo verden!, model: gpt-3.5-turbo, token: 4                   | text: Hallo verden!, model: gpt-3.5-turbo, token: 4                   |
| text: Hallo verden!, model: text-davinci-003, token: 5                | text: Hallo verden!, model: text-davinci-003, token: 5                |
| text: Hallo verden!, model: text-davinci-002, token: 5                | text: Hallo verden!, model: text-davinci-002, token: 5                |
| text: Hallo verden!, model: text-davinci-001, token: 5                | text: Hallo verden!, model: text-davinci-001, token: 5                |
| text: Hallo verden!, model: text-curie-001, token: 5                  | text: Hallo verden!, model: text-curie-001, token: 5                  |
| text: Hallo verden!, model: text-babbage-001, token: 5                | text: Hallo verden!, model: text-babbage-001, token: 5                |
| text: Hallo verden!, model: text-ada-001, token: 5                    | text: Hallo verden!, model: text-ada-001, token: 5                    |
| text: Hallo verden!, model: davinci, token: 5                         | text: Hallo verden!, model: davinci, token: 5                         |
| text: Hallo verden!, model: curie, token: 5                           | text: Hallo verden!, model: curie, token: 5                           |
| text: Hallo verden!, model: babbage, token: 5                         | text: Hallo verden!, model: babbage, token: 5                         |
| text: Hallo verden!, model: ada, token: 5                             | text: Hallo verden!, model: ada, token: 5                             |
| text: Hallo verden!, model: code-davinci-002, token: 5                | text: Hallo verden!, model: code-davinci-002, token: 5                |
| text: Hallo verden!, model: code-davinci-001, token: 5                | text: Hallo verden!, model: code-davinci-001, token: 5                |
| text: Hallo verden!, model: code-cushman-002, token: 5                | text: Hallo verden!, model: code-cushman-002, token: 5                |
| text: Hallo verden!, model: code-cushman-001, token: 5                | text: Hallo verden!, model: code-cushman-001, token: 5                |
| text: Hallo verden!, model: davinci-codex, token: 5                   | text: Hallo verden!, model: davinci-codex, token: 5                   |
| text: Hallo verden!, model: cushman-codex, token: 5                   | text: Hallo verden!, model: cushman-codex, token: 5                   |
| text: Hallo verden!, model: text-davinci-edit-001, token: 5           | text: Hallo verden!, model: text-davinci-edit-001, token: 5           |
| text: Hallo verden!, model: code-davinci-edit-001, token: 5           | text: Hallo verden!, model: code-davinci-edit-001, token: 5           |
| text: Hallo verden!, model: text-embedding-ada-002, token: 4          | text: Hallo verden!, model: text-embedding-ada-002, token: 4          |
| text: Hallo verden!, model: text-similarity-davinci-001, token: 5     | text: Hallo verden!, model: text-similarity-davinci-001, token: 5     |
| text: Hallo wereld!, model: gpt-4, token: 3                           | text: Hallo wereld!, model: gpt-4, token: 3                           |
| text: Hallo wereld!, model: gpt-3.5-turbo, token: 3                   | text: Hallo wereld!, model: gpt-3.5-turbo, token: 3                   |
| text: Hallo wereld!, model: text-davinci-003, token: 5                | text: Hallo wereld!, model: text-davinci-003, token: 5                |
| text: Hallo wereld!, model: text-davinci-002, token: 5                | text: Hallo wereld!, model: text-davinci-002, token: 5                |
| text: Hallo wereld!, model: text-davinci-001, token: 5                | text: Hallo wereld!, model: text-davinci-001, token: 5                |
| text: Hallo wereld!, model: text-curie-001, token: 5                  | text: Hallo wereld!, model: text-curie-001, token: 5                  |
| text: Hallo wereld!, model: text-babbage-001, token: 5                | text: Hallo wereld!, model: text-babbage-001, token: 5                |
| text: Hallo wereld!, model: text-ada-001, token: 5                    | text: Hallo wereld!, model: text-ada-001, token: 5                    |
| text: Hallo wereld!, model: davinci, token: 5                         | text: Hallo wereld!, model: davinci, token: 5                         |
| text: Hallo wereld!, model: curie, token: 5                           | text: Hallo wereld!, model: curie, token: 5                           |
| text: Hallo wereld!, model: babbage, token: 5                         | text: Hallo wereld!, model: babbage, token: 5                         |
| text: Hallo wereld!, model: ada, token: 5                             | text: Hallo wereld!, model: ada, token: 5                             |
| text: Hallo wereld!, model: code-davinci-002, token: 5                | text: Hallo wereld!, model: code-davinci-002, token: 5                |
| text: Hallo wereld!, model: code-davinci-001, token: 5                | text: Hallo wereld!, model: code-davinci-001, token: 5                |
| text: Hallo wereld!, model: code-cushman-002, token: 5                | text: Hallo wereld!, model: code-cushman-002, token: 5                |
| text: Hallo wereld!, model: code-cushman-001, token: 5                | text: Hallo wereld!, model: code-cushman-001, token: 5                |
| text: Hallo wereld!, model: davinci-codex, token: 5                   | text: Hallo wereld!, model: davinci-codex, token: 5                   |
| text: Hallo wereld!, model: cushman-codex, token: 5                   | text: Hallo wereld!, model: cushman-codex, token: 5                   |
| text: Hallo wereld!, model: text-davinci-edit-001, token: 5           | text: Hallo wereld!, model: text-davinci-edit-001, token: 5           |
| text: Hallo wereld!, model: code-davinci-edit-001, token: 5           | text: Hallo wereld!, model: code-davinci-edit-001, token: 5           |
| text: Hallo wereld!, model: text-embedding-ada-002, token: 3          | text: Hallo wereld!, model: text-embedding-ada-002, token: 3          |
| text: Hallo wereld!, model: text-similarity-davinci-001, token: 5     | text: Hallo wereld!, model: text-similarity-davinci-001, token: 5     |
| text: Hallo verden!, model: gpt-4, token: 4                           | text: Hallo verden!, model: gpt-4, token: 4                           |
| text: Hallo verden!, model: gpt-3.5-turbo, token: 4                   | text: Hallo verden!, model: gpt-3.5-turbo, token: 4                   |
| text: Hallo verden!, model: text-davinci-003, token: 5                | text: Hallo verden!, model: text-davinci-003, token: 5                |
| text: Hallo verden!, model: text-davinci-002, token: 5                | text: Hallo verden!, model: text-davinci-002, token: 5                |
| text: Hallo verden!, model: text-davinci-001, token: 5                | text: Hallo verden!, model: text-davinci-001, token: 5                |
| text: Hallo verden!, model: text-curie-001, token: 5                  | text: Hallo verden!, model: text-curie-001, token: 5                  |
| text: Hallo verden!, model: text-babbage-001, token: 5                | text: Hallo verden!, model: text-babbage-001, token: 5                |
| text: Hallo verden!, model: text-ada-001, token: 5                    | text: Hallo verden!, model: text-ada-001, token: 5                    |
| text: Hallo verden!, model: davinci, token: 5                         | text: Hallo verden!, model: davinci, token: 5                         |
| text: Hallo verden!, model: curie, token: 5                           | text: Hallo verden!, model: curie, token: 5                           |
| text: Hallo verden!, model: babbage, token: 5                         | text: Hallo verden!, model: babbage, token: 5                         |
| text: Hallo verden!, model: ada, token: 5                             | text: Hallo verden!, model: ada, token: 5                             |
| text: Hallo verden!, model: code-davinci-002, token: 5                | text: Hallo verden!, model: code-davinci-002, token: 5                |
| text: Hallo verden!, model: code-davinci-001, token: 5                | text: Hallo verden!, model: code-davinci-001, token: 5                |
| text: Hallo verden!, model: code-cushman-002, token: 5                | text: Hallo verden!, model: code-cushman-002, token: 5                |
| text: Hallo verden!, model: code-cushman-001, token: 5                | text: Hallo verden!, model: code-cushman-001, token: 5                |
| text: Hallo verden!, model: davinci-codex, token: 5                   | text: Hallo verden!, model: davinci-codex, token: 5                   |
| text: Hallo verden!, model: cushman-codex, token: 5                   | text: Hallo verden!, model: cushman-codex, token: 5                   |
| text: Hallo verden!, model: text-davinci-edit-001, token: 5           | text: Hallo verden!, model: text-davinci-edit-001, token: 5           |
| text: Hallo verden!, model: code-davinci-edit-001, token: 5           | text: Hallo verden!, model: code-davinci-edit-001, token: 5           |
| text: Hallo verden!, model: text-embedding-ada-002, token: 4          | text: Hallo verden!, model: text-embedding-ada-002, token: 4          |
| text: Hallo verden!, model: text-similarity-davinci-001, token: 5     | text: Hallo verden!, model: text-similarity-davinci-001, token: 5     |