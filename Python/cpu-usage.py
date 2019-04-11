import psutil                                     #import tsutil
import time

def printit():
    while(1):
        print(psutil.cpu_percent())
        time.sleep(3)
printit()