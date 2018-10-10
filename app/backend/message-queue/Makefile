export MESSAGE_QUEUE_IMAGE_TAG=0.1-DEV

build:
	mypy build-image.py
	python build-image.py

run:
	docker run -d --name mq --hostname mq \
		--network=mimir-net -p 5672:5672 -p 15672:15672 \
		eu.gcr.io/mimir-185212/message-queue:$(MESSAGE_QUEUE_IMAGE_TAG)

run-producer:
	mypy --ignore-missing-imports clients/producer.py
	python clients/producer.py

run-worker:
	mypy --ignore-missing-imports clients/worker.py
	python clients/worker.py
