module Main exposing (..)


import Decoder_json exposing (..)
import Def_writer exposing (..)
import Random_picker exposing (..)
import Browser
import Html exposing (..)
import Http
import Random
import List
import Html.Attributes exposing (type_, placeholder, value, style)
import Html.Events exposing (onInput, onClick)



-- MAIN

main : Program () Model Msg
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
  | OneWord String String
  | SuccessDef { wordToFind : String
  , wordToGuess : String
  , listdef : (List Def)
  , reveal_word : Bool
  , listwords : String}
  | FailureJSON



-- MSG

type Msg
  = GotText (Result Http.Error String)
  | RandomNumber Int
  | GotDef (Result Http.Error (List Def))
  | WordToGuess String
  | Checkbox



-- INIT

init : () -> (Model, Cmd Msg)
init _ =
  ( Loading
  , Http.get
      { url = "https://raw.githubusercontent.com/Swar2424/ELP_Project1/main/elm/thousand_words_things_explainer.txt"
      , expect = Http.expectString GotText
      }
  )



-- UPDATE

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
          FullText words -> case OneWord (randomWord x (String.split " " words)) words of
            OneWord word listword -> (OneWord word listword, 
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
          OneWord _ _ -> (Failure, Cmd.none)
          SuccessDef _ -> (Failure, Cmd.none)
          FailureJSON -> (Failure, Cmd.none)

    GotDef result -> case result of
        Ok def -> case def of
          (x::_) -> case model of
              Failure -> (Failure, Cmd.none)
              Loading -> (Failure, Cmd.none)
              FullText _ -> (Failure, Cmd.none)
              SuccessDef _ -> (Failure, Cmd.none)
              FailureJSON -> (Failure, Cmd.none)
              OneWord word listwords-> (SuccessDef {wordToGuess = "", wordToFind = word, listdef = def, reveal_word = False, listwords = listwords}, Cmd.none)
          [] -> (FailureJSON, Cmd.none)
        Err _ -> (FailureJSON, Cmd.none)
    
    WordToGuess wordToGuess -> case model of
      Failure -> (Failure, Cmd.none)
      Loading -> (Failure, Cmd.none)
      FullText _ -> (Failure, Cmd.none)
      SuccessDef results -> (SuccessDef {results | wordToGuess = wordToGuess}, Cmd.none)
      FailureJSON -> (Failure, Cmd.none)
      OneWord _ _-> (Failure, Cmd.none)
    
    Checkbox -> case model of 
      Failure -> (Failure, Cmd.none)
      Loading -> (Failure, Cmd.none)
      FullText _ -> (Failure, Cmd.none)
      SuccessDef results -> (SuccessDef {results | reveal_word = not results.reveal_word}, Cmd.none)
      FailureJSON -> (Failure, Cmd.none)
      OneWord _ _ -> (Failure, Cmd.none)
      

          
-- SUBSCRIPTIONS

subscriptions : Model -> Sub Msg
subscriptions _ =
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
      text "Loading word..."

    FullText _ ->
      text "Loading word..."

    OneWord _ _ ->
      text "Loading def..."

    SuccessDef result -> div []
      [ h1 [style "text-align" "center"] [text ("Guess the word    (～￣▽￣)～")]
      , h2 [style "text-align" "center", style "color" "green"] [text (if result.reveal_word then ("\n It's " ++ result.wordToFind ++ " !\n") else "")]
      , pre [] (createDef result.listdef 1)
      , div [style "text-align" "center"] [ h1 [] [viewValidation model, viewInput "wordToGuess" "Enter the word to guess" result.wordToGuess WordToGuess],
          button [ onClick Checkbox ] [ text "Reveal word ?" ], pre [] [text ""],
          button [ onClick (GotText (Ok result.listwords)) ] [ text "Reroll word ?" ], pre [] [text "\n\r"]]
      ]



-- Définition de la zone d'input
viewInput : String -> String -> String -> (String -> msg) -> Html msg
viewInput t p v toMsg =
  input [ type_ t, placeholder p, value v, onInput toMsg ] []


-- Définition du comportement en fonction du contenu de la zone d'input
viewValidation : Model -> Html msg
viewValidation model = case model of
  SuccessDef result ->
    if result.wordToGuess == result.wordToFind then
      div [ style "color" "green" ] [ text "Niiice, well done !!" ]
    else
      div [ style "color" "red" ] [ text "Find the word" ]
  Failure -> div [ style "color" "orange" ] [ text "Error" ]
  Loading -> div [ style "color" "orange" ] [ text "Error" ]
  OneWord _ _ -> div [ style "color" "orange" ] [ text "Error" ]
  FailureJSON -> div [ style "color" "orange" ] [ text "Error" ]
  FullText _ -> div [ style "color" "orange" ] [ text "Error" ]