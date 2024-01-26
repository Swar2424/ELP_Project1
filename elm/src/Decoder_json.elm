module Decoder_json exposing (..)

import Browser
import Html exposing (..)
import Http
import Random
import List
import Json.Decode exposing (Decoder, map2, field, list, string)
import Html.Attributes exposing (type_, placeholder, value, style)
import Html.Events exposing (onInput, onClick)



type alias Def = { word : String,
  meanings : List Meaning
  }


type alias Meaning =
    { definitions : List String,
      partOfSpeech : String
    }


defDecoder : Decoder (List Def)
defDecoder =
  list listDecodage


listDecodage : Decoder Def
listDecodage =
  map2 Def
    (field "word" string)
    (field "meanings" (list meaningDecodage))


meaningDecodage : Decoder Meaning
meaningDecodage = 
  map2 Meaning
    (field "definitions" (list (field "definition" string)))
    (field "partOfSpeech" string)