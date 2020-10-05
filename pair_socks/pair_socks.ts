function sockMerchant(arr) {
  let temp = {};
  let counter: number = 0;

  for (let value of arr) {
    temp[value] = ++temp[value] || 1;
  }

  for (let value in temp) {
    if (temp[value] % 2 !== 0 && temp[value] > 1) {
      counter += (temp[value] - (temp[value] % 2)) / 2;
    }

    if (temp[value] % 2 === 0) {
      counter += temp[value] / 2;
    }
  }

  return counter;
}

console.log(sockMerchant([10, 20, 20, 10, 10, 30, 50, 1020]));
console.log(sockMerchant([10, 20, 20, 10, 10, 10, 30, 50, 50, 50, 1020]));
