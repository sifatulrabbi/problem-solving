from typing import List


class Solution:
    def carFleet(self, target: int, position: List[int], speed: List[int]) -> int:
        new_position: List[int] = position
        new_speed: List[int] = speed
        total_cars = len(position)
        finished = 0
        fleets = 0

        while finished < total_cars:
            updated_fleets = False
            for i in range(total_cars):
                new_position[i] += new_speed[i]
                if new_position[i] >= target:
                    finished += 1
                    if not updated_fleets:
                        updated_fleets = True
                        fleets += 1

        return fleets


if __name__ == "__main__":
    tests = [
        [10, [1, 4], [3, 2]],
        [10, [4, 1, 0, 7], [2, 2, 1, 1]],
    ]
    for [target, position, speed] in tests:
        s = Solution()
        print(f"test: target={target}", f"positions={position}", f"speeds={speed}")
        print("fleet count:", s.carFleet(target, position, speed))
