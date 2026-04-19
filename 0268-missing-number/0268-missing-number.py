class Solution:
    def missingNumber(self, nums: List[int]) -> int:
        help_list = [0 for i in range(len(nums)+1)]
        for idx, num in enumerate(nums):
            help_list[num] = 1
        for idx, num in enumerate(help_list):
            if num == 0:
                return idx
        raise Exception("Invalid input")