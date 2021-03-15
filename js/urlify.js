// Replaces spaces with %20
// string: the string to replace
// length: the lenght of the unpadded string
function urlify(string, length){
    for(let i = 0, j = 0; i < length; i++, j++){
        if(string[length - 1 - i] === ' '){
            string[string.length - 1 - j] = '0'
            string[string.length - 1 - j - 1] = '2'
            string[string.length - 1 - j - 2] = '%'

            j += 2;
        } else {
            string[string.length - 1 - j] = string[length - 1 - i];
        }
    }

    return string.join("");
}

console.log(urlify("Mr John Smith    ".split(''), 13));
console.log(urlify(" c a    ".split(''), 4));
console.log(urlify("There is reason for hope        ".split(''), 24));
console.log(urlify("   ".split(''), 1));
console.log(urlify("".split(''), 0));