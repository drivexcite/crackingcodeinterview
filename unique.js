function unique(string){
    const array = string.split("");

    for(let i = 0; i < array.length - 1; i++){
        for(let j = i + 1; j < array.length; j++){
            if(array[i] === array[j])
                return false;
        }
    }

    return true;
}

console.log(unique("abcdefghijklmnopqrstuvwxyz"));
console.log(unique("abcdefghijklmnopqrstuvwxyzz"));
console.log(unique("aa"));
console.log(unique("a"));
console.log(unique(""));