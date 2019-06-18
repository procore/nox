FROM golang:1.12.4

LABEL "com.github.actions.name"="Generate Nox CLI Docs"
LABEL "com.github.actions.description"="Generates documentation for the nox CLI and pushes it to the project wiki"
LABEL "com.github.actions.icon"="book"
LABEL "com.github.actions.color"="red"

LABEL "repository"="https://github.com/procore/nox"
LABEL "homepage"="https://github.com/procore/nox"
LABEL "maintainer"="Maxine Krebs <maxine.krebs@procore.com>"

ADD entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]


