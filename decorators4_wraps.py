'''
In all 3 decorator exercises we lose the help information for the decorated function.
Rectify this issue
'''

from functools import wraps

def dec(func):
    @wraps(func)
    def ifunc(x, y):
        return func(x, y)
    return ifunc

@dec
def foo(p1, p2):
    ' help for foo'
    return p1 + p2

assert foo(10, 20) == 30

print help(foo)
print foo.__name__
        
print'all tests pass'
