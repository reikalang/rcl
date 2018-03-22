// Generated from spec/RCL.g4 by ANTLR 4.7.1
package org.reika.rcl.parser;
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class RCLLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.7.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, BOOL=7, INT=8, DOUBLE=9, 
		STRING=10, ID=11, WS=12, BLOCK_COMMENT=13, SINGLE_COMMENT=14;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	public static final String[] ruleNames = {
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "BOOL", "INT", "DOUBLE", 
		"STRING", "ID", "DIGIT", "ESC", "ID_LETTER", "WS", "BLOCK_COMMENT", "SINGLE_COMMENT"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "'='", "'{'", "','", "'}'", "'['", "']'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, null, null, null, null, null, null, "BOOL", "INT", "DOUBLE", "STRING", 
		"ID", "WS", "BLOCK_COMMENT", "SINGLE_COMMENT"
	};
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public RCLLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "RCL.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\20\u0089\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3\6\3\6\3\7\3\7\3\b\3\b\3\b\3\b"+
		"\3\b\3\b\3\b\3\b\3\b\5\b;\n\b\3\t\3\t\3\t\7\t@\n\t\f\t\16\tC\13\t\5\t"+
		"E\n\t\3\n\6\nH\n\n\r\n\16\nI\3\n\3\n\6\nN\n\n\r\n\16\nO\3\13\3\13\3\13"+
		"\7\13U\n\13\f\13\16\13X\13\13\3\13\3\13\3\f\3\f\7\f^\n\f\f\f\16\fa\13"+
		"\f\3\r\3\r\3\16\3\16\3\16\3\17\5\17i\n\17\3\20\3\20\3\20\3\20\3\21\3\21"+
		"\3\21\3\21\7\21s\n\21\f\21\16\21v\13\21\3\21\3\21\3\21\3\21\3\21\3\22"+
		"\3\22\3\22\3\22\7\22\u0081\n\22\f\22\16\22\u0084\13\22\3\22\3\22\3\22"+
		"\3\22\4t\u0082\2\23\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r"+
		"\31\2\33\2\35\2\37\16!\17#\20\3\2\t\3\2\63;\4\2$$^^\3\2c|\3\2\62;\n\2"+
		"$$\61\61^^ddhhppttvv\5\2C\\aac|\5\2\13\f\17\17\"\"\2\u008f\2\3\3\2\2\2"+
		"\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2"+
		"\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\37\3\2\2\2"+
		"\2!\3\2\2\2\2#\3\2\2\2\3%\3\2\2\2\5\'\3\2\2\2\7)\3\2\2\2\t+\3\2\2\2\13"+
		"-\3\2\2\2\r/\3\2\2\2\17:\3\2\2\2\21D\3\2\2\2\23G\3\2\2\2\25Q\3\2\2\2\27"+
		"[\3\2\2\2\31b\3\2\2\2\33d\3\2\2\2\35h\3\2\2\2\37j\3\2\2\2!n\3\2\2\2#|"+
		"\3\2\2\2%&\7?\2\2&\4\3\2\2\2\'(\7}\2\2(\6\3\2\2\2)*\7.\2\2*\b\3\2\2\2"+
		"+,\7\177\2\2,\n\3\2\2\2-.\7]\2\2.\f\3\2\2\2/\60\7_\2\2\60\16\3\2\2\2\61"+
		"\62\7v\2\2\62\63\7t\2\2\63\64\7w\2\2\64;\7g\2\2\65\66\7h\2\2\66\67\7c"+
		"\2\2\678\7n\2\289\7u\2\29;\7g\2\2:\61\3\2\2\2:\65\3\2\2\2;\20\3\2\2\2"+
		"<E\7\62\2\2=A\t\2\2\2>@\5\31\r\2?>\3\2\2\2@C\3\2\2\2A?\3\2\2\2AB\3\2\2"+
		"\2BE\3\2\2\2CA\3\2\2\2D<\3\2\2\2D=\3\2\2\2E\22\3\2\2\2FH\5\31\r\2GF\3"+
		"\2\2\2HI\3\2\2\2IG\3\2\2\2IJ\3\2\2\2JK\3\2\2\2KM\7\60\2\2LN\5\31\r\2M"+
		"L\3\2\2\2NO\3\2\2\2OM\3\2\2\2OP\3\2\2\2P\24\3\2\2\2QV\7$\2\2RU\5\33\16"+
		"\2SU\n\3\2\2TR\3\2\2\2TS\3\2\2\2UX\3\2\2\2VT\3\2\2\2VW\3\2\2\2WY\3\2\2"+
		"\2XV\3\2\2\2YZ\7$\2\2Z\26\3\2\2\2[_\t\4\2\2\\^\5\35\17\2]\\\3\2\2\2^a"+
		"\3\2\2\2_]\3\2\2\2_`\3\2\2\2`\30\3\2\2\2a_\3\2\2\2bc\t\5\2\2c\32\3\2\2"+
		"\2de\7^\2\2ef\t\6\2\2f\34\3\2\2\2gi\t\7\2\2hg\3\2\2\2i\36\3\2\2\2jk\t"+
		"\b\2\2kl\3\2\2\2lm\b\20\2\2m \3\2\2\2no\7\61\2\2op\7,\2\2pt\3\2\2\2qs"+
		"\13\2\2\2rq\3\2\2\2sv\3\2\2\2tu\3\2\2\2tr\3\2\2\2uw\3\2\2\2vt\3\2\2\2"+
		"wx\7,\2\2xy\7\61\2\2yz\3\2\2\2z{\b\21\2\2{\"\3\2\2\2|}\7\61\2\2}~\7\61"+
		"\2\2~\u0082\3\2\2\2\177\u0081\13\2\2\2\u0080\177\3\2\2\2\u0081\u0084\3"+
		"\2\2\2\u0082\u0083\3\2\2\2\u0082\u0080\3\2\2\2\u0083\u0085\3\2\2\2\u0084"+
		"\u0082\3\2\2\2\u0085\u0086\7\f\2\2\u0086\u0087\3\2\2\2\u0087\u0088\b\22"+
		"\2\2\u0088$\3\2\2\2\16\2:ADIOTV_ht\u0082\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}