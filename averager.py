

def averageResults():
    total_time = 0
    numLines = 0.0
    for i in range(1, 2):
        f = open('raw' + str(i) + '.txt', 'r')
        for line in f:
            time = line.split()[-1]
            if 'ms' in time:
                try:
                    total_time += float(time.replace('ms', ''))
                except:
                    continue
            elif 's' in time:
                try:
                    total_time += float(time.replace('s', '')) * 1000.0
                except:
                    continue
            numLines += 1
        f.close()
    print total_time / numLines
    return total_time

if __name__ == "__main__":
    averageResults()