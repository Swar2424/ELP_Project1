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
import Json.Decode exposing (Decoder, map2, map3, field, list, string)



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
  | SuccessDef (List Def)
  | FailureBad


type alias Def = { word : String,
  phonetic : String,
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
            FailureBad -> (Failure, Cmd.none)
              

          Failure -> (Failure, Cmd.none)
          Loading -> (Failure, Cmd.none)
          OneWord _ -> (Failure, Cmd.none)
          SuccessDef _ -> (Failure, Cmd.none)
          FailureBad -> (Failure, Cmd.none)

    GotDef result -> case result of
        Ok def -> (SuccessDef def, Cmd.none)
        Err _ -> (FailureBad, Cmd.none)
          



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

    FailureBad ->
      text "No dico"

    Loading ->
      text "Loading..."

    FullText fullText ->
      pre [] [ text fullText ]

    OneWord word ->
      pre [] [ text ("https://api.dictionaryapi.dev/api/v2/entries/en/" ++ (word)) ]

    SuccessDef listdef -> case listdef of 
      (x::xs) -> createDef x
      _ -> pre [] [ text "bruh"]


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
  map3 Def
    (field "word" string)
    (field "phonetic" string)
    (field "meanings" (list meaningDecodage))

meaningDecodage : Decoder Meaning
meaningDecodage = 
  map2 Meaning
    (field "definitions" (list (field "definition" string)))
    (field "partOfSpeech" string)


createDef : Def -> Html msg
createDef x = pre [] [ text x.word ]