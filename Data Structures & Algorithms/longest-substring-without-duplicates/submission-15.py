class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        l = r = 0
        longest = 0
        seen = dict()
        while r < len(s):
            if s[r] in seen:
                l = max(seen[s[r]]+1, l)
            seen[s[r]] = r
            longest = max(longest, r-l+1)
            r+=1
        return longest

        