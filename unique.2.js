function unique(string){
    const dictionary = {};
    const array = string.split("");

    for(let i = 0; i < array.length; i++){
        if(dictionary[array[i]])
            return false;

        dictionary[array[i]] = true;
    }

    return true;
}

console.log(unique("abcdefghijklmnopqrstuvwxyz"));
console.log(unique("abcdefghijklmnopqrstuvwxyzz"));
console.log(unique("aa"));
console.log(unique("a"));
console.log(unique(""));