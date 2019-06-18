workflow "Release" {
  on = "push"
  resolves = ["goreleaser"]
}

workflow "Build and Publish CLI Docs" {
  on = "release"
  resolves = ["cli-docs"]
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

action "cli-docs" {
  uses = "./clidocs"
  secrets = [
    "GITHUB_TOKEN",
    "GH_USER",
    "GH_EMAIL",
  ]
}
