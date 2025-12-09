#!/bin/bash
# create-user.sh
USERNAME=$1
GROUP=$2
NAMESPACE=${3:-default}


openssl genrsa -out ${USERNAME}.key 2048
openssl req -new -key ${USERNAME}.key -out ${USERNAME}.csr -subj "/CN=${USERNAME}/O=${GROUP}"


openssl x509 -req -in ${USERNAME}.csr \
  -CA /etc/kubernetes/pki/ca.crt \
  -CAkey /etc/kubernetes/pki/ca.key \
  -CAcreateserial -out ${USERNAME}.crt -days 365


kubectl config set-credentials ${USERNAME} \
  --client-certificate=${USERNAME}.crt \
  --client-key=${USERNAME}.key

kubectl config set-context ${USERNAME}-context \
  --cluster=kubernetes \
  --namespace=${NAMESPACE} \
  --user=${USERNAME}

echo "Created user: ${USERNAME}"
echo "Switch with: kubectl config use-context ${USERNAME}-context"

