docker pull openapitools/openapi-generator-cli

docker run --rm \
    -v ${PWD}:/local openapitools/openapi-generator-cli generate \
    -i https://api.hyperone.com/openapi.json \
    -g go \
    -o /local
