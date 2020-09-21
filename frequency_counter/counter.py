def same(a,b):
    frequencyObj1 = {}
    frequencyObj2 = {}
    for value in a:
        frequencyObj1[value] = frequencyObj1.get(value, 0) + 1
    
    
    for value in b:
        frequencyObj2[value] = frequencyObj2.get(value, 0) + 1
        
    for value in frequencyObj1:
        if value ** 2 not in frequencyObj2:    
            return False
        
        if frequencyObj2[value ** 2] != frequencyObj1[value]:
            return False
    
    return True        
    
print(same([1,2,3,2,5,0], [9,1,4,4,25,0]))