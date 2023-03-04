from  typing import List
import logging

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

class Algorithms:
    @staticmethod
    def p704(nums: List[int], target: int) -> int:
        low, high = 0, len(nums) - 1
        if nums[low] == target:
            return low
        if nums[high] == target:
            return high
        while low < high:
            mid = (low + high) // 2
            logger.info(low, high, mid)
            if target == nums[mid]:
                return mid

            elif target > nums[mid]:
                low = mid + 1

            else:
                high = mid

        return -1
