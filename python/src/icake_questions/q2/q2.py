# pylint: disable=missing-docstring
# pylint: disable-msg=C0103

def get_product_of_all_ints_except_at_index(nums):
    if len(nums) in [0, 1]:
        return nums

    products = [1] * len(nums)
    for i in range(len(nums)):
        for j, num in enumerate(nums):
            if i == j:
                continue
            else:
                products[i] = products[i] * num
    return products
