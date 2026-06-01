class MinStack:

    def __init__(self):
        self.stack = []
        minvalue = 0

    def push(self, val: int) -> None:
        if not self.stack:
            self.stack.append(0)
            self.minvalue = val
        else:
            self.stack.append(val - self.minvalue)
            if val < self.minvalue:
                self.minvalue = val
        
    def pop(self) -> None:
        if not self.stack:
            return
        
        top = self.stack.pop()
        if top < 0:
            self.minvalue = self.minvalue - top       

    def top(self) -> int:
        top = self.stack[-1]
        if top > 0:
            return top + self.minvalue
        else:
            return self.minvalue

    def getMin(self) -> int:
        return self.minvalue
        
