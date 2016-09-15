PWD := $(shell pwd)
APP := app

dev:
	docker run --rm -it \
	 -p 8080:8080 \
	 -v $(PWD):/opt \
	 -w /opt/$(APP) \
	 ckeyer/dev:node bash

build:
	docker run --rm -it \
	 -v $(PWD):/opt \
	 -w /opt/$(APP) \
	 ckeyer/dev:node sh -c 'rm -rf ../static && node build/build.js'
