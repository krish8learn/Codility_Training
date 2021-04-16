import java.util.Arrays;

class Distinct{
    public static void main(String[] args) {
        Distinct di = new Distinct();
        System.out.println(di.solution(new int[]{2, 1, 1, 2, 3, 1}));
    }

    public int solution(int[] A){
        if (A.length==0){
            return 0;
        }
        int count =1;
        Arrays.sort(A);
        for (int i=0;i<A.length-1; i++){
            if (A[i]!=A[i+1]){
                count++;
            }
        }
        return count;
    }
}