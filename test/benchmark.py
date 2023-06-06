import tiktoken as tk
import requests
import time

def benchmark_test(text_list,enc):
    """
    Benchmark test
    :return: None
    """
    start = time.perf_counter_ns()
    for index in range(100000):
        text = text_list[index]
        num_tokens = len(enc.encode(text))
    end = time.perf_counter_ns()
    print('benchmark test: {} ns/op'.format((end - start)/100000))

if __name__ == '__main__':
    r = requests.get('https://unicode.org/udhr/assemblies/full_all.txt')
    text_list = r.text.splitlines()
    cursor = 0
    enc=tk.get_encoding('cl100k_base')
    benchmark_test(text_list,enc)

