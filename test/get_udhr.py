import os
import tarfile
import urllib.request

url = "http://research.ics.aalto.fi/cog/data/udhr/udhr_txt_20100325.tar.gz"
file_name = "/tmp/udhr_txt_20100325.tar.gz"

def download_file(url, file_name):
    urllib.request.urlretrieve(url, file_name)

def merge_files(source_dir, output_file):
    with open(output_file, 'wb') as outfile:
        for filename in os.listdir(source_dir):
            if os.path.isfile(os.path.join(source_dir, filename)):
                with open(os.path.join(source_dir, filename), 'rb') as infile:
                    outfile.write(infile.read())
                outfile.write(b'\n')

def untar(dest, file_name):
    with tarfile.open(file_name, "r:gz") as tar:
        tar.extractall(path=dest)

# download_file(url, file_name)

untar('/tmp', file_name)

merge_files('/tmp/udhr/txt', '/tmp/udhr.txt')
