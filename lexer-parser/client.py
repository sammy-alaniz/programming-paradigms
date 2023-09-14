from antlr4 import FileStream, CommonTokenStream
from antlr4.error.ErrorListener import ErrorListener

from generated.CypherLexer import CypherLexer
from generated.CypherParser import CypherParser

input = FileStream('/home/texas/dev/programming-paradigms/lexer-parser/test.txt')
lexer = CypherLexer(input)
tokens = CommonTokenStream(lexer)
parser = CypherParser(tokens)



class DetectError(ErrorListener):
    def syntaxError(self, recognizer, offendingSymbol, line, column, msg, e):
        print('sammy was here')
        print('Line: ', line)
        print('Column: ', column)
        print('Message: ', msg)
        print('E:', e)
        return super().syntaxError(recognizer, offendingSymbol, line, column, msg, e)

detect_error = DetectError()
parser.removeErrorListeners()
parser.addErrorListener(detect_error)


tree = parser.oC_Cypher()

