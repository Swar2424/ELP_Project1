module Random_picker exposing (..)

import Html exposing (..)
import Random
import List



randomWord : Int -> List String -> String
randomWord x list = case pick x list of
    Nothing -> ""
    Just word -> word


roll : Int -> Random.Generator Int
roll n = Random.int 0 n


pick : Int -> List a -> Maybe a
pick n list = List.head (List.drop n list)