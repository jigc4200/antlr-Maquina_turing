// Generated from /home/jigc4200/Documents/antlr-Maquina_turing/TuringParser.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class TuringParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		EN_EL_ESTADO=1, SI_SE_LEE_UN=2, SE_ESCRIBE_UN=3, SE_MUEVE_HACIA_LA=4, 
		NO_SE_MUEVE=5, CONTINUA_EN_EL_ESTADO=6, TERMINA=7, BLANCO=8, CERO=9, UNO=10, 
		IZQUIERDA=11, DERECHA=12, COMA=13, ESTADO=14, WS=15;
	public static final int
		RULE_programa = 0, RULE_instruccion = 1, RULE_estadoActual = 2, RULE_simboloLectura = 3, 
		RULE_simboloEscritura = 4, RULE_movimiento = 5, RULE_direccion = 6, RULE_siguienteEstado = 7;
	private static String[] makeRuleNames() {
		return new String[] {
			"programa", "instruccion", "estadoActual", "simboloLectura", "simboloEscritura", 
			"movimiento", "direccion", "siguienteEstado"
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

	@Override
	public String getGrammarFileName() { return "TuringParser.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public TuringParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProgramaContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(TuringParser.EOF, 0); }
		public List<InstruccionContext> instruccion() {
			return getRuleContexts(InstruccionContext.class);
		}
		public InstruccionContext instruccion(int i) {
			return getRuleContext(InstruccionContext.class,i);
		}
		public ProgramaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_programa; }
	}

	public final ProgramaContext programa() throws RecognitionException {
		ProgramaContext _localctx = new ProgramaContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_programa);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(17); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(16);
				instruccion();
				}
				}
				setState(19); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==EN_EL_ESTADO );
			setState(21);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class InstruccionContext extends ParserRuleContext {
		public EstadoActualContext estadoActual() {
			return getRuleContext(EstadoActualContext.class,0);
		}
		public List<TerminalNode> COMA() { return getTokens(TuringParser.COMA); }
		public TerminalNode COMA(int i) {
			return getToken(TuringParser.COMA, i);
		}
		public SimboloLecturaContext simboloLectura() {
			return getRuleContext(SimboloLecturaContext.class,0);
		}
		public SimboloEscrituraContext simboloEscritura() {
			return getRuleContext(SimboloEscrituraContext.class,0);
		}
		public MovimientoContext movimiento() {
			return getRuleContext(MovimientoContext.class,0);
		}
		public SiguienteEstadoContext siguienteEstado() {
			return getRuleContext(SiguienteEstadoContext.class,0);
		}
		public InstruccionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion; }
	}

	public final InstruccionContext instruccion() throws RecognitionException {
		InstruccionContext _localctx = new InstruccionContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_instruccion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(23);
			estadoActual();
			setState(24);
			match(COMA);
			setState(25);
			simboloLectura();
			setState(26);
			match(COMA);
			setState(27);
			simboloEscritura();
			setState(28);
			match(COMA);
			setState(29);
			movimiento();
			setState(30);
			match(COMA);
			setState(31);
			siguienteEstado();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class EstadoActualContext extends ParserRuleContext {
		public TerminalNode EN_EL_ESTADO() { return getToken(TuringParser.EN_EL_ESTADO, 0); }
		public TerminalNode ESTADO() { return getToken(TuringParser.ESTADO, 0); }
		public EstadoActualContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_estadoActual; }
	}

	public final EstadoActualContext estadoActual() throws RecognitionException {
		EstadoActualContext _localctx = new EstadoActualContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_estadoActual);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(33);
			match(EN_EL_ESTADO);
			setState(34);
			match(ESTADO);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SimboloLecturaContext extends ParserRuleContext {
		public TerminalNode SI_SE_LEE_UN() { return getToken(TuringParser.SI_SE_LEE_UN, 0); }
		public TerminalNode BLANCO() { return getToken(TuringParser.BLANCO, 0); }
		public TerminalNode CERO() { return getToken(TuringParser.CERO, 0); }
		public TerminalNode UNO() { return getToken(TuringParser.UNO, 0); }
		public SimboloLecturaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_simboloLectura; }
	}

	public final SimboloLecturaContext simboloLectura() throws RecognitionException {
		SimboloLecturaContext _localctx = new SimboloLecturaContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_simboloLectura);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(36);
			match(SI_SE_LEE_UN);
			setState(37);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 1792L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SimboloEscrituraContext extends ParserRuleContext {
		public TerminalNode SE_ESCRIBE_UN() { return getToken(TuringParser.SE_ESCRIBE_UN, 0); }
		public TerminalNode BLANCO() { return getToken(TuringParser.BLANCO, 0); }
		public TerminalNode CERO() { return getToken(TuringParser.CERO, 0); }
		public TerminalNode UNO() { return getToken(TuringParser.UNO, 0); }
		public SimboloEscrituraContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_simboloEscritura; }
	}

	public final SimboloEscrituraContext simboloEscritura() throws RecognitionException {
		SimboloEscrituraContext _localctx = new SimboloEscrituraContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_simboloEscritura);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(39);
			match(SE_ESCRIBE_UN);
			setState(40);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 1792L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class MovimientoContext extends ParserRuleContext {
		public TerminalNode SE_MUEVE_HACIA_LA() { return getToken(TuringParser.SE_MUEVE_HACIA_LA, 0); }
		public DireccionContext direccion() {
			return getRuleContext(DireccionContext.class,0);
		}
		public TerminalNode NO_SE_MUEVE() { return getToken(TuringParser.NO_SE_MUEVE, 0); }
		public MovimientoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_movimiento; }
	}

	public final MovimientoContext movimiento() throws RecognitionException {
		MovimientoContext _localctx = new MovimientoContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_movimiento);
		try {
			setState(45);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case SE_MUEVE_HACIA_LA:
				enterOuterAlt(_localctx, 1);
				{
				setState(42);
				match(SE_MUEVE_HACIA_LA);
				setState(43);
				direccion();
				}
				break;
			case NO_SE_MUEVE:
				enterOuterAlt(_localctx, 2);
				{
				setState(44);
				match(NO_SE_MUEVE);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class DireccionContext extends ParserRuleContext {
		public TerminalNode IZQUIERDA() { return getToken(TuringParser.IZQUIERDA, 0); }
		public TerminalNode DERECHA() { return getToken(TuringParser.DERECHA, 0); }
		public DireccionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_direccion; }
	}

	public final DireccionContext direccion() throws RecognitionException {
		DireccionContext _localctx = new DireccionContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_direccion);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(47);
			_la = _input.LA(1);
			if ( !(_la==IZQUIERDA || _la==DERECHA) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SiguienteEstadoContext extends ParserRuleContext {
		public TerminalNode CONTINUA_EN_EL_ESTADO() { return getToken(TuringParser.CONTINUA_EN_EL_ESTADO, 0); }
		public TerminalNode ESTADO() { return getToken(TuringParser.ESTADO, 0); }
		public TerminalNode TERMINA() { return getToken(TuringParser.TERMINA, 0); }
		public SiguienteEstadoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_siguienteEstado; }
	}

	public final SiguienteEstadoContext siguienteEstado() throws RecognitionException {
		SiguienteEstadoContext _localctx = new SiguienteEstadoContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_siguienteEstado);
		try {
			setState(52);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case CONTINUA_EN_EL_ESTADO:
				enterOuterAlt(_localctx, 1);
				{
				setState(49);
				match(CONTINUA_EN_EL_ESTADO);
				setState(50);
				match(ESTADO);
				}
				break;
			case TERMINA:
				enterOuterAlt(_localctx, 2);
				{
				setState(51);
				match(TERMINA);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\u0004\u0001\u000f7\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0001"+
		"\u0000\u0004\u0000\u0012\b\u0000\u000b\u0000\f\u0000\u0013\u0001\u0000"+
		"\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0005\u0001\u0005\u0001\u0005\u0003\u0005"+
		".\b\u0005\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0003\u00075\b\u0007\u0001\u0007\u0000\u0000\b\u0000\u0002\u0004\u0006"+
		"\b\n\f\u000e\u0000\u0002\u0001\u0000\b\n\u0001\u0000\u000b\f1\u0000\u0011"+
		"\u0001\u0000\u0000\u0000\u0002\u0017\u0001\u0000\u0000\u0000\u0004!\u0001"+
		"\u0000\u0000\u0000\u0006$\u0001\u0000\u0000\u0000\b\'\u0001\u0000\u0000"+
		"\u0000\n-\u0001\u0000\u0000\u0000\f/\u0001\u0000\u0000\u0000\u000e4\u0001"+
		"\u0000\u0000\u0000\u0010\u0012\u0003\u0002\u0001\u0000\u0011\u0010\u0001"+
		"\u0000\u0000\u0000\u0012\u0013\u0001\u0000\u0000\u0000\u0013\u0011\u0001"+
		"\u0000\u0000\u0000\u0013\u0014\u0001\u0000\u0000\u0000\u0014\u0015\u0001"+
		"\u0000\u0000\u0000\u0015\u0016\u0005\u0000\u0000\u0001\u0016\u0001\u0001"+
		"\u0000\u0000\u0000\u0017\u0018\u0003\u0004\u0002\u0000\u0018\u0019\u0005"+
		"\r\u0000\u0000\u0019\u001a\u0003\u0006\u0003\u0000\u001a\u001b\u0005\r"+
		"\u0000\u0000\u001b\u001c\u0003\b\u0004\u0000\u001c\u001d\u0005\r\u0000"+
		"\u0000\u001d\u001e\u0003\n\u0005\u0000\u001e\u001f\u0005\r\u0000\u0000"+
		"\u001f \u0003\u000e\u0007\u0000 \u0003\u0001\u0000\u0000\u0000!\"\u0005"+
		"\u0001\u0000\u0000\"#\u0005\u000e\u0000\u0000#\u0005\u0001\u0000\u0000"+
		"\u0000$%\u0005\u0002\u0000\u0000%&\u0007\u0000\u0000\u0000&\u0007\u0001"+
		"\u0000\u0000\u0000\'(\u0005\u0003\u0000\u0000()\u0007\u0000\u0000\u0000"+
		")\t\u0001\u0000\u0000\u0000*+\u0005\u0004\u0000\u0000+.\u0003\f\u0006"+
		"\u0000,.\u0005\u0005\u0000\u0000-*\u0001\u0000\u0000\u0000-,\u0001\u0000"+
		"\u0000\u0000.\u000b\u0001\u0000\u0000\u0000/0\u0007\u0001\u0000\u0000"+
		"0\r\u0001\u0000\u0000\u000012\u0005\u0006\u0000\u000025\u0005\u000e\u0000"+
		"\u000035\u0005\u0007\u0000\u000041\u0001\u0000\u0000\u000043\u0001\u0000"+
		"\u0000\u00005\u000f\u0001\u0000\u0000\u0000\u0003\u0013-4";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}