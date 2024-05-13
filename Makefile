test-all:
	bazel test //:test

build-binary:
	bazel build //:sample-go-bzl-ch

build-oci-image:
	bazel build //:oci_image

build-docker-image:
	docker build -t sample-go-bzl-ch:$(img_version) .

run-app-in-docker:
	docker compose --profiles be up -d

run-fullstack-in-docker:
	docker compose --profiles fullstack up -d

generate-rule:
	bazel run //:gazelle
	bazel run //:gazelle_update_repos

clean:
	echo "Don't try to clean everywhere u psycho!"
