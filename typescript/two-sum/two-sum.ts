function twoSum(nums: number[], target: number): number[] {
  let output: number[] = [];

  nums.forEach((baseNum: number, index: number) => {
    if (!output.length) {
      for (let i = index + 1; i < nums.length; i++) {
        if(baseNum + nums[i] === target) {
          output = [index, i];
          break;
        }
      }
    }
  });
  return output;
};