main:
	@echo Specify a target.

install-linux:
	go get github.com/go-gl/mathgl/mgl32
	go get github.com/DeedleFake/Go-PhysicsFS/physfs
	go get github.com/go-gl/glfw/v3.2/glfw
	go get github.com/go-gl/gl/v3.3-core/gl
	go get github.com/sheenobu/go-obj/obj
	
linux:
	@echo Compiling Engine data
	cd ../Data && zip -r -q public.zip *
	@echo Compiling Engine source
	GOARCH=amd64 go build -o build/Warnengine src/*.go
	@echo Launching app
	./build/Warnengine ../Data/public.zip

win:
	@echo Compiling Engine source
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -o build/Warnengine.exe src/*.go
	@echo Launching app
	wine build/Warnengine.exe ../data/public.zip