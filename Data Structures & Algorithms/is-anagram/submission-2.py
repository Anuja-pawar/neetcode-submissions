class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        s_count = {}
        t_count = {}
        if len(s) != len(t):
            return False
        for ch in s:
            s_count[ch] = s_count.get(ch, 0) + 1 
        
        for ch in t:
            t_count[ch] = t_count.get(ch, 0) + 1 
  
        for ch in s_count:
            if s_count.get(ch, 0) != t_count.get(ch, 0):
                return False

        return True

        