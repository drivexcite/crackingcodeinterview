function rotate(nums: number[], k: number): number[] {
    if (k < 1) return nums;

    const input = [...nums];

    for (let i = 0; i < nums.length; i++) {
        nums[(i + k) % nums.length] = input[i];
    }

    return nums;
}

// console.log(rotate([1, 2, 3, 4, 5], 2));
// console.log(rotate([1, 2, 3, 4, 5, 6, 7], 3));
console.log(rotate([-1, -100, 3, 99], 2));
// console.log(rotate([1, 2, 3], 2));