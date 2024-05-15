import tiktoken as tk
import time

def benchmark_test(text_list,enc):
    """
    Benchmark test
    :return: None
    """
    start = time.perf_counter_ns()
    n = 100000
    for i in range(n):
        text = text_list[i % len(text_list)]
        num_tokens = len(enc.encode(text))
    end = time.perf_counter_ns()
    print('benchmark test: {} ns/op'.format((end - start)/n))

if __name__ == '__main__':
    with open('/tmp/udhr.txt','r') as f:
        text_list = f.readlines()
        enc=tk.get_encoding('o200k_base')
        benchmark_test(text_list,enc)
