SERVER_CERTS_PATH="../certs/server"
CLIENT_CERTS_PATH="../certs/client"

generateSelfSignedCert() {
    CERT_PATH=$1

    mkdir -p $CERT_PATH
    openssl genrsa 2048 > $CERT_PATH/key.pem
    openssl req -new -x509 -nodes -key $CERT_PATH/key.pem -out $CERT_PATH/cert.pem -subj "/C=BR/ST=Espirito Santo/L=Vila Velha/CN=localhost"

    # rename
    mv $CERT_PATH/key.pem $CERT_PATH/ca.key
    mv $CERT_PATH/cert.pem $CERT_PATH/ca.crt
}

generateCASignedCert() {
    CA_CERT_PATH=$1
    CERT_PATH=$2
    NAME=$3

    mkdir -p $CERT_PATH
    openssl req -newkey rsa:2048 -nodes -keyout $CERT_PATH/key.pem -out $CERT_PATH/req.pem -subj "/C=BR/ST=Espirito Santo/L=Vila Velha/CN=localhost"

    openssl x509 -req  -set_serial 01 -in $CERT_PATH/req.pem -out $CERT_PATH/cert.pem \
    -CA $CA_CERT_PATH/ca.crt -CAkey $CA_CERT_PATH/ca.key

    rm $CERT_PATH/req.pem

    # rename
    mv $CERT_PATH/key.pem $CERT_PATH/$NAME.key
    mv $CERT_PATH/cert.pem $CERT_PATH/$NAME.crt
}

echo "Generating server CA cert"
generateSelfSignedCert $SERVER_CERTS_PATH/ca 2>/dev/null

echo "Generating server certs"
generateCASignedCert $SERVER_CERTS_PATH/ca $SERVER_CERTS_PATH "server" 2>/dev/null

