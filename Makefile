test-all:
	bazel test //...

build-binary:
	bazel build //:sample-go-bzl-ch

run-binary:
	bazel run //:sample-go-bzl-ch -- $(app-arg)

build-and-load-oci-image:
	bazel build //:oci-img-tarball
	docker load --input $$(bazel cquery --output=files :oci-img-tarball)

run-app-in-docker:
	make build-and-load-oci-image
	docker compose --profile be up -d

run-fullstack-in-docker:
	make build-and-load-oci-image
	docker compose --profile fullstack up -d

generate-rule:
	bazel run //:gazelle

clean:
	echo "Don't try to clean everywhere u psycho!"
