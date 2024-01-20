module Project_1 exposing (..)

import Browser
import Html exposing (..)
import Html.Attributes exposing (style)
import Html.Events exposing (..)
import Http
import Json.Decode exposing (Decoder, map4, field, int, string, decodeString, float, nullable)
import File exposing (File)

type Model
  = Failure
  | Loading
  | Success String

type alias Def = {word : String}

type Msg
  = MorePlease
  | GotWord (Result Http.Error String)


main =
  Browser.element
    { init = init
    , update = update
    , subscriptions = subscriptions
    , view = view
    }


init : () -> (Model, Cmd Msg)
init _ =
  (Loading, getRandomWord)


update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
  case msg of
    MorePlease ->
      (Loading, getRandomWord)

    GotWord result ->
      case result of
        Ok def ->
          (Success def, Cmd.none)

        Err _ ->
          (Failure, Cmd.none)


subscriptions : Model -> Sub Msg
subscriptions model =
  Sub.none


view : Model -> Html Msg
view model =
  div []
    [ h2 [] [ text "Random Quotes" ]
    , viewDef model
    ]


viewDef : Model -> Html Msg
viewDef model =
  case model of
    Failure ->
      div []
        [ text "I could not load a word for some reason. "
        , button [ onClick MorePlease ] [ text "Try Again!" ]
        ]

    Loading ->
      text "Loading..."

    Success def ->
      div []
        [ button [ onClick MorePlease, style "display" "block" ] [ text "More Please!" ]
        , blockquote [] [ text def]
        , p [ style "text-align" "right" ]
            [ text "— "
            , text (def)
            ]
        ]


getRandomWord : Cmd Msg
getRandomWord = 
    Http.get
    { url = "https://perso.liris.cnrs.fr/tristan.roussillon/GuessIt/thousand_words_things_explainer.txt"
    , expect = Http.expectString GotWord
    }