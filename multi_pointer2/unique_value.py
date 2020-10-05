def countUniqueValues(arr):
    return len(set(arr))      
            

def countUniqueValues(arr):
    if len(arr) < 1:
        return 0
    start = 0
    for i in range(1,len(arr)):
        if arr[start] != arr[i]:
            start += 1
            arr[start] = arr[i]
    
    return start + 1

            
print(countUniqueValues([1, 1, 1, 1, 1, 2]))
print(countUniqueValues([1, 2, 3, 4, 4, 4, 7, 7, 12, 12, 13]))
print(countUniqueValues([-2, -1, -1, 0, 1]))
print(countUniqueValues([]))         