import pytest
import client

def test_one():
    input = client.FileStream('/home/texas/dev/programming-paradigms/lexer-parser/test.txt')
    with pytest.raises(ValueError) as exec_info:
        client.parse_cypher(input)
    assert str(exec_info.value) == "Parser caught error!"