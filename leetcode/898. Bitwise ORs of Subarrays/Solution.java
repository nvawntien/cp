class Solution {
    public int subarrayBitwiseORs(int[] arr) {
        Set <Integer> cur = new HashSet<>();
        Set <Integer> res = new HashSet<>();

        for (int num : arr) {
            Set <Integer> newCur = new HashSet<>();
            newCur.add(num);

            for (int or : cur) {
                newCur.add(num | or);
            }

            cur = newCur;
            res.addAll(cur);
        }

        return res.size();
    }
}