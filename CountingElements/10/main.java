import java.util.Arrays;

class Permcheck {
    public static void main(String[] args){
        Permcheck m1 = new Permcheck();
        //System.out.println(m1.solution(new int[]{4,1,3,2}));
        System.out.println(m1.solution(new int[]{1}));
    }

    public int solution(int[] A){
        Arrays.sort(A);
        int N = A.length;
        if ((N<2)&& A[0]==1){
            return 1;
        }else if (N<2){
            return 0;
        }
        if (A[0]==1){
            for (int i=0;i<N-1;i++){
                if (A[i]+1 != A[i+1]){
                    //System.out.println(A[i]+","+A[i+1]);
                    return 0;
                }
            }
        }else {
            return 0;
        }
        
        return 1;
    }
    
}
