import { Input, Alert } from 'reactstrap';
import React, { useState, useContext} from 'react';
import { RenderContext } from '../contexts/RenderContext';

const RenderMultiSelect = (props) => {

    const [state, setState] = useState({ selected: [] });

    const { renderState, dispatch } = useContext(RenderContext);
    const question = renderState.questionMap.get(renderState.currentQuestionId);

    const renderOption = (answer) => {
        return <option key={answer.id} id={answer.id}>{answer.text}</option>
    }

    const renderAlert = () => {
        console.log("render alert");
        return <Alert>Please select one or more answers from the select.</Alert>
    }

    const handleChange = (event) => {
        let options = event.target.options;
        let value = [];
        for (var i = 0, l = options.length; i < l; i++) {
            if (options[i].selected) {
                value.push(options[i].id);
            }
        }

        setState({ selected: value });
    }

    const handleClick = () => {
        if (state.selected.length > 0) {
            dispatch({ type: 'AnswerQuestion', update: { answer: state.selected, nextPage: props.nextPage } });
        }
        else {
            setState({...state, viewAlert: true});
        }
    }

    return (
        <div>
            <div>
                <b>{question.title}</b>
            </div>
            <div>
                <Input type="select" name="select" id="exampleSelect" multiple onChange={handleChange}>
                    {question.answers.map(answer => {
                        return renderOption(answer);
                    })}
                </Input>
            </div>

            <div>
                <button onClick={() => handleClick()}>Continue</button>
            </div>

            {state.viewAlert && renderAlert()}
        </div>
    )
}

export default RenderMultiSelect;

