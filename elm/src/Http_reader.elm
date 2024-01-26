module Http_reader exposing (..)

-- Make a GET request to load a book called "Public Opinion"
--
-- Read how it works:
--   https://guide.elm-lang.org/effects/http.html
--

import Browser
import Html exposing (Html, text, pre, div, input, form, h1)
import Http
import Random
import List
import Json.Decode exposing (Decoder, map2, field, list, string)
import Html.Attributes exposing (type_, placeholder, value, style)
import Html.Events exposing (onInput)



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
  | OneWord String
  | SuccessDef { wordToFind : String
  , wordToGuess : String
  , listdef : (List Def)}
  | FailureJSON


type alias Def = { word : String,
  meanings : List Meaning
  }

type alias Meaning =
    { definitions : List String,
      partOfSpeech : String
    }


init : () -> (Model, Cmd Msg)
init _ =
  ( Loading
  , Http.get
      { url = "https://raw.githubusercontent.com/Swar2424/ELP_Project1/main/elm/thousand_words_things_explainer.txt"
      , expect = Http.expectString GotText
      }
  )



-- UPDATE


type Msg
  = GotText (Result Http.Error String)
  | RandomNumber Int
  | GotDef (Result Http.Error (List Def))
  | WordToGuess String



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
          FullText words -> case OneWord (randomWord x (String.split " " words)) of
            OneWord word -> (OneWord word, 
              Http.get
                { url = ("https://api.dictionaryapi.dev/api/v2/entries/en/" ++ (word))
                , expect = Http.expectJson GotDef defDecoder
                })
            Failure -> (Failure, Cmd.none)
            Loading -> (Failure, Cmd.none)
            FullText _ -> (Failure, Cmd.none)
            SuccessDef _ -> (Failure, Cmd.none)
            FailureJSON -> (Failure, Cmd.none)
              

          Failure -> (Failure, Cmd.none)
          Loading -> (Failure, Cmd.none)
          OneWord _ -> (Failure, Cmd.none)
          SuccessDef _ -> (Failure, Cmd.none)
          FailureJSON -> (Failure, Cmd.none)

    GotDef result -> case result of
        Ok def -> case def of
          (x::xs) -> (SuccessDef {wordToGuess = "", wordToFind = x.word, listdef = def}, Cmd.none)
          [] -> (FailureJSON, Cmd.none)
        Err _ -> (FailureJSON, Cmd.none)
    
    WordToGuess wordToGuess -> case model of
      Failure -> (Failure, Cmd.none)
      Loading -> (Failure, Cmd.none)
      FullText _ -> (Failure, Cmd.none)
      SuccessDef results -> (SuccessDef {results | wordToGuess = wordToGuess}, Cmd.none)
      FailureJSON -> (Failure, Cmd.none)
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
      text "There is an error somewhere."

    FailureJSON ->
      text "No dico"

    Loading ->
      text "Loading..."

    FullText fullText ->
      pre [] [ text fullText ]

    OneWord word ->
      pre [] [ text ("https://api.dictionaryapi.dev/api/v2/entries/en/" ++ (word)) ]

    SuccessDef result -> div []
      [ h1 [] [text "Guess the word"], 
      Html.form [] 
      [viewValidation model, viewInput "wordToGuess" "Enter the word to guess" result.wordToGuess WordToGuess
      , pre [] (createDef result.listdef 1)]]


randomWord : Int -> List String -> String
randomWord x list = case pick x list of
    Nothing -> ""
    Just word -> word

roll : Int -> Random.Generator Int
roll n = Random.int 0 n

pick : Int -> List a -> Maybe a
pick n list = List.head (List.drop n list)



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


viewInput : String -> String -> String -> (String -> msg) -> Html msg
viewInput t p v toMsg =
  input [ type_ t, placeholder p, value v, onInput toMsg ] []


viewValidation : Model -> Html msg
viewValidation model = case model of
  SuccessDef result ->
    if result.wordToGuess == result.wordToFind then
      div [ style "color" "green" ] [ text "Niiice, well done !!" ]
    else
      div [ style "color" "red" ] [ text "Find the word" ]
  Failure -> div [ style "color" "green" ] [ text "Bruh" ]
  Loading -> div [ style "color" "green" ] [ text "Bruh" ]
  OneWord _ -> div [ style "color" "green" ] [ text "Bruh" ]
  FailureJSON -> div [ style "color" "green" ] [ text "Bruh" ]
  FullText _ -> div [ style "color" "green" ] [ text "Bruh" ]