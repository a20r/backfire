
all: tests

tests:
	@cd backfire; go test -v; cd ..;
