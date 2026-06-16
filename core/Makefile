pk:
	@openssl ecparam -name prime256v1 -genkey -noout -out tmp.pem
	@base64 -i tmp.pem -o "keys/$$(date -u +%Y-%m-%dT%H-%M-%S%Z)"
	@rm tmp.pem