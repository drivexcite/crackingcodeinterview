function rotate(nums: number[], k: number): number[] {
    for (let i = 0; i < k; i++) {
        const temp = nums[nums.length - 1];

        for (let j = nums.length - 1; j > 0; j--) {
            nums[j] = nums[j - 1];
        }

        nums[0] = temp;
    }

    return nums;
}

console.log(rotate([1, 2, 3, 4, 5], 2));
console.log(rotate([-1, -100, 3, 99], 2));