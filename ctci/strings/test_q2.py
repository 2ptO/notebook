
"""
Unit test for question q2
"""

import unittest
from q2 import Text

# pylint: disable=missing-docstring

class PermutationTestCase(unittest.TestCase):
    """ this pylint is annoying """
    def test_equal_strings(self):
        """ doc string for test_equal_strings """
        text1 = Text("hellow")
        self.assertTrue(text1.is_permutation_of("wohell"))

    def test_unequal_strings(self):
        text1 = Text("hello")
        self.assertFalse(text1.is_permutation_of("hellomello"))

if __name__ == "__main__":
    unittest.main()
