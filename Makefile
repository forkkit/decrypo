BINARY=decrypo
GITHUB_REPO="ajdnik/decrypo"
VERSION="0.1.0"
BUILD=`date +%FT%T%z`
LDFLAGS=-ldflags "-s -w"

clean:
	@rm -rf build/
	@rm -rf dist/

changelog:
	git-chglog -c .chglog/changelog/config.yml -o CHANGELOG.md --next-tag ${VERSION} ..${VERSION}

devdeps:
	go get -u github.com/git-chglog/git-chglog/cmd/git-chglog
	go get -u github.com/mitchellh/gox
	go get -u github.com/c4milo/github-release

compile:
	@rm -rf build/
	dep ensure
	@gox ${LDFLAGS} \
		-osarch="darwin/amd64" \
		-osarch="windows/amd64" \
		-output "build/{{.Dir}}_{{.OS}}_{{.Arch}}/$(BINARY)" \
		./...

dist: compile
	$(eval FILES := $(shell ls build))
	@rm -rf dist && mkdir dist
	@for f in $(FILES); do \
		(cd $(shell pwd)/build/$$f && tar -cvzf ../../dist/$$f.tar.gz *); \
		(cd $(shell pwd)/dist && shasum -a 512 $$f.tar.gz > $$f.sha512); \
		echo $$f; \
	done

release: dist changelog
	git add CHANGELOG.md
	git commit -m "chore: updated changelog"
	git add Makefile
	git commit -m "chore: version bumped"
	git push
	git-chglog -c .chglog/release/config.yml -o RELEASE.md --next-tag ${VERSION} ${VERSION}
	github-release $(GITHUB_REPO) $(VERSION) "$$(git rev-parse --abbrev-ref HEAD)" "## Changelog<br/>$$(cat RELEASE.md)" 'dist/*'
	@rm RELEASE.md
	git pull

default: changelog

.PHONY: dist release changelog compile devdeps clean
