go-lint:
	@sh ./scripts/go-lint.sh	

go-run:
	@sh ./scripts/go-run.sh

go-update:
	@sh ./scripts/go-update.sh

go-image-build:
	@sh ./scripts/go-image-build.sh

go-image-run:
	@sh ./scripts/go-image-run.sh

go-image-build-run:
	@sh ./scripts/go-image-build.sh
	@sh ./scripts/go-image-run.sh

deploy-development:
	git push origin master && git push origin development

deploy-staging:
	git push origin master && git push origin master:staging

deploy-all:
	@sh ./scripts/deploy-all.sh

swag-init:
	swag init --parseInternal --exclude build,developments,docs,scripts,vendor -g cmd/article-service/main.go 

swag-fmt:
	swag fmt --exclude build,developments,docs,scripts -g cmd/article-service/main.go 
