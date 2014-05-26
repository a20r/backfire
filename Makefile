
all: tests

tests:
	@cd backfire; go test; cd ..;
