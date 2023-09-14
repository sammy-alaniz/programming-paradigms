import pytest
import client
from antlr4 import InputStream


def test_one():
    input = InputStream("blah")
    with pytest.raises(ValueError) as exec_info:
        client.parse_cypher(input)
    assert str(exec_info.value) == "Parser caught error!"