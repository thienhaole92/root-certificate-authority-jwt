# Setting up a Root Certificate Authority with OpenSSL

In secure communication establishing a secure channel is very important. One way to achieve this is by setting up a Root Certificate Authority (CA) to sign and manage digital certificates. In this article, I'll walk you through the process of creating your own Root CA and signing service certificates using OpenSSL, a versatile open-source tool for cryptography.

## Create the Root CA

The Root CA serves as the trust anchor for all certificates. Follow these steps to generate the Root CA key and certificate:

1. Generate the private key for the Root CA

```sh
openssl genrsa -out rootCA.key 4096
```

2. Create the Root CA certificate

```sh
openssl req -x509 -sha256 -new -nodes -key rootCA.key -days 3650 -out rootCA.crt
```

3. Review the Root CA certificate

```sh
openssl x509 -in rootCA.crt -text
```

## Create the Service Key and Certificate Signing Request (CSR)

Let's generate the key for your service and create a Certificate Signing Request (CSR)

1. Generate the private key for your service

```sh
openssl genrsa -out service.key 4096
```

2. Create a CSR for your service

```sh
openssl req -new -key service.key -out service.csr
```

## Sign the Service Certificate with the Root CA

Sign the CSR with the Root CA to produce a valid service certificate

1. Sign the CSR with the Root CA

```sh
openssl x509 -req -in service.csr -CA rootCA.crt -CAkey rootCA.key -CAcreateserial -out service.crt -days 1
```

## Verify the Certificates

To ensure the integrity and correctness of the generated certificates, use the following commands:

1. Verify the Root CA certificate

```sh
openssl x509 -noout -text -in rootCA.crt -enddate
```

2. Verify the Service certificate

```sh
openssl x509 -noout -text -in service.crt -enddate
```

## Export public key from certificate signing request

```sh
openssl req -in service.csr -noout -pubkey -out publickey.pem
```

Or

```sh
openssl x509 -pubkey -noout -in service.crt -out pubkey.pem
```

## RSA sign and verify using OpenSSL

1. Create sample data file

```sh
echo abcdefghijklmnopqrstuvwxyz > myfile.txt
```

2. Sign the file using sha1 digest and PKCS1 padding scheme

```sh
openssl dgst -sha1 -sign service.key -out sha1.sign myfile.txt
```

3. Verify the signature of file

```sh
openssl dgst -sha1 -verify pubkey.pem -signature sha1.sign myfile.txt
```
