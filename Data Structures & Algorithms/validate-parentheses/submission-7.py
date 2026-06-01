class Solution:
    def isValid(self, s: str) -> bool:
        if len(s) == 1:
            return False
        stack = []
        openBrackets = ['[','{','(']
        closeBrackets = [']','}',')']
        for i in s:
            if i in openBrackets:
                stack.append(i)
            else:
                if i == '}' and len(stack) > 0:
                    last = stack.pop()
                    if last == '{':
                        continue
                    else:
                        return False
                elif i == ']' and len(stack) > 0:
                    last = stack.pop()
                    if last == '[':
                        continue
                    else:
                        return False
                elif i == ')' and len(stack) > 0:
                    last = stack.pop()
                    if last == '(':
                        continue
                    else:
                        return False
                else:
                    return False
        if len(stack) > 0:
            return False
        else:
            return True




        