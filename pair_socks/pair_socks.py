def socksMerchant(arr):
    temp = {}
    counter = 0
    
    for i in range(len(arr)):
        temp[arr[i]] = temp.get(arr[i], 0) + 1
        
        
    for i in temp:
        if temp[i] % 2 != 0 and temp[i] > 1:
            counter += (temp[i] - (temp[i] % 2)) / 2 
            
        if temp[i] % 2 == 0:
            counter += temp[i] / 2
    
    return counter
    
print(socksMerchant([10, 20, 20, 10, 10, 10, 30, 50, 1020]))
