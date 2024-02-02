module Def_writer exposing (..)

import Html exposing (..)
import List
import Decoder_json exposing (..)



-- Renvoie un message Html pour un mot
createDef : List Def -> Int -> List (Html msg)
createDef list n = case list of
    (x::xs) -> (text ((String.fromInt n)++".") ::List.append (showDef x.meanings) (createDef xs (n+1)))
    [] -> []


-- Renvoie un message Html pour une catégorie de définitions d'un mot
showDef : List Meaning -> List (Html msg)
showDef meanings = case meanings of
    (x::xs) ->  ((text ("\n\r\n\r - " ++ x.partOfSpeech ++ " : \r\n" ++ writeListDef x.definitions)) :: showDef xs)
    [] -> [text ""]


-- Renvoie un message Html pour une définition spécifique d'un mot
writeListDef : List String -> String
writeListDef list = case list of
  (x::xs) -> "     - " ++ x ++ "\r\n" ++ (writeListDef xs)
  [] -> "\r\n"