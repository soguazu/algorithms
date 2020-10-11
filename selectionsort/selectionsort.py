def selection_sort(arr):
    for i in range(0, len(arr), 1):
        lowest = i
        for j in range(i+1, len(arr)):
            if arr[j] < arr[lowest]:
                lowest = j

        # This condition is checked because if i  and lowest is the same,
		# They carry the same index, so there is no need to swap
        if i != lowest:
            arr[i], arr[lowest] = arr[lowest], arr[i]
       
    print(arr)


selection_sort([4, 9, 3, 14, 11, 50, 34, 23, 12])
selection_sort([1, 2, 3, 4, 5])