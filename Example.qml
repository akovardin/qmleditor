import QtQuick 2.0
import QtQuick.Controls 1.0
import QtQuick.Dialogs 1.0

Rectangle {
    width: 360
    height: 360
    color: "grey"

    TextArea {
        id: textArea
        x: 8
        y: 74
        width: 344
        height: 278
        text: editor.text
    }

    ExampleButton {
        id: loadButton
        x: 8
        y: 8
        text: "Load"
        onClicked: {
            console.log("Load")
            fileDialogLoad.open()
        }
    }

    ExampleButton {
        id: saveButton
        x: 120
        y: 8
        text: "Save"
        onClicked: {
            console.log("Save")
            editor.saveFile(textArea.text)
        }
    }

    FileDialog {
        id: fileDialogLoad
        folder: "."
        title: "Choose a file to open"
        selectMultiple: false
        // nameFilters: [ "Image files (*.png *.jpg)", "All files (*)" ]
        onAccepted: { 
            console.log("Accepted: " + fileDialogLoad.fileUrl) 
            editor.selectFile(fileDialogLoad.fileUrl)
        }
    }
}