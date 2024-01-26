module Def_writer exposing (..)

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


createDef : List Def -> Int -> List (Html msg)
createDef list n = case list of
    (x::xs) -> (text ((String.fromInt n)++".") ::List.append (showDef x.meanings) (createDef xs (n+1)))
    [] -> []


showDef : List Meaning -> List (Html msg)
showDef meanings = case meanings of
    (x::xs) ->  ((text ("\n\r\n\r - " ++ x.partOfSpeech ++ " : \r\n" ++ writeListDef x.definitions)) :: showDef xs)
    [] -> [text ""]


writeListDef : List String -> String
writeListDef list = case list of
  (x::xs) -> "     - " ++ x ++ "\r\n" ++ (writeListDef xs)
  [] -> "\r\n"