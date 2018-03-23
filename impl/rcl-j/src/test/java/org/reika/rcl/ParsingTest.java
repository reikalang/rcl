package org.reika.rcl;

import org.antlr.v4.runtime.tree.ParseTree;
import org.junit.jupiter.api.Test;

import java.io.IOException;

import static org.junit.jupiter.api.Assertions.assertNotNull;

public class ParsingTest implements TestBase {
    @Test
    void parseDB() throws IOException {
        ParseTree tree = getParseTreeFromResource("db-xephonk.rcl");
        assertNotNull(tree);
    }
}
