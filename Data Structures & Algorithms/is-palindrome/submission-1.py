class Solution:
    def isPalindrome(self, s: str) -> bool:
        newStr = ""
        for i in s:
            if i.isalnum():
                newStr += i
        newStr = newStr.lower()
        l = 0
        r = len(newStr) - 1
        while l <= r:   
            if newStr[l] != newStr[r]:
                return False
            l += 1
            r -= 1
        return True


