module Random_picker exposing (..)

import Html exposing (..)
import Random
import List



-- Renvoie un string à l'indice x d'une liste list de string
randomWord : Int -> List String -> String
randomWord x list = case List.head (List.drop x list) of
    Nothing -> ""
    Just word -> word


-- Renvoie un message de type Generator avec un Int associé
roll : Int -> Random.Generator Int
roll n = Random.int 0 n