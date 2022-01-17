go test -mod=vendor ./... -cover -coverprofile out.cover && 
go tool cover -html=out.cover

find . -type f \
    -name '*.go' \
    -exec sed -i -e 's,github.com/my/project,github.com/my/project/v2,g' {} \;