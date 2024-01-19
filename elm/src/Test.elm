module Test exposing (..)

import Browser
import Html exposing (Html, div, text)
import File exposing (File)
import Task exposing (Task)

main =
    Browser.sandbox { init = init, update = update, view = view }

-- Model
type alias Model =
    { fileContent : String }

init : Model
init =
    { fileContent = "" }

-- Msg
type Msg
    = FileSelected (File String)

-- Update
update : Msg -> Model -> Model
update msg model =
    case msg of
        FileSelected file ->
            { model | fileContent = file.content }

-- View
view : Model -> Html Msg
view model =
    div []
        [ input [ type_ "file", onChange FileSelected ] []
        , div [] [ text model.fileContent ]
        ]

-- File input
input : List (Html.Attribute Msg) -> List (Html Msg) -> Html Msg
input attributes children =
    Html.node "input" attributes children