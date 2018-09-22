.PHONY : clean build

clean: 
	@echo "cleaning out directory..."
	@rm -fR out

depends:
	

build: clean
	@echo "building..."
	@go build -o bin/run app/main.go 
	@echo "done. make run to run"
run: build
	@echo "running..."
	@./bin/run
