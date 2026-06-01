class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        res = defaultdict(list)
        for word in strs:
            char_list = [0 for x in range(26)]
            for ch in word:
                ind = ord(ch) - ord('a')
                char_list[ind] += 1
            
            res[tuple(char_list)].append(word)
        return list(res.values())

        