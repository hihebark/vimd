TARGET=pickle

all: clean build install

build:
	@echo "+ Building ..."
	@go build -o $(TARGET) .

clean:
	@rm -rf $(TARGET)
	@rm -rf build

install:
	@echo "+ Installing ..."
	@cp $(TARGET) ${GOPATH}/bin/
