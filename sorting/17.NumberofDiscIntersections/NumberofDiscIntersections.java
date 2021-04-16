import java.util.Arrays;
class NumberofDiscIntersections{
    public static void main(String[] args) {
        NumberofDiscIntersections nb = new NumberofDiscIntersections();
        System.out.println(nb.Solution(new int[]{1,5,2,1,4,0}));
    }

    public int Solution(int[] A){
        long[] start = new long[A.length];
        long[] end = new long[A.length];

        for(int i=0; i<A.length; i++){
            start[i] = i +(long)A[i];
            end[i] = i-(long)A[i];
        }

        Arrays.sort(start);
        Arrays.sort(end);

        int result = 0, j=0;

        for(int i=0; i<A.length; i++){
            while(j<A.length && start[i]>=end[j]){
                result = result +j;
                result = result-i;
                j++;
            }
        }

        if(result>10000000){
            return -1;
        }
        return result;
    }
}