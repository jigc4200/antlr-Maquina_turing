// Generated from /home/jigc4200/Documents/antlr-Maquina_turing/TuringLexer.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue", "this-escape"})
public class TuringLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		EN_EL_ESTADO=1, SI_SE_LEE_UN=2, SE_ESCRIBE_UN=3, SE_MUEVE_HACIA_LA=4, 
		NO_SE_MUEVE=5, CONTINUA_EN_EL_ESTADO=6, TERMINA=7, BLANCO=8, CERO=9, UNO=10, 
		IZQUIERDA=11, DERECHA=12, COMA=13, ESTADO=14, WS=15;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"EN_EL_ESTADO", "SI_SE_LEE_UN", "SE_ESCRIBE_UN", "SE_MUEVE_HACIA_LA", 
			"NO_SE_MUEVE", "CONTINUA_EN_EL_ESTADO", "TERMINA", "BLANCO", "CERO", 
			"UNO", "IZQUIERDA", "DERECHA", "COMA", "ESTADO", "WS"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'en_el_estado_'", "'si_se_lee_un_'", "'se_escribe_un_'", "'se_mueve_hacia_la_'", 
			"'no_se_mueve'", "'continua_en_el_estado_'", "'termina'", "'blanco'", 
			"'cero'", "'uno'", "'izquierda'", "'derecha'", "','"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "EN_EL_ESTADO", "SI_SE_LEE_UN", "SE_ESCRIBE_UN", "SE_MUEVE_HACIA_LA", 
			"NO_SE_MUEVE", "CONTINUA_EN_EL_ESTADO", "TERMINA", "BLANCO", "CERO", 
			"UNO", "IZQUIERDA", "DERECHA", "COMA", "ESTADO", "WS"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
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


	public TuringLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "TuringLexer.g4"; }

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
		"\u0004\u0000\u000f\u00bd\u0006\uffff\uffff\u0002\u0000\u0007\u0000\u0002"+
		"\u0001\u0007\u0001\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002"+
		"\u0004\u0007\u0004\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002"+
		"\u0007\u0007\u0007\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002"+
		"\u000b\u0007\u000b\u0002\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e"+
		"\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000"+
		"\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000"+
		"\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\b\u0001\b\u0001"+
		"\b\u0001\b\u0001\b\u0001\t\u0001\t\u0001\t\u0001\t\u0001\n\u0001\n\u0001"+
		"\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\u000b"+
		"\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b"+
		"\u0001\u000b\u0001\f\u0001\f\u0001\r\u0001\r\u0001\r\u0005\r\u00b0\b\r"+
		"\n\r\f\r\u00b3\t\r\u0003\r\u00b5\b\r\u0001\u000e\u0004\u000e\u00b8\b\u000e"+
		"\u000b\u000e\f\u000e\u00b9\u0001\u000e\u0001\u000e\u0000\u0000\u000f\u0001"+
		"\u0001\u0003\u0002\u0005\u0003\u0007\u0004\t\u0005\u000b\u0006\r\u0007"+
		"\u000f\b\u0011\t\u0013\n\u0015\u000b\u0017\f\u0019\r\u001b\u000e\u001d"+
		"\u000f\u0001\u0000\u0004\u0001\u0000az\u0001\u000019\u0001\u000009\u0003"+
		"\u0000\t\n\r\r  \u00bf\u0000\u0001\u0001\u0000\u0000\u0000\u0000\u0003"+
		"\u0001\u0000\u0000\u0000\u0000\u0005\u0001\u0000\u0000\u0000\u0000\u0007"+
		"\u0001\u0000\u0000\u0000\u0000\t\u0001\u0000\u0000\u0000\u0000\u000b\u0001"+
		"\u0000\u0000\u0000\u0000\r\u0001\u0000\u0000\u0000\u0000\u000f\u0001\u0000"+
		"\u0000\u0000\u0000\u0011\u0001\u0000\u0000\u0000\u0000\u0013\u0001\u0000"+
		"\u0000\u0000\u0000\u0015\u0001\u0000\u0000\u0000\u0000\u0017\u0001\u0000"+
		"\u0000\u0000\u0000\u0019\u0001\u0000\u0000\u0000\u0000\u001b\u0001\u0000"+
		"\u0000\u0000\u0000\u001d\u0001\u0000\u0000\u0000\u0001\u001f\u0001\u0000"+
		"\u0000\u0000\u0003-\u0001\u0000\u0000\u0000\u0005;\u0001\u0000\u0000\u0000"+
		"\u0007J\u0001\u0000\u0000\u0000\t]\u0001\u0000\u0000\u0000\u000bi\u0001"+
		"\u0000\u0000\u0000\r\u0080\u0001\u0000\u0000\u0000\u000f\u0088\u0001\u0000"+
		"\u0000\u0000\u0011\u008f\u0001\u0000\u0000\u0000\u0013\u0094\u0001\u0000"+
		"\u0000\u0000\u0015\u0098\u0001\u0000\u0000\u0000\u0017\u00a2\u0001\u0000"+
		"\u0000\u0000\u0019\u00aa\u0001\u0000\u0000\u0000\u001b\u00b4\u0001\u0000"+
		"\u0000\u0000\u001d\u00b7\u0001\u0000\u0000\u0000\u001f \u0005e\u0000\u0000"+
		" !\u0005n\u0000\u0000!\"\u0005_\u0000\u0000\"#\u0005e\u0000\u0000#$\u0005"+
		"l\u0000\u0000$%\u0005_\u0000\u0000%&\u0005e\u0000\u0000&\'\u0005s\u0000"+
		"\u0000\'(\u0005t\u0000\u0000()\u0005a\u0000\u0000)*\u0005d\u0000\u0000"+
		"*+\u0005o\u0000\u0000+,\u0005_\u0000\u0000,\u0002\u0001\u0000\u0000\u0000"+
		"-.\u0005s\u0000\u0000./\u0005i\u0000\u0000/0\u0005_\u0000\u000001\u0005"+
		"s\u0000\u000012\u0005e\u0000\u000023\u0005_\u0000\u000034\u0005l\u0000"+
		"\u000045\u0005e\u0000\u000056\u0005e\u0000\u000067\u0005_\u0000\u0000"+
		"78\u0005u\u0000\u000089\u0005n\u0000\u00009:\u0005_\u0000\u0000:\u0004"+
		"\u0001\u0000\u0000\u0000;<\u0005s\u0000\u0000<=\u0005e\u0000\u0000=>\u0005"+
		"_\u0000\u0000>?\u0005e\u0000\u0000?@\u0005s\u0000\u0000@A\u0005c\u0000"+
		"\u0000AB\u0005r\u0000\u0000BC\u0005i\u0000\u0000CD\u0005b\u0000\u0000"+
		"DE\u0005e\u0000\u0000EF\u0005_\u0000\u0000FG\u0005u\u0000\u0000GH\u0005"+
		"n\u0000\u0000HI\u0005_\u0000\u0000I\u0006\u0001\u0000\u0000\u0000JK\u0005"+
		"s\u0000\u0000KL\u0005e\u0000\u0000LM\u0005_\u0000\u0000MN\u0005m\u0000"+
		"\u0000NO\u0005u\u0000\u0000OP\u0005e\u0000\u0000PQ\u0005v\u0000\u0000"+
		"QR\u0005e\u0000\u0000RS\u0005_\u0000\u0000ST\u0005h\u0000\u0000TU\u0005"+
		"a\u0000\u0000UV\u0005c\u0000\u0000VW\u0005i\u0000\u0000WX\u0005a\u0000"+
		"\u0000XY\u0005_\u0000\u0000YZ\u0005l\u0000\u0000Z[\u0005a\u0000\u0000"+
		"[\\\u0005_\u0000\u0000\\\b\u0001\u0000\u0000\u0000]^\u0005n\u0000\u0000"+
		"^_\u0005o\u0000\u0000_`\u0005_\u0000\u0000`a\u0005s\u0000\u0000ab\u0005"+
		"e\u0000\u0000bc\u0005_\u0000\u0000cd\u0005m\u0000\u0000de\u0005u\u0000"+
		"\u0000ef\u0005e\u0000\u0000fg\u0005v\u0000\u0000gh\u0005e\u0000\u0000"+
		"h\n\u0001\u0000\u0000\u0000ij\u0005c\u0000\u0000jk\u0005o\u0000\u0000"+
		"kl\u0005n\u0000\u0000lm\u0005t\u0000\u0000mn\u0005i\u0000\u0000no\u0005"+
		"n\u0000\u0000op\u0005u\u0000\u0000pq\u0005a\u0000\u0000qr\u0005_\u0000"+
		"\u0000rs\u0005e\u0000\u0000st\u0005n\u0000\u0000tu\u0005_\u0000\u0000"+
		"uv\u0005e\u0000\u0000vw\u0005l\u0000\u0000wx\u0005_\u0000\u0000xy\u0005"+
		"e\u0000\u0000yz\u0005s\u0000\u0000z{\u0005t\u0000\u0000{|\u0005a\u0000"+
		"\u0000|}\u0005d\u0000\u0000}~\u0005o\u0000\u0000~\u007f\u0005_\u0000\u0000"+
		"\u007f\f\u0001\u0000\u0000\u0000\u0080\u0081\u0005t\u0000\u0000\u0081"+
		"\u0082\u0005e\u0000\u0000\u0082\u0083\u0005r\u0000\u0000\u0083\u0084\u0005"+
		"m\u0000\u0000\u0084\u0085\u0005i\u0000\u0000\u0085\u0086\u0005n\u0000"+
		"\u0000\u0086\u0087\u0005a\u0000\u0000\u0087\u000e\u0001\u0000\u0000\u0000"+
		"\u0088\u0089\u0005b\u0000\u0000\u0089\u008a\u0005l\u0000\u0000\u008a\u008b"+
		"\u0005a\u0000\u0000\u008b\u008c\u0005n\u0000\u0000\u008c\u008d\u0005c"+
		"\u0000\u0000\u008d\u008e\u0005o\u0000\u0000\u008e\u0010\u0001\u0000\u0000"+
		"\u0000\u008f\u0090\u0005c\u0000\u0000\u0090\u0091\u0005e\u0000\u0000\u0091"+
		"\u0092\u0005r\u0000\u0000\u0092\u0093\u0005o\u0000\u0000\u0093\u0012\u0001"+
		"\u0000\u0000\u0000\u0094\u0095\u0005u\u0000\u0000\u0095\u0096\u0005n\u0000"+
		"\u0000\u0096\u0097\u0005o\u0000\u0000\u0097\u0014\u0001\u0000\u0000\u0000"+
		"\u0098\u0099\u0005i\u0000\u0000\u0099\u009a\u0005z\u0000\u0000\u009a\u009b"+
		"\u0005q\u0000\u0000\u009b\u009c\u0005u\u0000\u0000\u009c\u009d\u0005i"+
		"\u0000\u0000\u009d\u009e\u0005e\u0000\u0000\u009e\u009f\u0005r\u0000\u0000"+
		"\u009f\u00a0\u0005d\u0000\u0000\u00a0\u00a1\u0005a\u0000\u0000\u00a1\u0016"+
		"\u0001\u0000\u0000\u0000\u00a2\u00a3\u0005d\u0000\u0000\u00a3\u00a4\u0005"+
		"e\u0000\u0000\u00a4\u00a5\u0005r\u0000\u0000\u00a5\u00a6\u0005e\u0000"+
		"\u0000\u00a6\u00a7\u0005c\u0000\u0000\u00a7\u00a8\u0005h\u0000\u0000\u00a8"+
		"\u00a9\u0005a\u0000\u0000\u00a9\u0018\u0001\u0000\u0000\u0000\u00aa\u00ab"+
		"\u0005,\u0000\u0000\u00ab\u001a\u0001\u0000\u0000\u0000\u00ac\u00b5\u0007"+
		"\u0000\u0000\u0000\u00ad\u00b1\u0007\u0001\u0000\u0000\u00ae\u00b0\u0007"+
		"\u0002\u0000\u0000\u00af\u00ae\u0001\u0000\u0000\u0000\u00b0\u00b3\u0001"+
		"\u0000\u0000\u0000\u00b1\u00af\u0001\u0000\u0000\u0000\u00b1\u00b2\u0001"+
		"\u0000\u0000\u0000\u00b2\u00b5\u0001\u0000\u0000\u0000\u00b3\u00b1\u0001"+
		"\u0000\u0000\u0000\u00b4\u00ac\u0001\u0000\u0000\u0000\u00b4\u00ad\u0001"+
		"\u0000\u0000\u0000\u00b5\u001c\u0001\u0000\u0000\u0000\u00b6\u00b8\u0007"+
		"\u0003\u0000\u0000\u00b7\u00b6\u0001\u0000\u0000\u0000\u00b8\u00b9\u0001"+
		"\u0000\u0000\u0000\u00b9\u00b7\u0001\u0000\u0000\u0000\u00b9\u00ba\u0001"+
		"\u0000\u0000\u0000\u00ba\u00bb\u0001\u0000\u0000\u0000\u00bb\u00bc\u0006"+
		"\u000e\u0000\u0000\u00bc\u001e\u0001\u0000\u0000\u0000\u0004\u0000\u00b1"+
		"\u00b4\u00b9\u0001\u0006\u0000\u0000";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}