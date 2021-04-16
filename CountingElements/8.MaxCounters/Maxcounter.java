import java.util.Arrays;

class Maxcounter {
    public static void main(String[] args){
        Maxcounter ma = new Maxcounter();
        System.out.println(ma.solution(5,new int[]{3,4,4,6,1,4,4}));
    }

    public int[] solution(int N,int[] A){
        int[] Arr = new int[N];
        for (int i=0;i<A.length; i++){
            if (A[i] >N){
                Arrays.sort(Arr);
                for(int j=0;j<N;j++){
                    Arr[j] = Arr[N-1];
                }
            }else{
                Arr[A[i]-1]++;
            }
        }

        for (int k=0; k<N;k++){
            System.out.println(Arr[k]);
        }

        return Arr;
    }
}
