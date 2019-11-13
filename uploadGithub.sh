TIME=$(date +"%Y%m%d.%H%M%S")
VERSION=0.0.1-alpha
DEST=build/dist
TAG=$VERSION
DESCRIPTION="Automated Distribution"
RELEASENAME=$VERSION

echo "Remember to setup GITHUB_TOKEN env variable"
echo "Uploading to GITHUB"

github-release release \
    --user jpramirez \
    --repo epicFDA \
    --tag $TAG \
    --name "$RELEASENAME" \
    --description "$DESCRIPTION" \
    --pre-release

github-release upload \
    --user jpramirez \
    --repo epicFDA \
    --tag $TAG  \
    --name linux-dist.zip \
    --file $DEST/linux-dist.zip



