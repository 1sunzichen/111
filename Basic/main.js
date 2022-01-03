function getCount(num){
    let nu=0
    for(let i=num;i>0;i--){
         nu=nu+getNum(i)
    }
    return nu
}
function getNum(num){
    let n=0;
    for(var i=1;i<num-1;i++){
        n+=i
    }
    console.log(n,num,'num')
    return n
}
//4 =>4321  4+32 31  21;  3+ 21=> 2+1 +1
//5 =>54321 5+43 42 41  32 31  21;  4+ 32 31  21  3; 21 =>3+2+1 +2+1 +1
console.log(getCount(7))