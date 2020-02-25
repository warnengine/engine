main:
	@echo Specify a target.
	
linux:
	@echo Compiling Engine data
	cd ../Data && zip -r -q public.zip *
	@echo Compiling Engine source
	GOARCH=amd64 go build -o build/Warnengine .
	@echo Launching app
	./build/Warnengine ../Data/

win:
	@echo Compiling Engine source
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -o build/Warnengine.exe src/*.go
	@echo Launching app
	wine build/Warnengine.exe ../Data/
