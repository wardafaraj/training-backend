# Generation of Keys
 - Generate private and public key using openssl as shown below
 - private key: `openssl ecparam -name prime256v1 -genkey -noout -out crweb_private_key.pem`
 - public key: `openssl pkey -in crweb_private_key.pem -pubout -out crweb_public_key.pem`
 - copy these keys into .storage/keys
 - specify the path of these keys into the config.yaml file