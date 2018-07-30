# -*- coding: utf-8 -*-

import csv
import sys
import string

def main(name):
    '''
    生成的map可能与go代码中名字不同，酌情修改
    '''
    csv_reader = csv.reader(open(name))
    obj = {}
    keys = []
    for row in csv_reader:
        # print(len(row))
        for k in range(0,len(row)):
            if k%2 == 1 and row[k-1] != "":
                keys.append(int(row[k-1]))
                obj[int(row[k-1])] = row[k]
    keys.sort()

    structName = getStructName(name)

    print("var %sNames = map[int]string{" % structName)
    for v in keys:
        print("    %s: \"%s\"," % (v,obj[v]))
    print("}")

def getStructName(name):
    name1 = name.split(".")[0]
    name2 = name1.split("_")
    structName = ""
    for v in name2:
        structName += string.capwords(v)
    return structName

if __name__ == '__main__':
    if len(sys.argv) < 2:
        print("请输入csv文件名")
        sys.exit()
    main(sys.argv[1])

