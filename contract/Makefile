.PHONY: compile deploy-shard0 deploy-shard1 clean
compile:
	npm run compile

deploy-shard0:
	npm run migrate-local-shard0

deploy-shard1:
	npm run migrate-local-shard1

clean:
	rm -rf node_modules build
