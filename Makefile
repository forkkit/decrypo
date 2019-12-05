BINARY=decrypo
GITHUB_REPO="ajdnik/decrypo"
VERSION="0.2.2"
TIME=`date +%FT%T%z`
LDFLAGS=-ldflags "-s -w -X github.com/${GITHUB_REPO}/build.version=${VERSION} -X github.com/${GITHUB_REPO}/build.datetime=${TIME}"

clean:
	@rm -rf dist/

prepare:
	mkdir -p dist/gox
	mkdir -p dist/arch

changelog:
	git-chglog -c .chglog/changelog/config.yml -o CHANGELOG.md --next-tag ${VERSION} ..${VERSION}

devdeps:
	go get -u github.com/git-chglog/git-chglog/cmd/git-chglog
	go get -u github.com/mitchellh/gox
	go get -u github.com/c4milo/github-release

compile: clean prepare
	dep ensure
	@gox ${LDFLAGS} \
		-osarch="darwin/amd64" \
		-osarch="windows/amd64" \
		-output "dist/gox/{{.Dir}}_{{.OS}}_{{.Arch}}/$(BINARY)" \
		./...

dist: compile
	$(eval FILES := $(shell ls dist/gox))
	@for f in $(FILES); do \
		(cd $(shell pwd)/dist/gox/$$f && tar -cvzf ../../arch/$$f.tar.gz *); \
		(cd $(shell pwd)/dist/arch && shasum -a 256 $$f.tar.gz > $$f.sha256); \
		echo $$f; \
	done

release: dist changelog
	git add CHANGELOG.md
	git commit -m "chore: updated changelog"
	git add Makefile
	git commit -m "chore: version bumped"
	git push
	git-chglog -c .chglog/release/config.yml -o RELEASE.md --next-tag ${VERSION} ${VERSION}
	github-release $(GITHUB_REPO) $(VERSION) "$$(git rev-parse --abbrev-ref HEAD)" "## Changelog<br/>$$(cat RELEASE.md)" 'dist/arch/*'
	@rm RELEASE.md
	git pull

default: changelog

.PHONY: dist release changelog compile devdeps clean prepare
