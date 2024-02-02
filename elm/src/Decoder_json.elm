module Decoder_json exposing (..)

import Html exposing (..)
import Json.Decode exposing (Decoder, map2, field, list, string)



-- Définit le type Def utilisé pour les définitions
type alias Def = { word : String,
  meanings : List Meaning
  }


-- Définit le type Meaning utilisé pour une définition en particulier
type alias Meaning =
    { definitions : List String,
      partOfSpeech : String
    }


-- Définnit le décodeur global pour une liste à décoder
defDecoder : Decoder (List Def)
defDecoder =
  list listDecodage


-- Première étape de décodage pour un mot
listDecodage : Decoder Def
listDecodage =
  map2 Def
    (field "word" string)
    (field "meanings" (list meaningDecodage))


-- Deuxième étape de décodage pour les définitions de ce mot
meaningDecodage : Decoder Meaning
meaningDecodage = 
  map2 Meaning
    (field "definitions" (list (field "definition" string)))
    (field "partOfSpeech" string)