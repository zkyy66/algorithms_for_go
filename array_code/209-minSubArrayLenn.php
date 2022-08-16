class Solution {
    /**
     * @param Integer $target
     * @param Integer[] $nums
     * @return Integer
     */
    function minSubArrayLen($target, $nums) {
        if (count($nums) < 1) {
            return 0;
        }
        $sum = 0;
        $res = PHP_INT_MAX;
        $left = 0;
        for ($right = 0; $right < count($nums); $right++) {
            $sum += $nums[$right];
            while ($sum >= $target) {
                $res = min($res, $right - $left + 1);
                $sum -= $nums[$left];
                $left++;
            }
        }
        return $res == PHP_INT_MAX ? 0 : $res;
    }
}