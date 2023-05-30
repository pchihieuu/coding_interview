import java.util.Arrays;

class hIndex {
    public int hightIndex(int[] citations) {
        final int n = citations.length;

        Arrays.sort(citations);

        for (int i = 0; i < n; ++i)
            if (citations[i] >= n - i)
                return n - i;

        return 0;
    }
}
