all: container

container:
	docker build . -t arbi:public

run:
	docker run -P arbi:public
