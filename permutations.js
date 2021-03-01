function swap(array, i, j){
    const temp = array[i];
    array[i] = array[j];
    array[j] = temp;
}

function generatePermutations(array, size, results){
    if(size === 1) {
        results.push([...array]);
    } else {
        for(let i = 0; i < size; i++){
            const pivot = size % 2 == 1 ? 0 : i;

            generatePermutations(array, size - 1, results);
            swap(array, pivot, size -1);
        }
    }
}


var permute = function(nums) {
    results = []
    generatePermutations(nums, nums.length, results);
    
    return results;
};

console.log(permute(["b", "o", "a", "t"]))