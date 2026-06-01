# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        curr1 = l1
        counter1 = 0
        num1 = 0

        while curr1:
            n1 = curr1.val
            num1 = num1 + n1 * (10**counter1)   
            curr1 = curr1.next
            counter1 += 1

        curr2 = l2
        counter2 = 0
        num2 = 0

        while curr2:
            n2 = curr2.val
            num2 = num2 + n2 * (10 ** counter2)
            curr2 = curr2.next
            counter2 += 1

        sum = num1 + num2
        print(sum)

        if sum == 0:
            return ListNode(0)
        
        dummy = ListNode(0)
        curr = dummy
        while sum > 0:
            remainder = sum % 10
            print(remainder)
            curr.next = ListNode(remainder)
            curr = curr.next
            sum = sum // 10
        
        return dummy.next

        
