#! /usr/local/bin/python
"""
Given two strings, write a method to find whether one string
is a permutation of the other.
"""
class Text(object):
    """ String object """
    def __init__(self, val):
        self.val = val

    def is_permutation_of(self, other_text, use_sorting = False):
        """ uses sorting method """
        if len(self.val) != len(other_text):
            return False

        if use_sorting:
            return "".join(sorted(self.val)) == "".join(sorted(other_text))
        else:
            return self.is_permutation_of_nosort(other_text)

    def is_permutation_of_nosort(self, other_text):
        """ finds permutation using frequency table """
        letters = {}
        for letter in self.val:
            if letter in letters:
                letters[letter] = letters[letter] + 1
            else:
                letters[letter] = 1

        for letter in other_text:
            if letter in letters:
                letters[letter] = letters[letter] - 1
                if letters[letter] < 0:
                    return False
            else:
                return False
        return True
    
def main():
    """ Main """
    text1 = Text("hellow")
    print (text1.is_permutation_of("wohello"))
    print (text1.is_permutation_of("wohell"))

if __name__ == "__main__":
    main()
