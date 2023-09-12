from antlr4 import FileStream, CommonTokenStream

from generated.CypherLexer import CypherLexer
from generated.CypherParser import CypherParser

input = FileStream('/home/texas/dev/programming-paradigms/lexer-parser/test.txt')
lexer = CypherLexer(input)
tokens = CommonTokenStream(lexer)
parser = CypherParser(tokens)
tree = parser.oC_Cypher()
