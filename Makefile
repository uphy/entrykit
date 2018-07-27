PACKAGE_OS := linux darwin
PACKAGE_ARCH := amd64 386

build:
	go build .

package: clean
	go get github.com/mitchellh/gox && \
	mkdir -p build && \
	cd cmd && \
	gox -os="$(PACKAGE_OS)" -arch="$(PACKAGE_ARCH)" -output="../build/entrykit_{{.OS}}_{{.Arch}}/entrykit" && \
	cd - && \
	mkdir -p dist && \
	ls -1 build | xargs -I% tar zcf "dist/%.tar.gz" -C build "%"

clean:
	rm -rf build dist