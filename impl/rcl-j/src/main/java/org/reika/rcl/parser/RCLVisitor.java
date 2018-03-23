// Generated from spec/RCL.g4 by ANTLR 4.7.1
package org.reika.rcl.parser;
import org.antlr.v4.runtime.tree.ParseTreeVisitor;

/**
 * This interface defines a complete generic visitor for a parse tree produced
 * by {@link RCLParser}.
 *
 * @param <T> The return type of the visit operation. Use {@link Void} for
 * operations with no return type.
 */
public interface RCLVisitor<T> extends ParseTreeVisitor<T> {
	/**
	 * Visit a parse tree produced by {@link RCLParser#rcl}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitRcl(RCLParser.RclContext ctx);
	/**
	 * Visit a parse tree produced by {@link RCLParser#pair}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPair(RCLParser.PairContext ctx);
	/**
	 * Visit a parse tree produced by {@link RCLParser#obj}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitObj(RCLParser.ObjContext ctx);
	/**
	 * Visit a parse tree produced by {@link RCLParser#array}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitArray(RCLParser.ArrayContext ctx);
	/**
	 * Visit a parse tree produced by the {@code ValBool}
	 * labeled alternative in {@link RCLParser#value}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitValBool(RCLParser.ValBoolContext ctx);
	/**
	 * Visit a parse tree produced by the {@code ValInt}
	 * labeled alternative in {@link RCLParser#value}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitValInt(RCLParser.ValIntContext ctx);
	/**
	 * Visit a parse tree produced by the {@code ValDouble}
	 * labeled alternative in {@link RCLParser#value}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitValDouble(RCLParser.ValDoubleContext ctx);
	/**
	 * Visit a parse tree produced by the {@code ValString}
	 * labeled alternative in {@link RCLParser#value}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitValString(RCLParser.ValStringContext ctx);
	/**
	 * Visit a parse tree produced by the {@code ValObject}
	 * labeled alternative in {@link RCLParser#value}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitValObject(RCLParser.ValObjectContext ctx);
	/**
	 * Visit a parse tree produced by the {@code ValArray}
	 * labeled alternative in {@link RCLParser#value}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitValArray(RCLParser.ValArrayContext ctx);
}