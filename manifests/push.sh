#!/bin/sh

FILENAME="bundles"
NAMESPACE="akashem"
REPOSITORY="bundles"
RELEASE="1.7.1"

# Remove leading and trailing double quotes, if there are any.
TOKEN="${TOKEN%\"}"
TOKEN="${TOKEN#\"}"

function cleanup() {
    rm -f ${FILENAME}.tar.gz
    rm -f ${FILENAME}.encoded
}
# trap cleanup EXIT

tar czf ${FILENAME}.tar.gz etcd prometheus

BLOB=$(cat ${FILENAME}.tar.gz | base64 -w 0)
BODY='
{
    "blob": "'"${BLOB}"'",
    "release": "'"${RELEASE}"'",
    "media_type": "helm"
}'
echo -n "${BODY}" > "${FILENAME}.encoded"

curl -H "Content-Type: application/json" -H "Authorization: ${TOKEN}" -XPOST https://quay.io/cnr/api/v1/packages/${NAMESPACE}/${REPOSITORY} -d @${FILENAME}.encoded
