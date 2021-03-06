ANTLR_VERSION=4.8
ANTLR_JAR=./hack/antlr-$(ANTLR_VERSION)-complete.jar
ANTLR=java -Xmx500M -cp $(ANTLR_JAR) org.antlr.v4.Tool
GEN_PARSER=java -cp $(ANTLR_JAR) -Xmx500M org.antlr.v4.Tool -visitor -no-listener -Werror -Xexact-output-dir spec/RCL.g4

.PHONY: dep-download
dep-download:
	cd script; wget http://www.antlr.org/download/antlr-$(ANTLR_VERSION)-complete.jar

.PHONY: antlr
antlr:
	java -cp $(ANTLR_JAR) -Xmx500M org.antlr.v4.Tool

.PHONY: gen-parser
gen-parser:
	$(GEN_PARSER) -package org.reika.rcl.parser -o impl/rcl-j/src/main/java/org/reika/rcl/parser
# NOTE: go version now use hand written parser
#	$(GEN_PARSER) -package parser -o impl/rcl-go/rcl/parser -Dlanguage=Go

