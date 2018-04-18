# Reika

Survey of old Reika implementation https://github.com/at15/reika

- the new implementation is in `me.at15.reika.compiler`
- `Position` has start and end for both line and column, and is defined in `me.at15.reika.compiler.reporter`
  - might learned this from scala, reporter has `info(Position pos, String msg)`
  - [ ] TODO: there seems to be no way to get exact position of token across two lines
- `phases/ANTLR.java` generate parse tree using ANTLR
- `parser/SyntaxErrorListener` implements `ANTLRErrorListener` and `ErrorCollector`
  - when it goet `syntaxError` it would store in a list and log it (if log2Stdout is true)
- `ast/UntypedVisitor.java` convert parse tree to AST
  - it is used in `phases/AST.java`
- `ast/AStErrorListener.java` implements `reporter.ErrorCollector`
- RCL implementation is in `me.at15.reika.rcl` but it only has grammar file and generated parser