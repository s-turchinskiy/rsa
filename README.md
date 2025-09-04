
# Data encryption/decryption using RSA

## encryption:
rsa encrypt
&ensp;&ensp;-msg string
&ensp;&ensp;&ensp;&ensp; message to encode
&ensp;&ensp;-pub string
&ensp;&ensp;&ensp;&ensp;path to file with public key

Example for linux:
```
./rsa decrypt -msg "Hello" -pr "./privatekey.txt"
output: TMOEnVBrqrqnWQUQdThikk9qQaMSqkJTDBlilbB6XQT10fSRt9a7YLTbMYPU/0inBT2Lw5WhHiClKoIWbQNkUE5DOkcqUnTwEaEsW5cQ83OACF/SbVZq5Wn4xqd8bryiup/IoyjylQgKMEpflUbR9lCOtqurkiKBNFPxiPZhWH9U+oVc3/osqfmHDRPfbERh7hkkBI1sL+LOMdcGN3T6+5UIvzrGJ3T3lMQ2Ti5S8OZTsFhrdqsFFSdVNab8eB+LrwOIFuSjpl05RayNo5GTakqiYylYyxCh4PobVUHl7tApCaKydvnV0Mt9wxMZv92ewrbCgnpZmu66y9WtykXcRQ==
```
Example for windows:
```
rsa64.exe decrypt -msg "Hello" -pr "./privatekey.txt"
output: TMOEnVBrqrqnWQUQdThikk9qQaMSqkJTDBlilbB6XQT10fSRt9a7YLTbMYPU/0inBT2Lw5WhHiClKoIWbQNkUE5DOkcqUnTwEaEsW5cQ83OACF/SbVZq5Wn4xqd8bryiup/IoyjylQgKMEpflUbR9lCOtqurkiKBNFPxiPZhWH9U+oVc3/osqfmHDRPfbERh7hkkBI1sL+LOMdcGN3T6+5UIvzrGJ3T3lMQ2Ti5S8OZTsFhrdqsFFSdVNab8eB+LrwOIFuSjpl05RayNo5GTakqiYylYyxCh4PobVUHl7tApCaKydvnV0Mt9wxMZv92ewrbCgnpZmu66y9WtykXcRQ==
```
## decryption:
rsa decrypt
&ensp;&ensp;-msg string
&ensp;&ensp;&ensp;&ensp; message to decode
&ensp;&ensp;-pr string
&ensp;&ensp;&ensp;&ensp;path to file with private key

Example for linux:
```
./rsa decrypt -pr "./privatekey.txt" -msg "TMOEnVBrqrqnWQUQdThikk9qQaMSqkJTDBlilbB6XQT10fSRt9a7YLTbMYPU/0inBT2Lw5WhHiClKoIWbQNkUE5DOkcqUnTwEaEsW5cQ83OACF/SbVZq5Wn4xqd8bryiup/IoyjylQgKMEpflUbR9lCOtqurkiKBNFPxiPZhWH9U+oVc3/osqfmHDRPfbERh7hkkBI1sL+LOMdcGN3T6+5UIvzrGJ3T3lMQ2Ti5S8OZTsFhrdqsFFSdVNab8eB+LrwOIFuSjpl05RayNo5GTakqiYylYyxCh4PobVUHl7tApCaKydvnV0Mt9wxMZv92ewrbCgnpZmu66y9WtykXcRQ=="
output: Hello
```
Example for windows:
```
rsa64.exe decrypt -pr "./privatekey.txt" -msg "TMOEnVBrqrqnWQUQdThikk9qQaMSqkJTDBlilbB6XQT10fSRt9a7YLTbMYPU/0inBT2Lw5WhHiClKoIWbQNkUE5DOkcqUnTwEaEsW5cQ83OACF/SbVZq5Wn4xqd8bryiup/IoyjylQgKMEpflUbR9lCOtqurkiKBNFPxiPZhWH9U+oVc3/osqfmHDRPfbERh7hkkBI1sL+LOMdcGN3T6+5UIvzrGJ3T3lMQ2Ti5S8OZTsFhrdqsFFSdVNab8eB+LrwOIFuSjpl05RayNo5GTakqiYylYyxCh4PobVUHl7tApCaKydvnV0Mt9wxMZv92ewrbCgnpZmu66y9WtykXcRQ=="
output: Hello
```
## decryption of multiple messages:
*The messages are in a file and are separated by a newline character.*
*Output: the encrypted  messages are in  the  same  order  as  in the file.*

rsa decryptbatch
&ensp;&ensp;-data string
&ensp;&ensp;&ensp;&ensp;path to file with data
&ensp;&ensp;-pr string
&ensp;&ensp;&ensp;&ensp;path to file with private key

Example for linux:
```
./rsa decryptbatch -data "data.txt" -pr "./privatekey.txt"
```
Example for windows:
```
rsa64.exe decryptbatch -data "data.txt" -pr "./privatekey.txt"
```

# Шифрование/расшифровка данных с помощью RSA

