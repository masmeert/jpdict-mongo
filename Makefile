run:
	go run ./cmd/jpdict-mongo/main.go

test:
	@echo "Running tests..."
	go test -v ./...

dl-all: dl-jmdict dl-kanjidic

dl-jmdict:
	@echo "Downloading  JMDict..."
	wget -P ./assets http://ftp.edrdg.org/pub/Nihongo/JMdict_e.gz
	gunzip ./assets/JMdict_e.gz
	mv ./assets/JMdict_e ./assets/JMdict.xml

dl-kanjidic:
	@echo "Downloading KANJIDIC..."
	wget -P ./assets http://ftp.edrdg.org/pub/Nihongo/kanjidic2.xml.gz
	gunzip ./assets/kanjidic2.xml.gz
