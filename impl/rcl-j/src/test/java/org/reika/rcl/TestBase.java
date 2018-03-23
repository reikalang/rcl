package org.reika.rcl;

import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.CommonTokenStream;
import org.antlr.v4.runtime.tree.ParseTree;
import org.reika.rcl.parser.RCLLexer;
import org.reika.rcl.parser.RCLParser;

import java.io.FileNotFoundException;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.nio.charset.StandardCharsets;

public interface TestBase {
    default String readResourceText(String path) throws IOException {
        ClassLoader classLoader = TestBase.class.getClassLoader();
        InputStream is = classLoader.getResourceAsStream(path);
        return com.google.common.io.CharStreams.toString(new InputStreamReader(is, StandardCharsets.UTF_8));
    }

    default InputStream readResourceStream(String path) throws IOException {
        ClassLoader classLoader = TestBase.class.getClassLoader();
        InputStream is = classLoader.getResourceAsStream(path);
        if (is == null) {
            throw new FileNotFoundException(path);
        }
        return is;
    }

    default ParseTree getParseTreeFromResource(String path) throws IOException {
        CharStream charStream = org.antlr.v4.runtime.CharStreams.fromStream(readResourceStream(path));
        RCLLexer lexer = new RCLLexer(charStream);
        CommonTokenStream tokenStream = new CommonTokenStream(lexer);
        RCLParser parser = new RCLParser(tokenStream);
        return parser.rcl();
    }
}
