import sys
import hashlib

if len(sys.argv) != 2:
	print("[!]File not found.")
	sys.exit()

filename = sys.argv[1]

def md5(string):
	return hashlib.md5(string).hexdigest()

def sha1(string):
	return hashlib.sha1(string).hexdigest()

def sha256(string):
	return hashlib.sha256(string).hexdigest()

def sha512(string):
	return hashlib.sha512(string).hexdigest()

fi = open(filename, 'rb')
data = fi.read()

final_md5_hash = md5(data)
final_sha1_hash = sha1(data)
final_sha256_hash = sha256(data)
final_sha512_hash = sha512(data)

print("MD5 --> {0} : {1}".format(filename, final_md5_hash))
print("SHA1 --> {0} : {1}".format(filename, final_sha1_hash))
print("SHA256 --> {0} : {1}".format(filename, final_sha256_hash))
print("SHA512 --> {0} : {1}".format(filename, final_sha512_hash))