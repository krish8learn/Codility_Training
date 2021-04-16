import java.util.HashSet;

class Mistinteger {
    public static void main(String[] args){
        Mistinteger mi = new Mistinteger();
        System.out.println(mi.solution(new int[]{1, 3, 6, 4, 1, 2}));
    }
    
    public int solution(int[] A){
        int smallest = 1;
        HashSet<Integer> hs = new HashSet<Integer>();
        for (int i=0; i<A.length; i++){
            hs.add(A[i]);
        }
        System.out.println(hs);
        while(hs.contains(smallest)){
            smallest++;
        }

        return smallest;

    }
}
