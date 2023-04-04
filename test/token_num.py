import tiktoken as tk

def read_data_from_file(filename):
    """
    Read data from file
    :param filename: file name
    :return: text_list, model_list, encoding_list
    """
    with open(filename, 'r', encoding='utf-8') as f:
        data = f.read()
        text_list = data.splitlines()[0].split(',')
        model_list = data.splitlines()[1].split(',')
        encoding_list = data.splitlines()[2].split(',')
    
    return text_list, model_list, encoding_list

def get_token_by_model(text, model):
    """
    Get token by model
    :param text: text
    :param model: model
    :return: num_tokens
    """
    encoding = tk.encoding_for_model(model)
    num_tokens = len(encoding.encode(text))
    return num_tokens

def get_token_by_encoding(text, encoding_name):
    """
    Get token by encoding
    :param text: text
    :param encoding: encoding
    :return: num_tokens
    """
    encoding = tk.get_encoding(encoding_name)
    num_tokens = len(encoding.encode(text))
    return num_tokens


def test_token_by_model(text_list, model_list):
    """
    Test token by model
    :param text_list: text list
    :param model_list: model list
    :return: None
    """
    for text in text_list:
        for model in model_list:
            num_tokens = get_token_by_model(text, model)
            print('text: {}, model: {}, token: {}'.format(text, model, num_tokens))

def test_token_by_encoding(text_list, encoding_list):
    """
    Test token by encoding
    :param text_list: text list
    :param encoding_list: encoding list
    :return: None
    """
    for text in text_list:
        for encoding in encoding_list:
            num_tokens = get_token_by_encoding(text, encoding)
            print('text: {}, encoding: {}, token: {}'.format(text, encoding, num_tokens))

if __name__ == '__main__':
    text_list, model_list, encoding_list = read_data_from_file('test/test.txt')
    test_token_by_model(text_list, model_list)
    print("=====================================")
    test_token_by_encoding(text_list, encoding_list)

    