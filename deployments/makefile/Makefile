setup:
	if [ -e ./dist ]; then rm -rf ./dist; fi; mkdir ./dist; mkdir -p ./dist/darwin/bin; mkdir -p ./dist/linux/bin

build:
	@echo "******************"
	@echo "**** Building ****"
	@echo "******************"

	GO111MODULE=on GOOS=linux go build -mod=vendor ../../cmd/authservice_bootup.go;cp dist/linux/bin/authservice_bootup ../docker;mv ./authservice_bootup dist/linux/bin/authservice_bootup


docker-build-prep:
	cd docker;cp ../dist/linux/bin/authservice_bootup .

docker-build-cleanup:
	cd ../docker;rm authservice_bootup

release-tag: build

	cd ../; git tag `cat .version`
	git push --tags

release: release-tag
	cd ../docker; docker build -t "singaravelan21/todo_list_auth_srv" . --no-cache;
	cd ../docker; docker tag "singaravelan21/todo_list_auth_srv" "singaravelan21/todo_list_auth_srv:`cat ../.version`";
	cd ../docker; docker push "singaravelan21/todo_list_auth_srv:`cat ../.version`"
	make docker-build-cleanup

