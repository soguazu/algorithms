function same(arr1: number[], arr2: number[]): boolean {
  let frequencyObj1: object = {};
  let frequencyObj2: object = {};
  for (let value of arr1) {
    frequencyObj1[value] = ++frequencyObj1[value] || 1;
  }

  for (let value of arr2) {
    frequencyObj2[value] = ++frequencyObj2[value] || 1;
  }

  for (let temp in frequencyObj1) {
    if (!(Number(temp) ** 2 in frequencyObj2)) {
      return false;
    }

    if (frequencyObj2[Number(temp) ** 2] !== frequencyObj1[Number(temp)]) {
      return false;
    }
  }
  return true;
}

console.log(same([1, 2, 3, 2, 5, 0], [9, 1, 4, 4, 25, 0]));
