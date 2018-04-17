

def getMedian():
    allTimes = []
    numLines = 0
    for i in range(0, 10):
        f = open('raw' + str(i) + '.txt', 'r')
        for line in f:
            time = line.split()[-1]
            if 'ms' in time:
                try:
                    allTimes.append(float(time.replace('ms', '')))
                except:
                    continue
            elif 's' in time:
                try:
                    allTimes.append(float(time.replace('s', '')) * 1000.0)
                except:
                    continue
            numLines += 1
        f.close()
    middleIndex = len(allTimes) / 2
    median = allTimes[middleIndex]
    print "Number of lines: ", numLines
    print "Number of elements in list: ", len(allTimes)
    print "Middle index is: ", middleIndex
    print "The median is: ", median, "ms"
    return median

if __name__ == "__main__":
    getMedian()
