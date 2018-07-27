import csv

def main():
    csv_reader = csv.reader(open("vision_objectr.csv"))
    obj = {}
    keys = []
    for row in csv_reader:
        # print(len(row))
        for k in range(0,len(row)):
            if k%2 == 1 and row[k-1] != "":
                keys.append(int(row[k-1]))
                obj[int(row[k-1])] = row[k]
    keys.sort()
    print("var VisionObjectrNames = map[int]string{")
    for v in keys:
        print("    %s: \"%s\"," % (v,obj[v]))
    print("}")
if __name__ == '__main__':
    main()
