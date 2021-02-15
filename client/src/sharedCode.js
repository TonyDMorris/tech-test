const setRender = (question) => {
    if (question.type === "multiSelect"){
        return {renderButton:false, renderMultiSelect:true, renderText: false};
    }
    else if (question.type === "button"){
        return {renderButton:true, renderMultiSelect:false, renderText: false};
    }
    else if (question.type === "text") {
        return {renderButton:false, renderMultiSelect:false, renderText: true}; 
    }
}

exports.setRender = setRender;