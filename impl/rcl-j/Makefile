ANTLR_VERSION=4.7.1
ANTLR_JAR=./script/antlr-$(ANTLR_VERSION)-complete.jar
ANTLR=java -Xmx500M -cp $(ANTLR_JAR) org.antlr.v4.Tool
GEN_PARSER=java -cp $(ANTLR_JAR) -Xmx500M org.antlr.v4.Tool -visitor -no-listener -Werror -Xexact-output-dir
PLAYGROUND=src/main/java/org/reika/rcl/playground

.PHONY: test
test:
	gradle test

.PHONY: gen-yaml
gen-yaml:
	$(GEN_PARSER) -package org.reika.rcl.yaml $(PLAYGROUND)/yaml/YAML.g4 -o $(PLAYGROUND)/yaml

.PHONY: dep-download
dep-download:
	cd script; wget http://www.antlr.org/download/antlr-$(ANTLR_VERSION)-complete.jar

# --- docker --- #
.PHONY: docker-test
docker-test:
	docker-compose run --rm j8
	docker-compose run --rm j9
# --- docker --- #