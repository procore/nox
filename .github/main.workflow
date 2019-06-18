workflow "Release" {
  on = "push"
  resolves = ["goreleaser", "fmt", "docs"]
}

action "is-tag" {
  uses = "actions/bin/filter@master"
  args = "tag"
}

action "goreleaser" {
  uses = "docker://goreleaser/goreleaser"
  secrets = [
    "GORELEASER_GITHUB_TOKEN",
  ]
  args = "release"
  needs = ["is-tag"]
}

action "docs" {
  uses = "cedrickring/golang-action@1.3.0"
  args = "make docs"
  env={
    PROJECT_PATH = "./cmd/nox"
  }
}

action "fmt" {
  uses = "cedrickring/golang-action@1.3.0"
  args = "go fmt ."
}
