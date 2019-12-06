GITHUB_REPO="ajdnik/decrypo"
VERSION="0.2.4"

changelog:
	git-chglog -c .chglog/changelog/config.yml -o CHANGELOG.md --next-tag ${VERSION} ..${VERSION}

devdeps:
	go get -u github.com/git-chglog/git-chglog/cmd/git-chglog
	go get -u github.com/c4milo/github-release

release: changelog
	git add CHANGELOG.md
	git commit -m "chore: updated changelog"
	git add Makefile
	git commit -m "chore: version bumped"
	git push
	git-chglog -c .chglog/release/config.yml -o RELEASE.md --next-tag ${VERSION} ${VERSION}
	github-release $(GITHUB_REPO) $(VERSION) "$$(git rev-parse --abbrev-ref HEAD)" "## Changelog<br/>$$(cat RELEASE.md)"
	@rm RELEASE.md
	git pull

default: changelog

.PHONY: dist release changelog compile devdeps clean prepare
