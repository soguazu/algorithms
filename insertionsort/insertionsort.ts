function insertionSort(arr: number[]): number[] {
  for (let i = 1; i < arr.length; i++) {
    let currentValue = arr[i];
    let index = i;
    for (let j = i - 1; j >= 0 && arr[j] > currentValue; j--) {
      arr[j + 1] = arr[j];
      index--;
    }
    arr[index] = currentValue;
  }
  return arr;
}

console.log(insertionSort([4, 9, 3, 14, 11, 50, 34, 23, 12]));
console.log(insertionSort([2, 1, 9, 76, 4]));
console.log(insertionSort([1, 2, 3, 4, 6]));
