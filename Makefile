ANTLR_VERSION=4.7.1
ANTLR_JAR=./script/antlr-$(ANTLR_VERSION)-complete.jar
ANTLR=java -Xmx500M -cp $(ANTLR_JAR) org.antlr.v4.Tool
GEN_PARSER=java -cp $(ANTLR_JAR) -Xmx500M org.antlr.v4.Tool -visitor -no-listener -Werror -Xexact-output-dir spec/RCL.g4

.PHONY: download-dep
download-dep:
	cd script; wget http://www.antlr.org/download/antlr-$(ANTLR_VERSION)-complete.jar 

.PHONY: antlr
antlr:
	java -cp $(ANTLR_JAR) -Xmx500M org.antlr.v4.Tool

.PHONY: gen-parser
gen-parser:
	$(GEN_PARSER) -package org.reika.rcl.parser -o impl/rcl-j/src/main/java/org/reika/rcl/parser
	$(GEN_PARSER) -package rcl -o impl/rcl-go/rcl -Dlanguage=Go