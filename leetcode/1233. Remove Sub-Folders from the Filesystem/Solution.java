class Solution {
    public List<String> removeSubfolders(String[] folder) {
        int n = folder.length;
        Map <String, Boolean> freq = new HashMap<>();
        Map <String, Boolean> tick = new HashMap<>();
        List <String> ans = new ArrayList<>();

        for (int i = 0; i < n; i++) freq.put(folder[i], true); 
        
        loop: for (String f : folder) {
            String str = "";
            String[] parts = f.split("/");

            for (int i = 1; i < parts.length; i++) {
                str += "/" + parts[i];
                if (freq.getOrDefault(str, false)) {
                    if (!tick.getOrDefault(str, false)) {
                        tick.put(str, true);
                        ans.add(str);
                    }
                    continue loop;
                }
            }
        }

        return ans;
    }
}