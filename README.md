# tiktoken-go
Sadly, openai tiktoken is impossible to rewrite it in go.   
Golang doesn't support Regex negative look behinds(?!), which are required for this to work.    
I've done some work on this, maybe it's helpful for someone.

