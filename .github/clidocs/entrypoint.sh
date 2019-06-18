#!/bin/sh

set -eu

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y git

git config --global user.email "$GITHUB_EMAIL"
git config --global user.name "$GITHUB_USER"
git submodule add https://github.com/procore/nox.wiki.git ./cmd/nox/docs

cd ./cmd/nox
go fmt .
make docs

cd ./docs
git add . --force
git status
git commit -m "Update auto-generated documentation."
git push --set-upstream origin master

cd ../
make clean
