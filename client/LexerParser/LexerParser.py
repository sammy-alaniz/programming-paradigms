#!/usr/bin/python3


from antlr4 import InputStream, CommonTokenStream
from antlr4.error.ErrorListener import ErrorListener

from generated.CypherLexer import CypherLexer
from generated.CypherParser import CypherParser

class DetectError(ErrorListener):
    def syntaxError(self, recognizer, offendingSymbol, line, column, msg, e):
        print('sammy was here')
        print('Line: ', line)
        print('Column: ', column)
        print('Message: ', msg)
        print('E:', e)
        raise ValueError("Parser caught error!")
        #return super().syntaxError(recognizer, offendingSymbol, line, column, msg, e) # do I really need this line?
    
def parse_cypher(input):
    lexer = CypherLexer(input)
    tokens = CommonTokenStream(lexer)
    parser = CypherParser(tokens)
    detect_error = DetectError()
    parser.removeErrorListeners()
    parser.addErrorListener(detect_error)
    tree = parser.oC_Cypher()


if __name__ == "__main__":
    print('LexerParser main hit!')
    query = "MATCH (a:bike) RETURN a"
    input = InputStream(query)
    try:
        parse_cypher(input)
    except Exception as e:
        print('dang')
        print(e)

    print("passed")