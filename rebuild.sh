rm -rf docs *.go
openapi-generator generate --git-user-id hyperonecom --git-repo-id h1-client-go -g go -i ../rbx-api-public/openapi.json
