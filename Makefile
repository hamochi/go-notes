.PHONY: run
run:
	rm -rf docs/*
	cp template/style.css docs/style.css
	go run go-notes/generator
