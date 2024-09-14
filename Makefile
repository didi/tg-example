all: install

install: clean gen_output build
	mkdir -p output/bin
	mv $(GOPATH)/src/github.com/didi/tg-example/tg-example output/bin/

gen_output:
	mkdir -p output/log
	mkdir -p output/conf
	cp -R conf/* output/conf

	cp -R template output/
	cp -R public output/
	
	cp control.sh output/

clean:
	rm -rf output/bin
	rm -rf output/conf
	rm -rf output/control.sh

build:
	go build

run:
	./output/bin/tg-example # 前台运行,方便退出
