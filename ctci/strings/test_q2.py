
"""
Unit test for question q2
"""

import unittest
from q2 import Text

class PermutationTestCase(unittest.TestCase):
    """ this pylint is annoying """
    def test_unequal_strings(self): # pylint: disable=missing-docstring
        text1 = Text("hello")
        self.assertFalse(text1.is_permutation_of("hellomello"))
    
    def test_equal_strings(self): # pylint: disable=missing-docstring
        text1 = Text("hellow")
        self.assertTrue(text1.is_permutation_of("wohell"))

if __name__ == "__main__":
    unittest.main()
