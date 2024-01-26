module Random_picker exposing (..)

import Browser
import Html exposing (..)
import Http
import Random
import List
import Json.Decode exposing (Decoder, map2, field, list, string)
import Html.Attributes exposing (type_, placeholder, value, style)
import Html.Events exposing (onInput, onClick)



randomWord : Int -> List String -> String
randomWord x list = case pick x list of
    Nothing -> ""
    Just word -> word


roll : Int -> Random.Generator Int
roll n = Random.int 0 n


pick : Int -> List a -> Maybe a
pick n list = List.head (List.drop n list)