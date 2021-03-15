
function maxProfit(stockPrice: number[]): number {
    let maxProfit = 0;

    for (let i = 1; i < stockPrice.length; i++) {
        if (stockPrice[i] > stockPrice[i - 1]) {
            maxProfit += stockPrice[i] - stockPrice[i - 1];
        }
    }

    return maxProfit;
}

console.log(maxProfit([7, 1, 5, 3, 6, 4]));
console.log(maxProfit([1, 7, 2, 3, 6, 7, 6, 7]));