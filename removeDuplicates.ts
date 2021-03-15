function removeDuplicates(nums: number[]): number {
    let current = nums[0];
    let lastNonDuplicateIndex = 0;
    let totalDuplicates = 0;

    for (let i = 1; i < nums.length; i++) {
        if (nums[i] > current) {
            lastNonDuplicateIndex++;
            nums[lastNonDuplicateIndex] = nums[i];

            current = nums[i];
        } else {
            totalDuplicates++;
        }
    }

    nums.splice(nums.length - totalDuplicates);
    return nums.length;
};

const input1 = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4];
console.log(removeDuplicates(input1));
console.log(input1);