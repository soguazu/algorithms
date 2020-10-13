def insertion_sort(arr):
    for i in range(1, len(arr), 1):
        currentValue = arr[i]
        index = i
        for j in range(i, -1, -1):
            if arr[j] > currentValue:   
                arr[j + 1] = arr[j]
                index-=1
        
        arr[index] = currentValue
    return arr
    
    
print(insertion_sort([4, 9, 3, 14, 11, 50, 34, 23, 12]))