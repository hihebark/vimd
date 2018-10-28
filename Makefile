TARGET=pickle

all: clean build install clean

build:
	@echo "+ Building ..."
	@go build -o $(TARGET) .

clean:
	@echo "+ Cleaning ..."
	@rm -rf $(TARGET)
	@rm -rf build

install:
ifneq ($(wildcard $(TARGET)),)
	@go build -o $(TARGET) .
endif
	@echo "+ Installing ..."
	@cp $(TARGET) ${GOPATH}/bin/
