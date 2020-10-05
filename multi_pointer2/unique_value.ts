function countUniqueValues(arr: number[]): number {
  let start = 0;
  let step = 1;
  for (let i = 0; i < arr.length; i++) {
    if (arr[start] !== arr[step]) {
      start++;
      arr[start] = arr[step];
      step++;
    } else {
      step++;
    }
  }
  return start;
}

console.log(countUniqueValues([1, 1, 1, 1, 1, 2]));
console.log(countUniqueValues([1, 2, 3, 4, 4, 4, 7, 7, 12, 12, 13]));
console.log(countUniqueValues([-2, -1, -1, 0, 1]));
console.log(countUniqueValues([]));
