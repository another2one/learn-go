# from datetime import datetime

# num = 20000

# start = datetime.now().timestamp()

# def addNum(i):
#     sum = 0
#     while i > 0:
#         sum+= i
#         i-= 1
#     return sum
    
# i = 1
# while i <= num:
#     addNum(i)
#     i+= 1

# end = datetime.now().timestamp()

# print("used %.3fs" % (end - start))

import threading
import asyncio

@asyncio.coroutine
def hello():
    print('Hello world! (%s)' % threading.currentThread())
    yield from asyncio.sleep(2)
    print('Hello again! (%s)' % threading.currentThread())

loop = asyncio.get_event_loop()
tasks = [hello(), hello()]
loop.run_until_complete(asyncio.wait(tasks))
loop.close()