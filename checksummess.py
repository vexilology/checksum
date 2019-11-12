import hashlib

message = input("Enter the message: ")

a = hashlib.md5()
e = hashlib.sha1()
f = hashlib.sha256()
c = hashlib.sha512()

a.update(message.encode('utf-8'))
e.update(message.encode('utf-8'))
f.update(message.encode('utf-8'))
c.update(message.encode('utf-8'))

print("MD5: ", a.hexdigest())
print("SHA1: ", e.hexdigest())
print("SHA256: ", f.hexdigest())
print("SHA512: ", c.hexdigest())