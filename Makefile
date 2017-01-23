all: squidbot.pb.go

squidbot.pb.go: squidbot.proto
	protoc squidbot.proto --go_out=plugins=grpc:. squidbot.proto
	@echo -------- protos regenerated --------

push: squidbot.pb.go
	git add squidbot.pb.go
	git commit -m 'Updating regenerated files.'
	git push

.PHONY: all push
