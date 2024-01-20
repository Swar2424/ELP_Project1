module Http_reader exposing (..)

-- Make a GET request to load a book called "Public Opinion"
--
-- Read how it works:
--   https://guide.elm-lang.org/effects/http.html
--

import Browser
import Html exposing (Html, text, pre)
import Http
import Random
import List



-- MAIN


main =
  Browser.element
    { init = init
    , update = update
    , subscriptions = subscriptions
    , view = view
    }



-- MODEL


type Model
  = Failure
  | Loading
  | FullText String
  | OneWord { word : String}


init : () -> (Model, Cmd Msg)
init _ =
  ( Loading
  , Http.get
      { url = "https://raw.githubusercontent.com/Swar2424/ELP_Project1/main/elm/thousand_words_things_explainer.txt" -- "https://elm-lang.org/assets/public-opinion.txt"
      , expect = Http.expectString GotText                      -- "https://perso.liris.cnrs.fr/tristan.roussillon/GuessIt/thousand_words_things_explainer.txt"
      }
  )



-- UPDATE


type Msg
  = GotText (Result Http.Error String)
  | RandomNumber Int



update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
  case msg of
    GotText result ->
      case result of
        Ok fullText ->
          (FullText fullText, Random.generate RandomNumber (roll (List.length (String.split " " fullText))))

        Err _ ->
          (Failure, Cmd.none)
        
    RandomNumber x -> case model of
          FullText words -> (OneWord{ word = randomWord x (String.split " " words)}, Cmd.none)
          Failure -> (Failure, Cmd.none)
          Loading -> (Failure, Cmd.none)
          OneWord _ -> (Failure, Cmd.none)



-- SUBSCRIPTIONS


subscriptions : Model -> Sub Msg
subscriptions model =
  Sub.none



-- VIEW


view : Model -> Html Msg
view model =
  case model of
    Failure ->
      text "I was unable to load your book."

    Loading ->
      text "Loading..."

    FullText fullText ->
      pre [] [ text fullText ]

    OneWord word ->
      pre [] [ text word.word ]


randomWord : Int -> List String -> String
randomWord x list = case pick x list of
    Nothing -> ""
    Just word -> word

roll : Int -> Random.Generator Int
roll n = Random.int 0 n

pick : Int -> List a -> Maybe a
pick n list = List.head (List.drop n list)

getLast : (List a) -> (Maybe a)
getLast liste = case List.reverse liste of
    (x::xs) -> Just x
    [] -> Nothing