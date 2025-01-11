## Instructions to Setup your certificates

- Run the "make ca-cert" to create your CA certificates if you dont have one. If you have one, simply paste them as CA.crt and CA.key in the certs directory
- Run the "make certs" make target to create client and server certificates.Ensure that the SERVER_NAME environment variable is set. The server certificates will use that as a hostname