## Шифрование:
rsa encrypt
&ensp;&ensp;-msg string
&ensp;&ensp;&ensp;&ensp; текст для шифрования
&ensp;&ensp;-pub string
&ensp;&ensp;&ensp;&ensp; путь к файлу с публичным ключом

Пример для linux:
```
./rsa decrypt -msg "Hello" -pr "./privatekey.txt"
output: TMOEnVBrqrqnWQUQdThikk9qQaMSqkJTDBlilbB6XQT10fSRt9a7YLTbMYPU/0inBT2Lw5WhHiClKoIWbQNkUE5DOkcqUnTwEaEsW5cQ83OACF/SbVZq5Wn4xqd8bryiup/IoyjylQgKMEpflUbR9lCOtqurkiKBNFPxiPZhWH9U+oVc3/osqfmHDRPfbERh7hkkBI1sL+LOMdcGN3T6+5UIvzrGJ3T3lMQ2Ti5S8OZTsFhrdqsFFSdVNab8eB+LrwOIFuSjpl05RayNo5GTakqiYylYyxCh4PobVUHl7tApCaKydvnV0Mt9wxMZv92ewrbCgnpZmu66y9WtykXcRQ==
```
Пример для windows:
```
rsa64.exe decrypt -msg "Hello" -pr "./privatekey.txt"
output: TMOEnVBrqrqnWQUQdThikk9qQaMSqkJTDBlilbB6XQT10fSRt9a7YLTbMYPU/0inBT2Lw5WhHiClKoIWbQNkUE5DOkcqUnTwEaEsW5cQ83OACF/SbVZq5Wn4xqd8bryiup/IoyjylQgKMEpflUbR9lCOtqurkiKBNFPxiPZhWH9U+oVc3/osqfmHDRPfbERh7hkkBI1sL+LOMdcGN3T6+5UIvzrGJ3T3lMQ2Ti5S8OZTsFhrdqsFFSdVNab8eB+LrwOIFuSjpl05RayNo5GTakqiYylYyxCh4PobVUHl7tApCaKydvnV0Mt9wxMZv92ewrbCgnpZmu66y9WtykXcRQ==
```
## Расшифровка:
rsa decrypt
&ensp;&ensp;-msg string
&ensp;&ensp;&ensp;&ensp; зашифрованный текст
&ensp;&ensp;-pr string
&ensp;&ensp;&ensp;&ensp;путь к файлу с приватным ключом

Пример для linux:
```
./rsa decrypt -pr "./privatekey.txt" -msg "TMOEnVBrqrqnWQUQdThikk9qQaMSqkJTDBlilbB6XQT10fSRt9a7YLTbMYPU/0inBT2Lw5WhHiClKoIWbQNkUE5DOkcqUnTwEaEsW5cQ83OACF/SbVZq5Wn4xqd8bryiup/IoyjylQgKMEpflUbR9lCOtqurkiKBNFPxiPZhWH9U+oVc3/osqfmHDRPfbERh7hkkBI1sL+LOMdcGN3T6+5UIvzrGJ3T3lMQ2Ti5S8OZTsFhrdqsFFSdVNab8eB+LrwOIFuSjpl05RayNo5GTakqiYylYyxCh4PobVUHl7tApCaKydvnV0Mt9wxMZv92ewrbCgnpZmu66y9WtykXcRQ=="
output: Hello
```
Пример для windows:
```
rsa64.exe decrypt -pr "./privatekey.txt" -msg "TMOEnVBrqrqnWQUQdThikk9qQaMSqkJTDBlilbB6XQT10fSRt9a7YLTbMYPU/0inBT2Lw5WhHiClKoIWbQNkUE5DOkcqUnTwEaEsW5cQ83OACF/SbVZq5Wn4xqd8bryiup/IoyjylQgKMEpflUbR9lCOtqurkiKBNFPxiPZhWH9U+oVc3/osqfmHDRPfbERh7hkkBI1sL+LOMdcGN3T6+5UIvzrGJ3T3lMQ2Ti5S8OZTsFhrdqsFFSdVNab8eB+LrwOIFuSjpl05RayNo5GTakqiYylYyxCh4PobVUHl7tApCaKydvnV0Mt9wxMZv92ewrbCgnpZmu66y9WtykXcRQ=="
output: Hello
```
## Расшифровка нескольких сообщений:
*Сообщения  хранятся  в  файле  и  разделяются  символом новой строки.*
*Вывод: зашифрованные сообщения в том же порядке, что и в файле.*

rsa decryptbatch
&ensp;&ensp;-data string
&ensp;&ensp;&ensp;&ensp;путь к файлу с зашифрованными данными
&ensp;&ensp;-pr string
&ensp;&ensp;&ensp;&ensp;путь к файлу с приватным ключом

Пример для linux:
```
./rsa decryptbatch -data "data.txt" -pr "./privatekey.txt"
```
Пример для windows:
```
rsa64.exe decryptbatch -data "data.txt" -pr "./privatekey.txt"
```







