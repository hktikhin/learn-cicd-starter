# Install bootdev cli
go install github.com/bootdotdev/bootdev@latest

go install github.com/securego/gosec/v2/cmd/gosec@latest

curl -O https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-cli-linux-x86_64.tar.gz

tar -xf google-cloud-cli-linux-x86_64.tar.gz

gcloud builds submit --tag us-central1-docker.pkg.dev/notely-496807/notely-ar-repo/notely:latest .

turso auth login --headless

go install github.com/pressly/goose/v3/cmd/goose@latest