var crypto = require('crypto'),
  algorithm = 'aes-256-gcm',
  password = 'passwordpasswordpasswordpassword',
  iv = 'iviviviv', // generate a new one for each encryption
  aad = new Buffer('aadaadaad')

function encrypt(text, data) {
  var cipher = crypto.createCipheriv(algorithm, password, iv)
  cipher.setAAD(aad)
  var encrypted = cipher.update(text, 'utf8', 'hex')
  encrypted += cipher.final('hex');
  var tag = cipher.getAuthTag();
  return {
    content: encrypted,
    tag: tag
  };
}

function decrypt(encrypted) {
  var decipher = crypto.createDecipheriv(algorithm, password, iv)
  decipher.setAuthTag(encrypted.tag);
  decipher.setAAD(aad)
  var dec = decipher.update(encrypted.content, 'hex', 'utf8')
  dec += decipher.final('utf8');
  return dec;
}

var data = encrypt("kore nande shokai")
console.log(data)
console.log(decrypt(data));
