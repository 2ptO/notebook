#! /usr/local/bin/python
# pylint: disable=invalid-name

def main():
    """
    algo_crush logic testing
    """
    inputFile = open("input.txt", "r")
    #inputData = inputFile.readlines()

    N, M = [int(value) for value in inputFile.readline().split()]
 
    data = [0] * (N+1)

    for _ in range(M):
        a, b, adder = [int(value) for value in inputFile.readline().split()]
        data[a-1] += adder
        if b <= len(data):
            data[b] -= adder

    maxValue = x = 0
    for i in data:
        x = x + i
        if maxValue < (x):
            maxValue = x

    print maxValue
    inputFile.close()

def mainBackup():
    """
        Backup function to main()
    """
    inputFile = open("input.txt", "r")

    n, inputs = [int(n) for n in inputFile.readline().split(" ")]
    data = [0]*(n+1)
    for _ in range(inputs):
        x, y, incr = [int(n) for n in inputFile.readline().split(" ")]
        data[x-1] += incr
        if y <= len(data):
            data[y] -= incr

    maxValue = x = 0
    for i in data:
        x = x + i
        if maxValue < x:
            maxValue = x

    print maxValue
    inputFile.close()

if __name__ == "__main__":
    mainBackup()
    main()

