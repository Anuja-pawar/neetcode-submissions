class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        stack = []
        for i in range(len(temperatures)):
            counter = 1
            for j in range(i+1, len(temperatures)):
                if j == len(temperatures) - 1 and temperatures[j] <= temperatures[i]:
                    counter = 0
                    break
                if temperatures[j] > temperatures[i]:
                    break
                counter += 1
            if i == len(temperatures) - 1:
                stack.append(0)
            else:
                stack.append(counter)
        return stack
                    


