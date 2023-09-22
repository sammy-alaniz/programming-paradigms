import pytest
import client
from antlr4 import InputStream


def invalid_test(input):
    with pytest.raises(ValueError) as exec_info:
        client.parse_cypher(input)
    assert str(exec_info.value) == "Parser caught error!"

def test_one():
    input = InputStream("blah")
    invalid_test(input)
