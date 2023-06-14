# Replace this with your own github.com/<username>/<repository>
GO_MODULE :=github.com/MichaelYcCho/protobuf-go

.PHONY: tidy
tidy:
	go mod tidy


.PHONY: clean
clean:
ifeq ($(OS),Windows_NT)
	if exist "protogen" rd /s /q "protogen"
else
	rm -rf protogen
endif


.PHONY: protoc
# (un)comment or add folder that contains proto as required
protoc:
	protoc --go_opt=module==module=${GO_MODULE} --go_out=. \
	./proto/basic/*.proto    

.PHONY: build
build: clean protoc tidy


.PHONY: run
run:
	go run main.go