import pytest
from ..topics.algorithms import Algorithms

def test_algos0218():
    nums = [1,5]
    target = 5
    res = Algorithms.p704(nums=nums, target=target)
    assert res == 1

def test_algos0218_2():
    nums = [-1,0,3,5,9,12]
    target = 9
    res = Algorithms.p704(nums=nums, target=target)
    assert res == 4


