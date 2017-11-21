""" Unit tests for q2 """
import unittest
import q2

class Q2TestCase(unittest.TestCase):
    """ Unit tests for q2 """
    def test_empty_list(self):
        """ Test empty input list """
        numbers = []
        products = q2.get_product_of_all_ints_except_at_index(numbers)
        self.assertEqual(numbers, products)
    
    def test_multi_values(self):
        """ Test inputs with multiple values """
        numbers = [1, 7, 3, 4]
        expected_products = [84, 12, 28, 21]
        actual_products = q2.get_product_of_all_ints_except_at_index(numbers)
        self.assertEqual(expected_products, actual_products)

if __name__ == '__main__':
    unittest.main()
    