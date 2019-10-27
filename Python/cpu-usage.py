import psutil
import time

def printit():
    while(1):
        print('{:06.2f}'.format(psutil.cpu_percent()), end='\r')
        time.sleep(1)
printit()
