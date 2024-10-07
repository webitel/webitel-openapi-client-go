#!/usr/bin/env bash

# https://stackoverflow.com/a/246128
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

# Pull the schema (last commit from main branch)
SCHEMA="$(curl -s -L https://raw.githubusercontent.com/webitel/protos/refs/heads/main/swagger/api.json)"

# Custom extensions: https://goswagger.io/use/models/schemas.html#custom-extensions
modify() {
    SCHEMA="$(echo "${SCHEMA}" | jq "${1}")"
}

modify '(.paths."/roles/metadata".get.parameters[] | select(.name == "merge.struct_value") | .type) |= "string"'

# Write the schema to a file
echo "${SCHEMA}" > "${SCRIPT_DIR}/schema.json"
