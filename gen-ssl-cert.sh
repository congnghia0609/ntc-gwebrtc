#!/usr/bin/env bash

echo "Creating SSL Certificates needed for signal and web"
openssl req  -nodes -new -x509 -keyout localhost.key -out localhost.pem -days 365 -subj '/CN=localhost'

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cp $DIR/*.key $DIR/*.pem $DIR/signal
cp $DIR/*.key $DIR/*.pem $DIR/web
rm $DIR/*.key $DIR/*.pem
