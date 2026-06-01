class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        operators = ["+", "-", "*", "/"]
        numstack = []
        res = 0
        for i in tokens:
            if i not in operators:
                numstack.append(i)
            else:
                top1 = int(numstack.pop())
                top2 = int(numstack.pop())
                if i == "+":
                    res = top2 + top1
                elif i == "-":
                    res = top2 - top1
                elif i =="*":
                    res = top2 * top1
                else:
                    res = int(top2 / top1)
                numstack.append(res)
        return int(numstack[-1])