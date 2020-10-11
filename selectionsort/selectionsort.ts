function selectionSort(arr: number[]) {
  for (let i = 0; i < arr.length; i++) {
    let lowest = i;
    for (let j = i + 1; j < arr.length; j++) {
      if (arr[j] < arr[lowest]) {
        lowest = j;
      }
    }
    // This condition is checked because if i  and lowest is the same,
    // They carry the same index, so there is no need to swap
    if (i !== lowest) [arr[i], arr[lowest]] = [arr[lowest], arr[i]];
  }
  console.log(arr);
}

selectionSort([4, 9, 3, 14, 11, 50, 34, 23, 12]);
selectionSort([1, 2, 3, 4, 5]);
