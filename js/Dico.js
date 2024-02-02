async function get_word(name) {
    var dictionary = await fetch(`https://api.dictionaryapi.dev/api/v2/entries/en/${name}`)
    var convert = await dictionary.json()
    if (convert.title == 'No Definitions Found'){
        word = 'error'
    } else{
        var word = convert[0].word
    }
    return word
}

get_word('noun').then(
    word => console.log(word)
)


module.exports = {get_word};