function swap(array, i, j) {
    const temp = array[i];
    array[i] = array[j];
    array[j] = temp;
}

function reverse(array, start = 0) {
    for(let i = 0; i < (array.length - start) / 2; i++){
        swap(array, i + start, array.length - 1 - i)
    }

    return array;
}

function firstDecreasingFromRight (array) {
    for(let i = array.length - 2, j = array.length - 1; i >= 0; i--, j--)
        if(array[i] < array[j])
            return i;

    return -1;
}

function getMinumumLargerThan(array, pivotIndex = -1){
    let min = Number.MAX_VALUE;
    let minIndex = -1;

    for(let i = pivotIndex + 1; i < array.length; i++) {
        if(array[i] <= min && array[i] > array[pivotIndex]){
            min = array[i];
            minIndex = i;
        }
    }

    return minIndex;
}

function nextPermutation(array){
    const pivotIndex = firstDecreasingFromRight(array);
    if(pivotIndex < 0)
        return reverse(array);

    const nextIndex = getMinumumLargerThan(array, pivotIndex);
    if(nextIndex < 0)
        return reverse(array, pivotIndex);
    
    swap(array, pivotIndex, nextIndex);
    reverse(array, pivotIndex + 1)

    return array;
}

console.log(nextPermutation([2,3,1,3,3]))