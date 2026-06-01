class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        for i in range(len(matrix)):
            row = matrix[i]
            start, end = 0, len(row) - 1
            x=row[start]
            y=row[end]
            if target >= x and target <= y:
                mid = start + (end-start)//2
                while start<=end and end>=0:
                    if target > row[mid]:
                        start = mid + 1
                    elif target < row[mid]:
                        end = mid - 1
                    elif target == row[mid]:
                        return True
                    else:
                        return False
                    mid = start + (end-start)//2
            else:
                continue
        return False

