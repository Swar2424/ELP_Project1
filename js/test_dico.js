async function get_word(name) {
    var dictionary = await fetch(`https://api.dictionaryapi.dev/api/v2/entries/en/${name}`)
    var convert = await dictionary.json()
    var word = convert[0].word
    console.log (word)
}

get_word("noun")