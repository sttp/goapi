@echo off
echo Compiling FilterExpressionSyntax grammar for Go...
java -jar antlr-4.9.2-complete.jar -Dlanguage=Go -o parser FilterExpressionSyntax.g4
echo Finished.