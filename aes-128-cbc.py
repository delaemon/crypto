import os
import hashlib

from Crypto.Cipher import AES

def padding(data):
  pad = AES.block_size - (len(data) % AES.block_size) + len(data)
  out = data.ljust(pad, ' ')
  return out

def encrypt(key, iv, plain_text):
  data = padding(plain_text)
  encryptor = AES.new(key, AES.MODE_CBC, iv)
  encrypted = encryptor.encrypt(data)
  return encrypted

def decrypt(key, iv, encrypted):
  decryptor = AES.new(key, AES.MODE_CBC, iv)
  decrypted = decryptor.decrypt(encrypted)
  return decrypted

def main():
  key = hashlib.sha256('example no key death.').digest()
  iv = os.urandom(AES.block_size)

  plain_text = 'example no plain text death.'
  encrypted = encrypt(key, iv, plain_text)
  decrypted = decrypt(key, iv, encrypted)

  print 'encrypt: %s' % (encrypted.encode('hex'))
  print 'decrypt: %s' % (decrypted)

if __name__ == "__main__":
  main()
