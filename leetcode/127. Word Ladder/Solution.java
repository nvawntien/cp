class Solution {
    private List<List<Integer>> edges = new ArrayList<>();
    
    public int ladderLength(String beginWord, String endWord, List<String> wordList) {
        if (!wordList.contains(endWord) || beginWord.equals(endWord)) return 0;
        wordList.add(0, beginWord);

        int n = wordList.size(), index = 0;
        boolean[] visited = new boolean[n+1];

        for (int i = 0; i < n; i++) {
            edges.add(new ArrayList<>());
            if (wordList.get(i).equals(endWord)) index = i;
        }

        for (int i = 0; i < n-1; i++) {
            for (int j = i+1; j < n; j++) {
                if (compare(wordList.get(i), wordList.get(j))) {
                    edges.get(i).add(j);
                    edges.get(j).add(i);
                }
            }
        }

    
        return bfs(0, n, visited, index);
    }

    private int bfs(int parent, int n, boolean[] visited, int i) {
        Queue <Integer> Q = new LinkedList<>();
        Q.add(parent);
        int[] dist = new int[n];

        visited[parent] = true;

        while (!Q.isEmpty()) {
            int u = Q.poll();
            for (int v : edges.get(u)) {
                if (!visited[v]) {
                    dist[v] = dist[u] + 1;
                    Q.add(v);
                    visited[v] = true;
                }
            }
        }

        return visited[i] ? dist[i]+1:0;
    }

    private boolean compare(String a, String b) {
        int count = 0;
        for (int i = 0; i < a.length(); i++) {
            if (a.charAt(i) != b.charAt(i)) count++;
            if (count > 1) return false;
        }

        return count == 1;
    }
}