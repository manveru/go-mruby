all : libmruby.so
	go build -x .

libmruby.so : mruby/lib/libmruby.a
	ld --whole-archive -shared -o libmruby.so mruby/build/host/lib/libmruby.a

mruby/lib/libmruby.a :
	(cd mruby && \
	CFLAGS="-fPIC -g -O3 -Wall -Werror-implicit-function-declaration" make)

clean :
	(cd mruby && make clean)
	go clean .
