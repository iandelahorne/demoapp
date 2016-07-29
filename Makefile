all: deb

prep:
	mkdir -p build/opt/twitch/demoapp

demoapp:
	go build -ldflags "-X main.VERSION=`cat VERSION`" -o ./build/opt/twitch/demoapp/demoapp demoapp.go

deb: prep demoapp
	fpm -t deb -n demoapp --iteration 1 -v `cat VERSION` --deb-upstart demoapp.conf --after-install deb-post-install.sh -p build -s dir -C build
