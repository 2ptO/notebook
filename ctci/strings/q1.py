#!/usr/local/bin/python3
"""
Implement an algorithm to determine if a string has all unique characters.
What if you can not use additional data structures?
Note: using additional buffer(map) to start with
Assumptions:
character set is limited to asci
Naming convention from Google Python:
module_name, package_name, ClassName, method_name, ExceptionName, function_name,
GLOBAL_CONSTANT_NAME, global_var_name, instance_var_name, function_parameter_name, local_var_name.
"""

class UniqueCharFinder(): # pylint: disable=too-few-public-methods
    """ Tests whether a given character is unique or not through different methods"""
    def __init__(self, text):
        self.text = text

    def run(self):
        """Runs all the testing methods"""
        print("Find whether {text} contains unique chars or not".format(text=self.text))
        print("using dict ", self.is_unique_using_dict())
        print("using bitmap ", self.is_unique_using_bitmap())

    def is_unique_using_dict(self):
        """ Find whether a given string contains unique characters or not """
        charmap = {}
        for cha in self.text:
            if cha in charmap:
                return False
            else:
                charmap[cha] = True
        return True

    def is_unique_using_bitmap(self):
        """ Find whether a given string contains unique character or not using bitmap.
            ASSUMPTION: character set is only lower case letters.. otherwise things will go crazy"""
        char_bitmap = 0x0
        for ch in self.text: # pylint: disable=invalid-name
            bitpos = (ord(ch) - ord('a'))
            assert bitpos >= 0 and bitpos < 25
            if (char_bitmap & 1<<bitpos) > 0:
                return False
            else:
                char_bitmap = char_bitmap | (1<<bitpos)
        return True

def main():
    """ Main function.."""
    user_input = input("Enter String:")
    ucf = UniqueCharFinder(user_input)
    ucf.run()

if __name__ == "__main__":
    main()
    